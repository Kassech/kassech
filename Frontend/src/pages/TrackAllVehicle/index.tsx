import LoadingSpinner from '@/components/loading-spinner';
import { useGetAllVehicles } from '@/services/vehicleService';
import { VehicleMapDashboard } from './mapcomponent';
import Header from '@/components/header';
import { useGetAllPaths } from '@/services/pathService';

export default function TrackAllVehicle() {
  const { data: vehicles, isLoading, isError } = useGetAllVehicles('', 1000);
  const {
    data: path,
    isLoading: isPathLoading,
    isError: isPathError,
  } = useGetAllPaths();
  console.log('🚀 ~ TrackAllVehicle ~ path:', path);
  console.log('🚀 ~ TrackAllVehicle ~ path:', path);

  console.log('🚀 ~ TrackAllVehicle ~ vehicles:', vehicles);
  console.log('🚀 ~ TrackAllVehicle ~ isError:', isError);
  console.log('🚀 ~ TrackAllVehicle ~ isLoading:', isLoading);
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Dashboard', href: '/dashboard' },
  ];
  if (isLoading || isPathLoading) return <LoadingSpinner />;
  return (
    <div>
      <Header paths={paths} />

      <h1 className="text-2xl font-bold p-4">Track All Vehicles</h1>
      <VehicleMapDashboard
        vehicles={vehicles?.data ?? []}
        paths={path?.data ?? []}
      />
    </div>
  );
}
