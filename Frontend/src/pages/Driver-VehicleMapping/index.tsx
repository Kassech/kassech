import Header from '@/components/header';
import DriverVehicleMap from '@/sections/Driver-Vehicle Mapping';

export default function QueueManagerRegistration() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'DriverVehicleMap', href: '/driverVehicleMap' },
  ];
  return (
    <>
      <Header paths={paths} />
      <div className="w-full">
        <DriverVehicleMap />
      </div>
    </>
  );
}
