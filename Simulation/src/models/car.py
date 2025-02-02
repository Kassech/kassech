# models/car.py
import networkx as nx
import numpy as np
import time
from geopy.distance import great_circle
from config.settings import MAX_RETRIES

class Car:
    def __init__(self, graph, vehicle_data):
        self.graph = graph
        self.vehicle_id = vehicle_data["ID"]
        self.color = vehicle_data["Color"].lower()
        self.speed = 7  # m/s
        self.current_position = None
        self.route = []
        self._route_segments = []
        self._current_segment = 0
        self._segment_progress = 0.0  # Progress within current segment (0-1)
        self.forward = True
        self._last_update = time.monotonic()
        self._initialize_route()

    def _initialize_route(self):
        """Generate smooth route with segment vectors"""
        nodes = list(self.graph.nodes())
        for _ in range(MAX_RETRIES):
            start, end = np.random.choice(nodes, 2, False)
            try:
                path = nx.astar_path(self.graph, start, end,
                                   self._heuristic, 'length')
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
        """Create enhanced route segments with precomputed vectors"""
        self._route_segments = []
        for i in range(len(path)-1):
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
        """Optimized distance calculation"""
        u_pos = self.graph.nodes[u]['x'], self.graph.nodes[u]['y']
        v_pos = self.graph.nodes[v]['x'], self.graph.nodes[v]['y']
        return great_circle(u_pos[::-1], v_pos[::-1]).meters

    def update_position(self):
        """Smooth movement with proper segment progress and detailed logging"""
        if not self._route_segments:
            print("No route segments available.")
            return

        now = time.monotonic()
        delta_time = now - self._last_update
        self._last_update = now
        distance_moved = self.speed * delta_time
        print(f"Time elapsed: {delta_time:.2f}s, Distance to move: {distance_moved:.2f}m")

        while distance_moved > 0:
            segment = self._route_segments[self._current_segment]
            remaining = segment['length'] * (1 - self._segment_progress)
            print(f"Current segment: {self._current_segment}, Remaining distance: {remaining:.2f}m")

            if remaining <= distance_moved:
                # Move to next segment
                distance_moved -= remaining
                self._segment_progress = 0.0
                print(f"Completed segment: {self._current_segment}, Moving to next. Remaining distance to move: {distance_moved:.2f}m")

                if self.forward:
                    self._current_segment += 1
                    if self._current_segment >= len(self._route_segments):
                        # Reverse direction at end of route
                        self.forward = False
                        self._current_segment = len(self._route_segments) - 1
                        print("Reached end of route, reversing direction.")
                else:
                    self._current_segment -= 1
                    if self._current_segment < 0:
                        # Reverse direction at start of route
                        self.forward = True
                        self._current_segment = 0
                        print("Reached start of route, reversing direction.")
            else:
                # Move within current segment
                self._segment_progress += distance_moved / segment['length']
                distance_moved = 0
                print(f"Moving within segment: {self._current_segment}, Updated segment progress: {self._segment_progress:.2f}")

            # Update current position
            segment = self._route_segments[self._current_segment]
            progress = self._segment_progress if self.forward else (1 - self._segment_progress)
            self.current_position = (
                segment['start'][0] + progress * segment['dx'],
                segment['start'][1] + progress * segment['dy']
            )
            print(f"Updated position: {self.current_position}")

        return self.current_position

    @property
    def has_valid_route(self):
        return len(self._route_segments) > 0
