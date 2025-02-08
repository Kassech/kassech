# services/graph_loader.py
import logging
import os
import hashlib
import osmnx as ox
from config.settings import CACHE_DIR, PLACE_NAME, SIMPLIFY_GRAPH, NETWORK_TYPE

class GraphLoader:
    def __init__(self):
        ox.settings.cache_folder = CACHE_DIR
        ox.settings.use_cache = True
        ox.settings.log_console = False  # Reduce console logging
        ox.settings.requests_timeout = 120  # Increase timeout for API requests

    def _get_cache_filename(self):
        """Generate unique filename based on graph parameters"""
        params = f"{PLACE_NAME}-{NETWORK_TYPE}-{SIMPLIFY_GRAPH}"
        hash_key = hashlib.md5(params.encode()).hexdigest()
        return os.path.join(CACHE_DIR, f"graph_{hash_key}.graphml")

    def load_graph(self):
        """Load and optimize graph with persistent caching"""
        cache_file = self._get_cache_filename()

        try:
            if os.path.exists(cache_file):
                # Load pre-processed graph from cache
                return ox.load_graphml(cache_file)

            # Generate and cache new graph
            graph = ox.graph_from_place(
                PLACE_NAME,
                network_type=NETWORK_TYPE,
                simplify=SIMPLIFY_GRAPH,  # Simplify only once here
                retain_all=False          # Don't keep disconnected components
            )

            # Convert to undirected and save
            undirected_graph = ox.utils_graph.convert.to_undirected(graph)
            ox.save_graphml(undirected_graph, cache_file)
            return undirected_graph

        except Exception as e:
            raise RuntimeError(f"Failed to load graph: {str(e)}")

    def warm_cache(self):
        """Pre-generate and cache the graph if not exists"""
        if not os.path.exists(self._get_cache_filename()):
            try:
                self.load_graph()
            except Exception as e:
                logging.getLogger(__name__).warning(f"Cache warming failed: {e}")
