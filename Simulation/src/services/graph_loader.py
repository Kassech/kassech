import osmnx as ox
from config.settings import CACHE_DIR, PLACE_NAME, SIMPLIFY_GRAPH, NETWORK_TYPE

class GraphLoader:
    def __init__(self):
        ox.settings.cache_folder = CACHE_DIR
        ox.settings.use_cache = True
        ox.settings.log_console = True

    def load_graph(self):
        """Load and optimize graph with caching"""
        try:
            graph = ox.graph_from_place(
                PLACE_NAME,
                network_type=NETWORK_TYPE,
                simplify=SIMPLIFY_GRAPH
            )
            return ox.utils_graph.convert.to_undirected(graph)
        except Exception as e:
            raise RuntimeError(f"Failed to load graph: {str(e)}")
