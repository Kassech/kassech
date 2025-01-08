import React, { useEffect } from "react";
import { useMap } from "react-leaflet";
import L from "leaflet";

export const MapView: React.FC<{ stations: any[]; drawRoute: (start: any, end: any) => void }> = ({ stations, drawRoute }) => {
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

  return (
    <div>
      {stations.length >= 2 && (
        <button onClick={() => drawRoute(stations[0], stations[1])}>Draw Route</button>
      )}
    </div>
  );
};
