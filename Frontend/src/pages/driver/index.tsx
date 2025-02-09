import { useState } from 'react';
import Header from '@/components/header';
// import DriverForm from '@/sections/Driver/driverForm';
// import DriverAttachmentForm from '@/sections/Driver/driverAttacchmentForm';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { driverAttachmentSchema, driverSchema } from '@/types/schemas';
import { z } from 'zod';
import { DRIVER_ROLE } from '@/constants';
import DriverForm from '@/sections/driver/driverForm';
import DriverAttachmentForm from '@/sections/driver/driverAttacchmentForm';

export default function DriverPage({
  defaultValues = {},
}: {
  defaultValues?: Partial<
    z.infer<typeof driverAttachmentSchema> & z.infer<typeof driverSchema>
  > | null;
}) {
  const [activeTab, setActiveTab] = useState('person'); // State to control active tab
  const [isTabDisabled, setIsTabDisabled] = useState(true); // Disable tabs initially

  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Driver', href: '/driver' },
  ];

  const switchTab = (tab: any) => {
    setActiveTab(tab); // Programmatically switch tabs
    setIsTabDisabled(false); // Enable tabs after switching via code
  };

  const isHeaderVisible =
    defaultValues === null || Object.keys(defaultValues).length === 0;

  return (
    <>
      {isHeaderVisible ? <Header paths={paths} /> : null}
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
            <DriverForm
              defaultValues={{
                ID: defaultValues?.ID || '',
                FirstName: defaultValues?.FirstName || '',
                LastName: defaultValues?.LastName || '',
                Email: defaultValues?.Email || '',
                PhoneNumber: defaultValues?.PhoneNumber || '',
                Profile: defaultValues?.Profile || null,
                Role: defaultValues?.Role || DRIVER_ROLE.toString(),
              }}
              switchTab={switchTab}
            />
          </TabsContent>
          <TabsContent value="attachments">
            <DriverAttachmentForm
              defaultValues={{
                ID: defaultValues?.ID || '',
                driving_license: defaultValues?.driving_license || null,
                national_id: defaultValues?.national_id || null,
                insurance_document: defaultValues?.insurance_document || null,
                other_file: defaultValues?.other_file || null,
              }}
              switchTab={switchTab}
            />
          </TabsContent>
        </Tabs>
      </div>
    </>
  );
}
