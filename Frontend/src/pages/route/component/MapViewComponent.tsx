import React, { useEffect } from 'react';
import { useMap } from 'react-leaflet';
import L from 'leaflet';
import { Station } from '@/types/api';

export const MapView: React.FC<{ stations: Station[] }> = ({ stations }) => {
  const map = useMap();

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
            13
          );
        },
        () => {
          map.setView([0, 0], 2);
        }
      );
    }
  }, [stations, map]);

  L.Routing.control({
    waypoints: [L.latLng(57.74, 11.94), L.latLng(57.6792, 11.949)],
  }).addTo(map);

  return null;
};
