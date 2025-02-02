import asyncio
import json
import sys
import os
from matplotlib import pyplot as plt
import networkx as nx

import aiohttp
import websockets

# Add the src directory to the Python path
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), 'src')))

import osmnx as ox
from services.graph_loader import GraphLoader
from models.car import Car
from visualization.animator import Animator
from utils.logger import configure_logger
import time
import logging

VEHICLES_API = "http://localhost:5000/simulation/vehicle/"
WS_URL = "ws://localhost:5000/ws/location?test_id={}"  # WebSocket URL

async def fetch_vehicles():
    """Fetch vehicle data from API"""
    logger = logging.getLogger(__name__)
    async with aiohttp.ClientSession() as session:
        try:
            async with session.get(VEHICLES_API) as response:
                data = await response.json()
                logger.debug("Fetched vehicle data: %s", data)
                return data.get("data", [])
        except Exception as e:
            logger.error("Error fetching vehicles: %s", e)
            return []

async def run_animation(animator, cars):
    """High-frequency animation updates"""
    animator.init_cars(cars)
    plt.show(block=False)

    try:
        while True:
            # Update positions continuously
            for car in cars:
                car.update_position()

            animator.update_display()
            await asyncio.sleep(0.05)  # ~20 FPS
    except asyncio.CancelledError:
        animator.close()

async def send_location(car):
    """Send current position without modifying it"""
    async with websockets.connect(WS_URL.format(car.vehicle_id)) as ws:
        while True:
            if car.current_position:
                await ws.send(json.dumps({
                    'vehicle_id': car.vehicle_id,
                    'lat': car.current_position[1],
                    'lon': car.current_position[0],
                    'speed': car.speed,
                    'created_at': time.time()
                }))
            await asyncio.sleep(1)  # 1 second updates
async def main():
    logger = configure_logger()
    try:
        # Load graph data
        logger.info("Loading graph data...")
        start_time = time.time()
        loader = GraphLoader()
        graph = loader.load_graph()
        nodes, edges = ox.graph_to_gdfs(graph)
        logger.info(f"Graph loaded in {time.time()-start_time:.2f}s. Nodes: {len(nodes)}, Edges: {len(edges)}")

        # Fetch vehicles
        logger.info("Fetching vehicle data...")
        vehicles = await fetch_vehicles()
        logger.info(f"Fetched {len(vehicles)} vehicles")

        cars = [Car(graph, vehicle_data) for vehicle_data in vehicles]
        animator = Animator(edges)
        logger.info("Cars and animator initialized")

        # Create tasks
        ws_tasks = [asyncio.create_task(send_location(car)) for car in cars]
        animation_task = asyncio.create_task(animator.run_animation(cars))
        logger.info("Tasks created for sending locations and running animation")

        try:
            await asyncio.gather(*ws_tasks, animation_task)
        except KeyboardInterrupt:
            logger.warning("Keyboard interrupt received, closing animator")
            animator.close()
    except Exception as e:
        logger.error(f"Exception during task execution: {str(e)}")
        logger.error(f"Error in main execution: {str(e)}")


if __name__ == "__main__":
    asyncio.run(main())
