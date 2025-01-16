import Header from '@/components/header';
import QueueManager from '@/sections/QueueManager';

export default function QueueManagerRegistration() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'QueueManager', href: '/queueManager' },
  ];
  return (
    <>
      <Header paths={paths} />
      <div className="w-full">
        <QueueManager />
      </div>
    </>
  );
}
