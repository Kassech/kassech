// pages/driver/Driver.jsx

import Header from '@/components/header';
import DriverForm from '@/sections/driver/driverForm';
import DriverAttachmentForm from '@/sections/driver/driverAttacchmentForm';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';

export default function DriverPage() {
  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Driver', href: '/driver' },
  ];
  return (
    <>
      <Header paths={paths} />
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        <Tabs defaultValue="person" className="w-full">
          <TabsList className="w-full border-b border-gray-200">
            <TabsTrigger
              value="person"
              className="py-2 px-4 text-sm font-semibold text-gray-700 hover:text-blue-600"
            >
              Person
            </TabsTrigger>
            <TabsTrigger
              value="attachments"
              className="py-2 px-4 text-sm font-semibold text-gray-700 hover:text-blue-600"
            >
              Attachments
            </TabsTrigger>
          </TabsList>
          <TabsContent value="person">
            <DriverForm />
          </TabsContent>
          <TabsContent value="attachments">
            <DriverAttachmentForm />
          </TabsContent>
        </Tabs>
      </div>
    </>
  );
}
