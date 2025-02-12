import asyncio
import json
import sys
import os
import time
import logging
from matplotlib import pyplot as plt
import aiohttp
import websockets
import osmnx as ox

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), 'src')))

from services.graph_loader import GraphLoader
from models.car import Car
from visualization.animator import Animator
from utils.logger import configure_logger

VEHICLES_API = "http://localhost:5000/simulation/vehicle?limit=400&per_page=400"
WS_URL = "ws://localhost:5000/ws/location?test_id={}"
WS_PASSENGER_URL = "ws://localhost:5000/ws/passenger?test_id={}"
ROUTES_API = "http://localhost:5000/simulation/path?limit=1000&per_page=1000"

async def fetch_vehicles():
    async with aiohttp.ClientSession() as session:
        try:
            async with session.get(VEHICLES_API) as response:
                data = await response.json()
                return data.get("data", [])
        except Exception as e:
            logging.getLogger(__name__).error(f"Error fetching vehicles: {e}")
            return []

async def fetch_routes():
    async with aiohttp.ClientSession() as session:
        try:
            async with session.get(ROUTES_API) as response:
                data = await response.json()
                return data.get("data", [])
        except Exception as e:
            logging.getLogger(__name__).error(f"Error fetching routes: {e}")
            return []

async def send_location(car):
    while True:
        try:
            async with websockets.connect(
                WS_URL.format(car.vehicle_id),
                ping_interval=15,
                ping_timeout=10,
                close_timeout=5
            ) as ws:
                while True:
                    if car.current_position and car.has_valid_route:
                        await ws.send(json.dumps({
                            'vehicle_id': car.vehicle_id,
                            'lat': car.current_position[1],
                            'lon': car.current_position[0],
                            'created_at': time.time()
                        }))
                    await asyncio.sleep(1)
        except (websockets.exceptions.ConnectionClosed, asyncio.CancelledError) as e:
            logging.error(f"WebSocket closed for vehicle {car.vehicle_id}: {e}")
            await asyncio.sleep(2)

async def listen_destination(car):
    dest_ws_url = f"ws://localhost:5000/ws/destination?test_id={car.driver_id}"
    while True:
        try:
            async with websockets.connect(dest_ws_url) as ws:
                logger = logging.getLogger(__name__)
                logger.info(f"Connected to destination WebSocket for driver {car.driver_id}")
                while True:
                    message = await ws.recv()
                    data = json.loads(message)
                    if "route" in data:
                        car.reset_for_new_route()
                        car.set_destination(data["route"]["station_a"], data["route"]["station_b"])
        except Exception as e:
            logger.error(f"WebSocket error for driver {car.driver_id}: {e}")
            await asyncio.sleep(2)

async def send_arrival_notification(car):
    url = "http://localhost:5000/api/arrival"
    data = {"vehicle_id": car.vehicle_id, "status": "arrived"}
    try:
        async with aiohttp.ClientSession() as session:
            async with session.post(url, json=data) as response:
                if response.status == 200:
                    logging.info(f"Car {car.vehicle_id} arrival notified")
    except Exception as e:
        logging.error(f"Arrival notification failed: {e}")

async def check_arrivals(cars):
    while True:
        for car in [c for c in cars if c.arrived]:
            await send_arrival_notification(car)
            car.arrived = False
        await asyncio.sleep(1)

async def main():
    logger = configure_logger()
    try:
        loader = GraphLoader()
        logger.info("Loading graph data...")
        graph = loader.load_graph()

        vehicles, routes_data = await asyncio.gather(
            fetch_vehicles(),
            fetch_routes()
        )

        cars = [Car(graph, vehicle_data) for vehicle_data in vehicles]
        animator = Animator(ox.graph_to_gdfs(graph)[1], routes_data)

        tasks = [
            *[asyncio.create_task(send_location(car)) for car in cars],
            *[asyncio.create_task(listen_destination(car)) for car in cars],
            asyncio.create_task(animator.run_animation(cars)),
            asyncio.create_task(check_arrivals(cars))
        ]

        await asyncio.gather(*tasks)
    except KeyboardInterrupt:
        logger.info("Simulation stopped")
    except Exception as e:
        logger.error(f"Critical error: {e}", exc_info=True)
    finally:
        if 'animator' in locals():
            animator.close()

if __name__ == "__main__":
    asyncio.run(main())
