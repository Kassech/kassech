import React, { useState } from "react";
import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  useMapEvents,
} from "react-leaflet";
import "leaflet/dist/leaflet.css";
import * as L from "leaflet";

const SearchMap: React.FC = () => {
  const [position, setPosition] = useState<L.LatLng | null>(null);

  // Handle click event on the map
  const MapClickHandler = () => {
    const map = useMapEvents({
      click(event: L.LeafletMouseEvent) {
        const { lat, lng } = event.latlng;
        setPosition(new L.LatLng(lat, lng));
      },
    });

    return null;
  };

  return (
    <MapContainer
      center={[51.505, -0.09]} // Default center for the map
      zoom={13} // Initial zoom level
      style={{ height: "100vh", width: "100%" }} // Full-screen map
    >
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />

      {/* Handle map click event */}
      <MapClickHandler />

      {position && (
        <Marker position={position}>
          <Popup>
            <div>
              <p>Latitude: {position.lat.toFixed(4)}</p>
              <p>Longitude: {position.lng.toFixed(4)}</p>
            </div>
          </Popup>
        </Marker>
      )}
    </MapContainer>
  );
};

export default SearchMap;
