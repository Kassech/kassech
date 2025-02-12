import json
import logging
import os
import random
import matplotlib
import websockets
matplotlib.use('TkAgg')
import matplotlib.pyplot as plt
import contextily as ctx
import asyncio
from matplotlib.offsetbox import OffsetImage, AnnotationBbox
import matplotlib.image as mpimg
import geopandas as gpd
from shapely.geometry import Point

class Animator:
    def __init__(self, edges, routes_data):
        self.fig, self.ax = plt.subplots(figsize=(10, 10))
        self.edges = edges
        self.routes_data = routes_data
        self.car_data = {}
        self.station_a_projections = {}
        self.passenger_counts = {}
        self._is_running = False
        self._setup_base_map()
        self._process_route_data()
        self._plot_station_a_home_icons()
        self.passenger_counts = {}
        self.passenger_texts = {}
        self._add_passenger_websocket_tasks()

    def _add_passenger_websocket_tasks(self):
        self.passenger_tasks = [
            asyncio.create_task(self.listen_passenger_updates()),
            asyncio.create_task(self.send_random_path_id())
        ]

    def _setup_base_map(self):
        self.edges.plot(ax=self.ax, linewidth=1, edgecolor='gray', alpha=0.7, zorder=1)
        ctx.add_basemap(self.ax, source=ctx.providers.CartoDB.Positron, crs=self.edges.crs, alpha=0.8)
        self.ax.set_axis_off()

    def _process_route_data(self):
        gdf_edges = gpd.GeoDataFrame(geometry=self.edges.geometry)
        target_crs = gdf_edges.crs
        points = [Point(r['route']['station_a']['Longitude'], r['route']['station_a']['Latitude'])
                 for r in self.routes_data]
        gdf = gpd.GeoDataFrame(geometry=points, crs='EPSG:4326').to_crs(target_crs)
        self.station_a_projections = {r['ID']: (p.x, p.y) for r, p in zip(self.routes_data, gdf.geometry)}

    def _plot_station_a_home_icons(self):
        for x, y in self.station_a_projections.values():
            self.ax.plot(x, y, marker='^', markersize=10, color='green', linestyle='none', zorder=3)
        self.ax.legend(['Station A'], loc='upper right')

    def _init_markers(self, cars):
        car_img = mpimg.imread(os.path.join(os.path.dirname(__file__), "car.png"))
        imagebox = OffsetImage(car_img, zoom=0.05)

        for car in cars:
            marker, = self.ax.plot([], [], 'o', alpha=0)
            label = self.ax.text(0, 0, f"{car.driver_name}", fontsize=8,
                               ha='center', va='top', backgroundcolor='white', alpha=0.7)
            ab = AnnotationBbox(imagebox, (0,0), frameon=False, zorder=3)
            self.ax.add_artist(ab)

            self.car_data[car.vehicle_id] = {
                "marker": marker,
                "label": label,
                "image": ab,
                "visible": True
            }

    async def run_animation(self, cars):
        self._init_markers(cars)
        plt.show(block=False)
        plt.pause(0.1)
        self._is_running = True

        try:
            while self._is_running:
                valid_cars = [car for car in cars if car.has_valid_route]
                for car in valid_cars:
                    car.update_position()
                self._update_display(valid_cars)
                await asyncio.sleep(0.05)
        except asyncio.CancelledError:
            self._is_running = False
        finally:
            plt.close()
            for task in self.passenger_tasks:
                task.cancel()
            try:
                await asyncio.gather(*self.passenger_tasks)
            except:
                pass


    async def listen_passenger_updates(self):
        logger = logging.getLogger(__name__)
        ws_url = "ws://localhost:5000/ws/passengers?test_id=1"
        while True:
            try:
                async with websockets.connect(
                    ws_url,
                    ping_interval=25,
                    ping_timeout=10,
                    close_timeout=15
                ) as ws:
                    logger.info("Connected to passenger WebSocket")
                    while True:
                        message = await ws.recv()
                        data = json.loads(message)
                        if 'pathID' in data and 'passengerCount' in data:
                            self.passenger_counts[data['pathID']] = data['passengerCount']
            except Exception as e:
                logger.error(f"Passenger WS error: {e}")
                await asyncio.sleep(5)

    async def send_random_path_id(self):
        logger = logging.getLogger(__name__)
        ws_url = "ws://localhost:5000/ws/passengers?test_id=1"
        while True:
            try:
                async with websockets.connect(ws_url) as ws:
                    while True:
                        if self.station_a_projections:
                            path_id = random.choice(list(self.station_a_projections.keys()))
                            message = json.dumps({
                                "action": "increment",
                                "pathID": path_id,
                                "amount": 1
                            })
                            await ws.send(message)
                        await asyncio.sleep(0.5)
            except Exception as e:
                logger.error(f"Random path WS error: {e}")
                await asyncio.sleep(2)


    def _update_display(self, valid_cars):
        valid_ids = {car.vehicle_id for car in valid_cars}

        for vid, data in self.car_data.items():
            if vid not in valid_ids:
                data["marker"].set_alpha(0)
                data["label"].set_text("")
                data["image"].set_visible(False)
                continue

            car = next(c for c in valid_cars if c.vehicle_id == vid)
            if car.current_position:
                x, y = car.current_position
                data["marker"].set_data([x], [y])
                data["label"].set_position((x, y - 0.0002))
                data["image"].xy = (x, y + 0.0005)
                data["marker"].set_alpha(1)
                data["image"].set_visible(True)

        self.fig.canvas.draw_idle()
        self.fig.canvas.start_event_loop(0.001)
        self._update_passenger_counts()

    def _update_passenger_counts(self):
        current_ids = set(self.passenger_counts.keys())
        existing_ids = set(self.passenger_texts.keys())

        # Remove old texts
        for pid in existing_ids - current_ids:
            self.passenger_texts[pid].remove()
            del self.passenger_texts[pid]

        # Add/update texts
        for pid, count in self.passenger_counts.items():
            if pid not in self.station_a_projections:
                continue

            x, y = self.station_a_projections[pid]
            if pid in self.passenger_texts:
                self.passenger_texts[pid].set_text(str(count))
            else:
                self.passenger_texts[pid] = self.ax.text(
                    x, y, str(count),
                    fontsize=12, color='red', ha='center', va='bottom',
                    backgroundcolor='white', zorder=4
                )


    def close(self):
        self._is_running = False
        plt.close(self.fig)
