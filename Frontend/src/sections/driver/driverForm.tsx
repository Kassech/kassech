'use client';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from '@/components/ui/form';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { driverSchema } from '@/types/schemas';
import ImageUploader from '@/components/image-uploader';
import { Card } from '@/components/ui/card';
import { DRIVER_ROLE } from '@/constants';
import { useDriverStore } from '@/store/driverStore';

export default function DriverForm({
  switchTab,
}: {
  switchTab: (tab: string) => void;
}) {
  const { formData, setField } = useDriverStore();
  console.log('🚀 ~ DriverForm ~ formData:', formData);

  const form = useForm<z.infer<typeof driverSchema>>({
    resolver: zodResolver(driverSchema),
    mode: 'onBlur',
    defaultValues: {
      ...formData,
      Role: DRIVER_ROLE,
    },
  });

  const onSubmit = (values: z.infer<typeof driverSchema>) => {
    Object.entries(values).forEach(([key, value]) => {
      setField(key, value); // Save each field to the store
    });

    switchTab('attachments');
  };

  return (
    <Card className="py-10 px-4 w-full mx-2 flex items-center justify-center">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-1 gap-4 md:grid-cols-2"
        >
          <div className="col-span-full">
            <ImageUploader
              initialPreview={formData.Profile}
              onImageUpload={(file) => form.setValue('Profile', file)}
              maxFileSize={2000000}
              acceptedFormats={{
                'image/png': [],
                'image/jpg': [],
                'image/jpeg': [],
              }}
            />
          </div>

          <FormField
            control={form.control}
            name="FirstName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>First Name</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="Enter your first name" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="LastName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Last Name</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="Enter your last name" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="Email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="Enter your email" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="PhoneNumber"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Phone Number</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="+251..." />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* <FormField
            control={form.control}
            name="Password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input
                    type="password"
                    {...field}
                    placeholder="Enter your password"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          /> */}

          <FormItem>
            <FormLabel>Go To Next Step</FormLabel>
            <FormControl>
              <Button
                type="submit"
                disabled={form.formState.isSubmitting}
                className="w-full rounded-lg mt-7"
              >
                Next
              </Button>
            </FormControl>
            <FormMessage />
          </FormItem>
        </form>
      </Form>
    </Card>
  );
}
