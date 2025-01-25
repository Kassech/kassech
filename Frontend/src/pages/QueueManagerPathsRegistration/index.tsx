import Header from '@/components/header';
import { AssignmentForm } from './components/AssignmentForm';

export default function QueueManagerPaths() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Dashboard', href: '/b' },
  ];

  return (
    <>
      <Header paths={paths} />
      <div className="h-screen flex flex-col">
        <AssignmentForm />
      </div>
    </>
  );
}
