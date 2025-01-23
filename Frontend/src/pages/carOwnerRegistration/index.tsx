// pages/driver/Driver.jsx

import Header from '@/components/header';
import Owner from '@/sections/Owner';

export default function OwnerRegistration() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Driver', href: '/driver' },
  ];
  return (
    <>
      <Header paths={paths} />
      <div className="w-full">
        <Owner />
      </div>
    </>
  );
}
