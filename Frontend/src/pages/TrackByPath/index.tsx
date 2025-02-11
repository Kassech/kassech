import MapComponent from '@/sections/RealTimeMonitoring/MapComponent';
import PathSelectionSection from '@/sections/RealTimeMonitoring/PathSelectionSection';
import useWebSocketData from '@/services/useWebSocketData';
import React, { useState } from 'react';

const TrackPathPage: React.FC = () => {
  const { data } = useWebSocketData();
  const [pathId, setPathId] = useState<number | null>(null);

  const filteredData =
    data?.allCars.filter((car) => car.path_id === pathId) || [];

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Track by Path</h1>
      <PathSelectionSection
        paths={data?.paths || []}
        onPathSelect={(selectedPathId) => setPathId(selectedPathId)}
      />
      <MapComponent data={filteredData} center={[41.0321, 9.7512]} zoom={13} />
    </div>
  );
};

export default TrackPathPage;
