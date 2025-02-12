import Header from '@/components/header';
import Owner from '@/sections/owner';

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
