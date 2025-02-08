// pages/Dashboard.jsx
import Header from '@/components/header';
import VehicleDelegation from '@/sections/VehicleDelegation';

export default function Delegation() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Delegation', href: '/delegation' },
  ];
  return (
    <>
      <Header paths={paths} />
      <VehicleDelegation />
    </>
  );
}
