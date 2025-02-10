import MapComponent from '@/sections/RealTimeMonitoring/MapComponent';
import useWebSocketData from '@/services/useWebSocketData';
import React, { useState } from 'react';

const TrackVehicle: React.FC = () => {
  const { data } = useWebSocketData();
  const [vehicleId, setVehicleId] = useState<number | null>(null);

  const filteredData =
    data?.allCars.filter((car) => car.vehicle_id === vehicleId) || [];

  return (
    <div>
      <h1 className="text-2xl font-bold p-4">Track Vehicle</h1>
      <div className="p-4">
        <input
          type="number"
          placeholder="Enter Vehicle ID"
          className="p-2 border rounded"
          onChange={(e) => setVehicleId(Number(e.target.value))}
        />
      </div>
      <MapComponent data={filteredData} center={[41.0321, 9.7512]} zoom={13} />
    </div>
  );
};

export default TrackVehicle;
