import os

# OSMnx configuration
CACHE_DIR = os.path.join(os.path.expanduser('~'), '.cache/osmnx')
PLACE_NAME = "Addis Ababa, Ethiopia"
SIMPLIFY_GRAPH = True
NETWORK_TYPE = 'drive'

# Animation settings
CAR_COLORS = ['red', 'blue', 'lime', 'green', 'yellow', 'orange', 'purple', 'pink', 'brown', 'gray', 'black', 'cyan', 'magenta']
FRAMES = 500
INTERVAL = 100
UPDATE_INTERVAL = 0.001
# Performance
MAX_RETRIES = 3
