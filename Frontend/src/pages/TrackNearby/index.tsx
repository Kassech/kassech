import MapComponent from '@/sections/RealTimeMonitoring/MapComponent';
import useWebSocketData from '@/services/useWebSocketData';
import React, { useState } from 'react';

const TrackNearbyPage: React.FC = () => {
  const { data } = useWebSocketData();
  const [lat, setLat] = useState<number>(41.0321);
  const [lon, setLon] = useState<number>(9.7512);
  const [radius, setRadius] = useState<number>(500000000);

  const filteredData = data?.nearbyCars || [];

  return (
    <div>
      <h1 className="text-2xl font-bold p-4">Track Nearby Vehicles</h1>
      <div className="p-4 space-y-2">
        <input
          type="number"
          placeholder="Latitude"
          className="p-2 border rounded"
          value={lat}
          onChange={(e) => setLat(Number(e.target.value))}
        />
        <input
          type="number"
          placeholder="Longitude"
          className="p-2 border rounded"
          value={lon}
          onChange={(e) => setLon(Number(e.target.value))}
        />
        <input
          type="number"
          placeholder="Radius"
          className="p-2 border rounded"
          value={radius}
          onChange={(e) => setRadius(Number(e.target.value))}
        />
      </div>
      <MapComponent data={filteredData} center={[lat, lon]} zoom={13} />
    </div>
  );
};

export default TrackNearbyPage;
