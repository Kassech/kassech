import Header from '@/components/header';
import CarOwnerForm from '@/sections/owner';

export default function OwnerRegistration() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Driver', href: '/driver' },
  ];
  return (
    <>
      <Header paths={paths} />
      <div className="w-full">
        <CarOwnerForm />
      </div>
    </>
  );
}
