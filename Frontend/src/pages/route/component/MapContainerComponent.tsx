import React from 'react';
import { MapContainer, TileLayer, Marker } from 'react-leaflet';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-routing-machine';
import { MapView } from './MapViewComponent';
import { generateRandomColor } from '@/utils/colorUtils';
import { Station } from '@/types/api';

const MapContainerComponent: React.FC<{ stations: Station[] }> = ({
  stations,
}) => {
  const stationMarkers = stations.map((station: Station) => {
    const randomColor = generateRandomColor();
    const customIcon = L.divIcon({
      className: 'custom-icon',
      html: `<div style="position: relative; text-align: center;">
          <span
            style="
              position: absolute;
              bottom: 30px;
              left: 50%;
              transform: translateX(-50%);
              background-color: rgba(0, 0, 0, 0.7);
              color: white;
              padding: 2px 5px;
              border-radius: 3px;
              font-size: 12px;">
            ${station.LocationName}
          </span>
          <div style="background-color: ${randomColor}; width: 20px; height: 20px; border-radius: 50%;"></div>
        </div>`,
      iconSize: [20, 20],
    });

    return (
      <Marker
        key={station.ID}
        position={[station.Latitude, station.Longitude]}
        icon={customIcon}
      />
    );
  });

  return (
    <MapContainer
      center={[0, 0]}
      zoom={2}
      style={{ height: '100%', width: '100%' }}
    >
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      <MapView stations={stations} />
      {stationMarkers}
    </MapContainer>
  );
};

export default MapContainerComponent;
