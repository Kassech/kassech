import osmnx as ox
import logging

class GraphLoader:
    def __init__(self):
        self.graph = None

    def load_graph(self, location="Addis Ababa, Ethiopia", network_type="drive"):
        try:
            self.graph = ox.graph_from_place(location, network_type=network_type)
            return self.graph
        except Exception as e:
            logging.error(f"Error loading graph: {e}")
            raise
