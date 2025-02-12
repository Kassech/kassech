import networkx as nx
import numpy as np
import time
import logging
from geopy.distance import great_circle
import osmnx as ox

logger = logging.getLogger(__name__)

class Car:
    def __init__(self, graph, vehicle_data):
        self.graph = graph
        self.vehicle_id = vehicle_data["ID"]
        self.driver_id = vehicle_data["Driver"]["User"]["ID"]
        self.driver_name = vehicle_data["Driver"]["User"]["FirstName"]
        self.color = vehicle_data["Color"].lower()
        self.speed = 20  # m/s
        self.current_position = None
        self._route_segments = []
        self._current_segment = 0
        self._segment_progress = 0.0
        self.forward = True
        self.arrived = False
        self.destination_mode = False
        self._last_update = time.monotonic()
        self._initialize_route()

    def _initialize_route(self):
        nodes = list(self.graph.nodes())
        for _ in range(3):  # Reduced retries
            try:
                start, end = np.random.choice(nodes, 2, replace=False)
                path = nx.astar_path(self.graph, start, end, self._heuristic)
                self._create_route_segments(path)
                if self._route_segments:
                    self.current_position = self._route_segments[0]['start']
                return
            except Exception:
                continue
        self._route_segments = []

    def _create_route_segments(self, path):
        self._route_segments = []
        for i in range(len(path)-1):
            u, v = path[i], path[i+1]
            start = (self.graph.nodes[u]['x'], self.graph.nodes[u]['y'])
            end = (self.graph.nodes[v]['x'], self.graph.nodes[v]['y'])
            self._route_segments.append({
                'start': start,
                'end': end,
                'dx': end[0] - start[0],
                'dy': end[1] - start[1],
                'length': self.graph[u][v][0]['length']
            })

    def set_destination(self, station_a, station_b, path_id):
        try:
            start = ox.distance.nearest_nodes(self.graph, station_a["Longitude"], station_a["Latitude"])
            end = ox.distance.nearest_nodes(self.graph, station_b["Longitude"], station_b["Latitude"])
            path = nx.astar_path(self.graph, start, end, self._heuristic)
            self._create_route_segments(path)
            self._current_segment = 0
            self._segment_progress = 0.0
            self.forward = True
            self.destination_mode = True
            self.arrived = False
            self.path_id = path_id
        except Exception as e:
            logger.error(f"Route error: {e}")
            self._route_segments = []

    def update_position(self):
        if not self._route_segments:
            return

        now = time.monotonic()
        delta_time = now - self._last_update
        self._last_update = now
        distance_moved = self.speed * delta_time

        if self._should_stop():
            self.current_position = self._route_segments[-1]['end']
            self.arrived = True
            self._route_segments = []
            return

        while distance_moved > 0 and self._route_segments:
            self._current_segment = max(0, min(self._current_segment, len(self._route_segments)-1))
            segment = self._route_segments[self._current_segment]

            remaining = segment['length'] * (1 - self._segment_progress)
            if remaining <= distance_moved:
                distance_moved -= remaining
                self._advance_segment()
            else:
                self._segment_progress += distance_moved / segment['length']
                distance_moved = 0

            self._update_current_position()

    def _should_stop(self):
        return (self.destination_mode
                and self.forward
                and self._current_segment >= len(self._route_segments)-1
                and self._segment_progress >= 1)

    def _advance_segment(self):
        self._segment_progress = 0.0
        if self.forward:
            if self._current_segment < len(self._route_segments)-1:
                self._current_segment += 1
            else:
                self.forward = False
        else:
            if self._current_segment > 0:
                self._current_segment -= 1
            else:
                self.forward = True

    def _update_current_position(self):
        segment = self._route_segments[self._current_segment]
        progress = self._segment_progress if self.forward else (1 - self._segment_progress)
        self.current_position = (
            segment['start'][0] + progress * segment['dx'],
            segment['start'][1] + progress * segment['dy']
        )

    def reset_for_new_route(self):
        self._route_segments = []
        self.current_position = None
        self.arrived = False

    @property
    def has_valid_route(self):
        return bool(self._route_segments)

    def _heuristic(self, u, v):
        u_pos = (self.graph.nodes[u]['x'], self.graph.nodes[u]['y'])
        v_pos = (self.graph.nodes[v]['x'], self.graph.nodes[v]['y'])
        return great_circle(u_pos[::-1], v_pos[::-1]).meters
