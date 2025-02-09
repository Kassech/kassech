# visualization/animator.py
import logging
import os
import matplotlib
matplotlib.use('TkAgg')
import matplotlib.pyplot as plt
import contextily as ctx
import asyncio, time, json, websockets
import geopandas as gpd
from shapely.geometry import Point
from matplotlib.offsetbox import OffsetImage, AnnotationBbox
from config.settings import UPDATE_INTERVAL
import numpy as np
import random
import matplotlib.image as mpimg


class Animator:
    def __init__(self, edges, routes_data):
        self.fig, self.ax = plt.subplots(figsize=(10, 10))
        self.edges = edges
        self.routes_data = routes_data
        self.car_markers = []
        self.car_labels = []
        self.car_images = []
        self.passenger_counts = {}
        self.passenger_texts = {}
        self.station_a_projections = {}
        self._is_running = False

        self._setup_base_map()
        self._process_route_data()
        self._plot_station_a_home_icons()

    def _setup_base_map(self):
        """Set up the base map with edges and basemap."""
        self.edges.plot(ax=self.ax, linewidth=1, edgecolor='gray', alpha=0.7, zorder=1)
        ctx.add_basemap(self.ax, source=ctx.providers.CartoDB.Positron, crs=self.edges.crs, alpha=0.8)
        self.ax.set_axis_off()

    def _process_route_data(self):
        """Pre-process station A coordinates and project to map CRS."""
        # Create a GeoDataFrame from the edges to get the CRS
        gdf_edges = gpd.GeoDataFrame(geometry=self.edges.geometry)
        target_crs = gdf_edges.crs

        # Create a list of points from station A coordinates
        points = []
        for route in self.routes_data:
            station_a = route['route']['station_a']
            points.append(Point(station_a['Longitude'], station_a['Latitude']))

        # Create a GeoDataFrame in WGS84 (EPSG:4326)
        gdf = gpd.GeoDataFrame(geometry=points, crs='EPSG:4326')

        # Project to the map's CRS
        gdf = gdf.to_crs(target_crs)

        # Store projected coordinates
        for route, point in zip(self.routes_data, gdf.geometry):
            self.station_a_projections[route['ID']] = (point.x, point.y)

    def _plot_station_a_home_icons(self):
        """Plot home icons for all station A locations."""
        for route_id, (x, y) in self.station_a_projections.items():
            self.ax.plot(
                x, y,
                marker='^',
                markersize=10,
                color='green',
                linestyle='none',
                label='Station A' if route_id == 1 else None,
                zorder=3
            )
        self.ax.legend(loc='upper right')

    async def listen_passenger_updates(self):
        logger = logging.getLogger(__name__)
        ws_url = "ws://localhost:5000/ws/passengers?test_id=1"
        while True:
            try:
                logger.info("Connecting to WebSocket %s", ws_url)
                async with websockets.connect(ws_url) as ws:
                    logger.info("Connected to WebSocket")
                    while True:
                        try:
                            message = await ws.recv()
                            logger.info("Received WebSocket message: %s", message)
                            data = json.loads(message)
                            path_id = data['pathID']
                            count = data['passengerCount']
                            self.passenger_counts[path_id] = count
                            logger.info("Updated passenger count for path %s to %s", path_id, count)
                        except websockets.ConnectionClosedError as e:
                            logger.error("WebSocket connection closed: %s", e)
                            break
                        except Exception as e:
                            logger.error("Unexpected error: %s", e)
                            break
            except Exception as e:
                logger.error("Failed to connect to WebSocket: %s", e)
                await asyncio.sleep(5)  # Wait before retrying

    async def send_random_path_id(self):
        logger = logging.getLogger(__name__)
        ws_url = "ws://localhost:5000/ws/passengers?test_id=1"
        while True:
            try:
                logger.info("Connecting to WebSocket %s", ws_url)
                async with websockets.connect(ws_url) as ws:
                    logger.info("Connected to WebSocket")
                    while True:
                        path_id = random.choice(list(self.station_a_projections.keys()))
                        message = json.dumps({"action": "increment", "pathID": path_id, "amount": 1})
                        await ws.send(message)
                        logger.info("Sent WebSocket message: %s", message)
                        await asyncio.sleep(0.1)  # Send every 1/10 a second
            except Exception as e:
                logger.error("Failed to connect to WebSocket: %s", e)
                await asyncio.sleep(5)  # Wait before retrying

    async def run_animation(self, cars):
        """Main animation loop with WebSocket integration."""
        valid_cars = [car for car in cars if car.has_valid_route]
        self._init_markers(valid_cars)  # Remove the 'await' keyword here

        # Start WebSocket listener and sender
        ws_listener_task = asyncio.create_task(self.listen_passenger_updates())
        ws_sender_task = asyncio.create_task(self.send_random_path_id())

        plt.show(block=False)
        plt.pause(0.1)
        self._is_running = True

        try:
            while self._is_running:
                start = time.monotonic()
                for car in valid_cars:
                    car.update_position()
                self._update_display(valid_cars)
                elapsed = time.monotonic() - start
                await asyncio.sleep(max(UPDATE_INTERVAL - elapsed, 0))
        except asyncio.CancelledError:
            self._is_running = False
        finally:
            ws_listener_task.cancel()
            ws_sender_task.cancel()
            try:
                await ws_listener_task
                await ws_sender_task
            except:
                pass
            plt.close()

    def _init_markers(self, cars):
        """Initialize car markers and labels."""
        logging.info("Initializing markers for %d cars", len(cars))
        self.car_markers = []
        self.car_labels = []
        self.car_images = []
        print(os.path.join(os.path.dirname(__file__), "car.png"))

        car_img = mpimg.imread(os.path.join(os.path.dirname(__file__), "car.png"))  # Load car image
        imagebox = OffsetImage(car_img, zoom=0.05)  # Adjust size

        for car in cars:
            marker, = self.ax.plot([], [], 'o', alpha=0)  # Invisible marker
            self.car_markers.append(marker)
            label = self.ax.text(0, 0, f"{car.driver_name}", fontsize=8, color='black',
                                ha='center', va='top', backgroundcolor='white', alpha=0.7)
            self.car_labels.append(label)

            ab = AnnotationBbox(imagebox, (0, 0), frameon=False, zorder=3)
            self.ax.add_artist(ab)
            self.car_images.append(ab)  # Store image annotations

        self.info_text = self.ax.text(0.05, 0.95, '', transform=self.ax.transAxes,
                                    fontsize=9, color='black', backgroundcolor='white', zorder=4)
    def _update_display(self, cars):
        """Update all visual elements including passenger counts."""
        info = []
        renderer = self.fig.canvas.get_renderer()

        # Update car positions
        for i, (car, marker, label, img) in enumerate(zip(cars, self.car_markers, self.car_labels, self.car_images)):
            if car.current_position:
                x, y = car.current_position
                marker.set_data([x], [y])
                label.set_position((x, y - 0.0002))
                if img:
                    img.xy = (x, y + 0.0005)
                    img.update_positions(renderer)
                info.append(f"Car {i+1}: {car.driver_name} - {car.speed*3.6:.1f} km/h")

        # Update passenger counts
        current_path_ids = set(self.passenger_counts.keys())
        existing_path_ids = set(self.passenger_texts.keys())

        # Remove old passenger texts
        for pid in existing_path_ids - current_path_ids:
            self.passenger_texts[pid].remove()
            del self.passenger_texts[pid]

        # Add/update current passenger texts
        for path_id in current_path_ids:
            count = self.passenger_counts[path_id]
            if path_id not in self.station_a_projections:
                continue

            x, y = self.station_a_projections[path_id]
            if path_id in self.passenger_texts:
                text = self.passenger_texts[path_id]
                text.set_text(str(count))
                text.set_position((x, y))
            else:
                text = self.ax.text(
                    x, y, str(count),
                    fontsize=12, color='red', ha='center', va='bottom',
                    backgroundcolor='white', zorder=4
                )
                self.passenger_texts[path_id] = text

        # Update info text
        self.info_text.set_text('\n'.join(info))
        self.fig.canvas.draw_idle()
        self.fig.canvas.start_event_loop(0.001)

    def close(self):
        """Clean up the animator."""
        self._is_running = False
        plt.close(self.fig)
