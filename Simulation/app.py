import osmnx as ox
import networkx as nx
import matplotlib.pyplot as plt
import matplotlib.animation as animation
import numpy as np
import contextily as ctx  # For realistic basemap
from geopy.distance import great_circle

# Fetch OSM data for a specific location
place_name = "Dire Dawa, Ethiopia"
graph = ox.graph_from_place(place_name, network_type='drive', simplify=False)
graph = ox.utils_graph.convert.to_undirected(graph)

# Convert graph to GeoDataFrames
nodes, edges = ox.graph_to_gdfs(graph)

# Create a figure with a real map background
fig, ax = plt.subplots(figsize=(10, 10))

# Plot roads (edges) with realistic color
edges.plot(ax=ax, linewidth=1, edgecolor="gray", alpha=0.7, zorder=1)

# Add a realistic basemap (satellite or terrain)
ctx.add_basemap(ax, source=ctx.providers.CartoDB.Positron, crs=edges.crs)  # Street map

# Car class
class Car:
    def __init__(self, graph, start_node, end_node, color):
        self.graph = graph
        self.route = ox.shortest_path(graph, start_node, end_node)
        self.current_index = 0
        self.color = color  # Assign color for the car
        self.speed = 13.89  # ~50 km/h in m/s

        if self.route:
            self.coords = [(graph.nodes[n]['x'], graph.nodes[n]['y']) for n in self.route]
        else:
            self.coords = []

    def update_position(self):
        if self.current_index < len(self.coords) - 1:
            self.current_index += 1
            return self.coords[self.current_index]
        return None

# Create multiple cars with distinct colors
colors = ['red', 'blue', 'lime']
cars = []
for i in range(3000):
    nodes_list = list(graph.nodes())
    start = np.random.choice(nodes_list)
    end = np.random.choice(nodes_list)
    cars.append(Car(graph, start, end, colors[i % len(colors)]))

# Create plot elements
car_icons = [ax.plot([], [], 'o', color=car.color, markersize=8, alpha=0.9, zorder=3)[0] for car in cars]
route_lines = [ax.plot([], [], color=car.color, linewidth=2, alpha=0.8, zorder=2)[0] for car in cars]
info_text = ax.text(0.05, 0.95, '', transform=ax.transAxes, fontsize=10, color='black', backgroundcolor='white')

# Animation initialization
def init():
    for icon, line in zip(car_icons, route_lines):
        icon.set_data([], [])
        line.set_data([], [])
    info_text.set_text('')
    return car_icons + route_lines + [info_text]

# Update function
def update(frame):
    info = []
    for i, car in enumerate(cars):
        pos = car.update_position()
        if pos:  # Ensure pos is not None
            x, y = pos
            car_icons[i].set_data([x], [y])  # Ensure x, y are sequences

            if len(car.coords) > 1:
                xs, ys = zip(*car.coords[:car.current_index])
                route_lines[i].set_data(xs, ys)

            info.append(f"Car {i+1}: {car.speed*3.6:.1f} km/h")

    info_text.set_text('\n'.join(info))
    return car_icons + route_lines + [info_text]

# Run animation
ani = animation.FuncAnimation(fig, update, init_func=init,
                              frames=200, interval=100, blit=True, repeat=False)

plt.show()
