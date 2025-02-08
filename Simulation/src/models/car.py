# models/car.py
import networkx as nx, numpy as np, time, logging
from geopy.distance import great_circle
from config.settings import MAX_RETRIES
import osmnx as ox

logger = logging.getLogger(__name__)

class Car:
    def __init__(self, graph, vehicle_data):
        self.graph = graph
        self.vehicle_id = vehicle_data["ID"]
        self.color = vehicle_data["Color"].lower()
        self.driver_name = f"{vehicle_data['Driver']['User']['FirstName']}"
        # self.driver_picture = vehicle_data["Driver"]["User"]["ProfilePicture"]
        self.driver_picture = "http://localhost:5000/uploads/testing.jpg"
        self.speed = 7  # m/s
        self.current_position = None
        self.route = []
        self._route_segments = []
        self._current_segment = 0
        self.destination_mode = False
        self._segment_progress = 0.0
        self.forward = True
        self._last_update = time.monotonic()
        self._initialize_route()

    def _initialize_route(self):
        # Generate a random route with precomputed segments for smooth movement
        nodes = list(self.graph.nodes())
        for _ in range(MAX_RETRIES):
            start, end = np.random.choice(nodes, 2, replace=False)
            try:
                path = nx.astar_path(self.graph, start, end, self._heuristic, weight='length')
                self.route = path
                self._create_route_segments(path)
                if self._route_segments:
                    self.current_position = self._route_segments[0]['start']
                return
            except (nx.NetworkXNoPath, KeyError):
                continue
        self.route = []
        self._route_segments = []

    def _create_route_segments(self, path):
        # Precompute segments with direction vectors and lengths
        self._route_segments = []
        for i in range(len(path) - 1):
            u, v = path[i], path[i+1]
            start = (self.graph.nodes[u]['x'], self.graph.nodes[u]['y'])
            end = (self.graph.nodes[v]['x'], self.graph.nodes[v]['y'])
            dx = end[0] - start[0]
            dy = end[1] - start[1]
            length = self.graph[u][v][0]['length']
            self._route_segments.append({
                'start': start,
                'end': end,
                'dx': dx,
                'dy': dy,
                'length': length
            })

    def _heuristic(self, u, v):
        # Estimate distance using great circle distance
        u_pos = (self.graph.nodes[u]['x'], self.graph.nodes[u]['y'])
        v_pos = (self.graph.nodes[v]['x'], self.graph.nodes[v]['y'])
        return great_circle(u_pos[::-1], v_pos[::-1]).meters

    def set_destination(self, station_a, station_b):
        # Set a new destination route based on station coordinates
        try:
            nearest_a = ox.distance.nearest_nodes(self.graph, station_a["Longitude"], station_a["Latitude"])
            nearest_b = ox.distance.nearest_nodes(self.graph, station_b["Longitude"], station_b["Latitude"])
            path = nx.astar_path(self.graph, nearest_a, nearest_b, heuristic=self._heuristic, weight='length')
            self.route = path
            self._create_route_segments(path)
            self._current_segment = 0
            self._segment_progress = 0.0
            self.forward = True
            self.destination_mode = True
            if self._route_segments:
                self.current_position = self._route_segments[0]['start']
        except Exception as e:
            logger.error("Invalid destination path: %s; reverting to random movement", e)
            self.destination_mode = False

    def update_position(self):
        # Update the car's position along its route based on elapsed time and speed
        if not self._route_segments:
            logger.debug("No route segments available for car %s.", self.vehicle_id)
            return self.current_position

        now = time.monotonic()
        delta_time = now - self._last_update
        self._last_update = now
        distance_moved = self.speed * delta_time

        # Stop if destination reached
        if (self.destination_mode and self.forward and
            self._current_segment == len(self._route_segments) - 1 and
            self._segment_progress >= 1):
            self.current_position = self._route_segments[-1]['end']
            return self.current_position

        while distance_moved > 0:
            segment = self._route_segments[self._current_segment]
            remaining = segment['length'] * (1 - self._segment_progress)
            if remaining <= distance_moved:
                distance_moved -= remaining
                self._segment_progress = 0.0
                if self.forward:
                    self._current_segment += 1
                    if self._current_segment >= len(self._route_segments):
                        if self.destination_mode:
                            self.current_position = self._route_segments[-1]['end']
                            return self.current_position
                        self.forward = False
                        self._current_segment = len(self._route_segments) - 1
                else:
                    self._current_segment -= 1
                    if self._current_segment < 0:
                        self.forward = True
                        self._current_segment = 0
            else:
                self._segment_progress += distance_moved / segment['length']
                distance_moved = 0

            segment = self._route_segments[self._current_segment]
            progress = self._segment_progress if self.forward else (1 - self._segment_progress)
            self.current_position = (
                segment['start'][0] + progress * segment['dx'],
                segment['start'][1] + progress * segment['dy']
            )
        return self.current_position

    @property
    def has_valid_route(self):
        return bool(self._route_segments)
