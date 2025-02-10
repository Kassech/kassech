import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';

interface MapComponentProps {
  data: { lat: number; lon: number; vehicle_id?: number; path_id?: number }[];
  center: [number, number];
  zoom: number;
}

const MapComponent: React.FC<MapComponentProps> = ({ data, center, zoom }) => {
  return (
    <MapContainer
      center={center}
      zoom={zoom}
      style={{ height: '90vh', width: '100%' }}
    >
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {data.map((item, index) => (
        <Marker key={index} position={[item.lat, item.lon]}>
          <Popup>
            {item.vehicle_id && `Vehicle ID: ${item.vehicle_id}`}
            {item.path_id && `Path ID: ${item.path_id}`}
            <br />
            Latitude: {item.lat} <br />
            Longitude: {item.lon}
          </Popup>
        </Marker>
      ))}
    </MapContainer>
  );
};

export default MapComponent;
