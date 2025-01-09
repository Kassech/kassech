// pages/driver/Driver.jsx

import Header from '@/components/header';
import Driver from '@/sections/driver';

export default function DriverPage() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Driver', href: '/driver' },
  ];
  return (
    <>
      <Header paths={paths} />
      <div className="flex flex-1 flex-col gap-4 p-2 pt-0">
        <div className="grid auto-rows-min gap-4 md:grid-cols-3 overflow-x-hidden">
          <Driver />
        </div>
      </div>
    </>
  );
}
