# main.py (optimized)
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

VEHICLES_API = "http://localhost:5000/simulation/vehicle?limit=1&per_page=1"
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
            async with websockets.connect(WS_URL.format(car.vehicle_id), ping_interval=10, ping_timeout=20) as ws:
                while True:
                    if car.current_position:
                        await ws.send(json.dumps({
                            'vehicle_id': car.vehicle_id,
                            'lat': car.current_position[1],
                            'lon': car.current_position[0],
                            'created_at': time.time()
                        }))
                    await asyncio.sleep(1)
        except (websockets.exceptions.ConnectionClosed, asyncio.CancelledError) as e:
            logging.error(f"WebSocket closed for vehicle {car.vehicle_id}: {e}")
            await asyncio.sleep(5)  # Wait before reconnecting

async def listen_destination(car):
    dest_ws_url = f"ws://localhost:5000/ws/destination?test_id={car.vehicle_id}"
    async with websockets.connect(dest_ws_url) as ws:
        while True:
            message = await ws.recv()
            data = json.loads(message)
            if "station_a" in data and "station_b" in data:
                car.set_destination(data["station_a"], data["station_b"])
            await asyncio.sleep(1)

# main.py (updated)
async def main():
    logger = configure_logger()
    try:
        # Initialize graph loader
        loader = GraphLoader()

        # Load graph with timing
        logger.info("Loading graph data...")
        start_time = time.time()
        graph = loader.load_graph()
        logger.info(f"Graph loaded in {time.time()-start_time:.2f}s")

        # Convert graph to GeoDataFrames
        _, edges = ox.graph_to_gdfs(graph)  # We only need edges for visualization

        # Parallel data fetching
        logger.info("Fetching initial data...")
        vehicles, routes_data = await asyncio.gather(
            fetch_vehicles(),
            fetch_routes()
        )

        # Initialize components
        cars = [Car(graph, vehicle_data) for vehicle_data in vehicles]
        animator = Animator(edges, routes_data)  # Pass only edges and routes_data

        # Create async tasks
        tasks = [
            *[asyncio.create_task(send_location(car)) for car in cars],
            *[asyncio.create_task(listen_destination(car)) for car in cars],
            asyncio.create_task(animator.run_animation(cars))
        ]

        await asyncio.gather(*tasks)
    except KeyboardInterrupt:
        logger.info("Simulation stopped by user")
    except Exception as e:
        logger.error(f"Critical error: {e}", exc_info=True)
    finally:
        if 'animator' in locals():
            animator.close()

if __name__ == "__main__":
    asyncio.run(main())
