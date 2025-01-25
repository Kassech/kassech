import React, { useEffect, useRef } from 'react';
import { useMap } from 'react-leaflet';
import L from 'leaflet';
import { Station } from '@/types/api';
import useFormStore from '@/store/routemap';

export const MapView: React.FC<{ stations: Station[] }> = ({ stations }) => {
  const map = useMap();
  const { location1, location2 } = useFormStore();
  const routeControlRef = useRef<L.Routing.Control | null>(null);

  useEffect(() => {
    if (stations.length > 0) {
      const bounds = L.latLngBounds(
        stations.map((station) => [station.Latitude, station.Longitude])
      );
      map.fitBounds(bounds, { padding: [50, 50] });
    } else {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          map.setView(
            [position.coords.latitude, position.coords.longitude],
            15
          );
        },
        () => {
          map.setView([0, 0], 2);
        }
      );
    }
  }, [stations, map]);

  useEffect(() => {
    // Clear previous route when locations change
    if (routeControlRef.current) {
      routeControlRef.current.remove();
    }

    if (location1 && location2) {
      const waypoint1 = stations.find((station) => {
        return station.ID.toString() === location1;
      });
      const waypoint2 = stations.find((station) => {
        return station.ID.toString() === location2;
      });

      if (waypoint1 && waypoint2) {
        // Create new route control and add it to the map
        const newRouteControl = L.Routing.control({
          waypoints: [
            L.latLng(waypoint1.Latitude, waypoint1.Longitude),
            L.latLng(waypoint2.Latitude, waypoint2.Longitude),
          ],
        }).addTo(map);

        // Store the route control in the ref to remove it later if needed
        routeControlRef.current = newRouteControl;
      }
    }
  }, [location1, location2, stations, map]);

  return null;
};
