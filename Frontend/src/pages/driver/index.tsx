import { useState } from 'react';
import Header from '@/components/header';
import DriverForm from '@/sections/Driver/driverForm';
import DriverAttachmentForm from '@/sections/Driver/driverAttacchmentForm';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';

export default function DriverPage() {
  const [activeTab, setActiveTab] = useState('person'); // State to control active tab
  const [isTabDisabled, setIsTabDisabled] = useState(true); // Disable tabs initially

  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Driver', href: '/driver' },
  ];

  const switchTab = (tab) => {
    setActiveTab(tab); // Programmatically switch tabs
    setIsTabDisabled(false); // Enable tabs after switching via code
  };

  return (
    <>
      <Header paths={paths} />
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        <Tabs value={activeTab} onValueChange={setActiveTab} className="w-full">
          <TabsList className="w-full border-b border-gray-200">
            <TabsTrigger
              value="person"
              className="py-2 px-4 text-sm font-semibold text-gray-700 hover:text-blue-600"
              disabled={isTabDisabled} // Disable trigger initially
            >
              Person
            </TabsTrigger>
            <TabsTrigger
              value="attachments"
              className="py-2 px-4 text-sm font-semibold text-gray-700 hover:text-blue-600"
              disabled={isTabDisabled} // Disable trigger initially
            >
              Attachments
            </TabsTrigger>
          </TabsList>
          <TabsContent value="person">
            <DriverForm switchTab={switchTab} />
          </TabsContent>
          <TabsContent value="attachments">
            <DriverAttachmentForm switchTab={switchTab}  />
          </TabsContent>
        </Tabs>
      </div>
    </>
  );
}
