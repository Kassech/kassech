# visualization/animator.py
import matplotlib
matplotlib.use('TkAgg')  # Must come before pyplot import
import matplotlib.pyplot as plt
import contextily as ctx
import asyncio
import time
from config.settings import UPDATE_INTERVAL

class Animator:
    def __init__(self, edges):
        self.fig, self.ax = plt.subplots(figsize=(10, 10))
        self.edges = edges
        self.car_markers = []
        self.info_text = None
        self._setup_base_map()
        self._is_running = False

    def _setup_base_map(self):
        """Initialize static map elements"""
        self.edges.plot(ax=self.ax, linewidth=1, edgecolor='gray', alpha=0.7, zorder=1)
        ctx.add_basemap(self.ax, source=ctx.providers.CartoDB.Positron,
                      crs=self.edges.crs, alpha=0.8)
        self.ax.set_axis_off()

    async def run_animation(self, cars):
        """Main animation loop with position updates"""
        valid_cars = [car for car in cars if car.has_valid_route]
        self._init_markers(valid_cars)

        # Start non-blocking display
        plt.show(block=False)
        plt.pause(0.1)  # Initial render

        self._is_running = True
        try:
            while self._is_running:
                frame_start = time.monotonic()

                # Update all car positions first
                for car in valid_cars:
                    car.update_position()

                # Then update visualization
                self._update_display(valid_cars)

                # Maintain frame rate
                elapsed = time.monotonic() - frame_start
                await asyncio.sleep(max(UPDATE_INTERVAL - elapsed, 0))

        except asyncio.CancelledError:
            self._is_running = False
        finally:
            plt.close()

    def _init_markers(self, cars):
        """Initialize car markers and info text"""
        self.car_markers = [
            self.ax.plot([], [], 'o', color=car.color,
                       markersize=10, alpha=0.9, zorder=3)[0]
            for car in cars
        ]
        self.info_text = self.ax.text(
            0.05, 0.95, '', transform=self.ax.transAxes,
            fontsize=9, color='black', backgroundcolor='white', zorder=4
        )

    def _update_display(self, cars):
        """Update visualization elements"""
        info = []
        for i, (car, marker) in enumerate(zip(cars, self.car_markers)):
            if car.current_position:
                marker.set_data([car.current_position[0]], [car.current_position[1]])
                info.append(f"Car {i+1}: {car.speed*3.6:.1f} km/h")

        self.info_text.set_text('\n'.join(info))

        # Process GUI events
        self.fig.canvas.draw_idle()
        self.fig.canvas.start_event_loop(0.001)

    def close(self):
        """Properly close the animation"""
        self._is_running = False
        plt.close(self.fig)
