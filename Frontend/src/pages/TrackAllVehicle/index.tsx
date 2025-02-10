import MapComponent from '@/sections/RealTimeMonitoring/MapComponent';
import useWebSocketData from '@/services/useWebSocketData';

export default function TrackAllVehicle(){
  const { data } = useWebSocketData();

  return (
    <div>
      <h1 className="text-2xl font-bold p-4">Track All Vehicles</h1>
      <MapComponent
        data={data?.allCars || []}
        center={[41.0321, 9.7512]}
        zoom={13}
      />
    </div>
  );
};


