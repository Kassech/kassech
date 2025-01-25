import { MapContainer, TileLayer, Marker, Popup, useMap } from 'react-leaflet';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import { useEffect } from 'react';
import { useRouteStore } from '@/store/pathStore';

L.Marker.prototype.options.icon = L.icon({
  iconUrl:
    'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon.png',
  iconSize: [25, 41],
  iconAnchor: [12, 41],
});

function MapController() {
  const map = useMap();
  const { selectedRoute } = useRouteStore();

  useEffect(() => {
    if (selectedRoute) {
      const bounds = L.latLngBounds([
        [selectedRoute.start.lat, selectedRoute.start.lng],
        [selectedRoute.end.lat, selectedRoute.end.lng],
      ]);
      map.flyToBounds(bounds, { padding: [50, 50] });

      // Calculate distance and time
      const distance = calculateDistance(
        selectedRoute.start.lat,
        selectedRoute.start.lng,
        selectedRoute.end.lat,
        selectedRoute.end.lng
      );
      const time = calculateTime(distance);
      useRouteStore.getState().setCalculations(distance, time);
    }
  }, [selectedRoute, map]);

  return null;
}

export default function RouteMap() {
  const { selectedRoute } = useRouteStore();

  return (
    <div className="h-full w-full rounded-lg overflow-hidden">
      <MapContainer
        center={[51.505, -0.09]}
        zoom={3}
        style={{ height: '100%', width: '100%' }}
        touchZoom={false}
        dragging={!L.Browser.mobile}
        doubleClickZoom={false}
        scrollWheelZoom={true}
        zoomControl={true}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />

        {selectedRoute && (
          <>
            <Marker
              position={[selectedRoute.start.lat, selectedRoute.start.lng]}
            >
              <Popup>{selectedRoute.start.name}</Popup>
            </Marker>
            <Marker position={[selectedRoute.end.lat, selectedRoute.end.lng]}>
              <Popup>{selectedRoute.end.name}</Popup>
            </Marker>
          </>
        )}

        <MapController />
      </MapContainer>
    </div>
  );
}

// Haversine formula for distance calculation
function calculateDistance(
  lat1: number,
  lon1: number,
  lat2: number,
  lon2: number
) {
  const R = 6371; // Earth radius in km
  const dLat = ((lat2 - lat1) * Math.PI) / 180;
  const dLon = ((lon2 - lon1) * Math.PI) / 180;
  const a =
    Math.sin(dLat / 2) * Math.sin(dLat / 2) +
    Math.cos((lat1 * Math.PI) / 180) *
      Math.cos((lat2 * Math.PI) / 180) *
      Math.sin(dLon / 2) *
      Math.sin(dLon / 2);
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
  return Number((R * c).toFixed(2));
}

function calculateTime(distance: number) {
  const averageSpeed = 50; // km/h
  const hours = distance / averageSpeed;
  const minutes = Math.round(hours * 60);
  return `${minutes}m`;
}
