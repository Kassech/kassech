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
import { toast } from 'sonner';
import { ownerSchema } from '@/types/schemas';
import ImageUploader from '@/components/image-uploader';
import { Card, CardHeader, CardTitle } from '@/components/ui/card';
import { useCreateUser } from '@/services/userService';
import { OWNER_ROLE } from '@/constants';

export default function CarOwnerForm({
  defaultValues = null,
}: {
  defaultValues?: Partial<z.infer<typeof ownerSchema>> | null;
}) {
  const { mutateAsync } = useCreateUser(); // Use async mutation

  const form = useForm<z.infer<typeof ownerSchema>>({
    resolver: zodResolver(ownerSchema),
    mode: 'onBlur',
    defaultValues: {
      FirstName: defaultValues?.FirstName || '',
      LastName: defaultValues?.LastName || '',
      Email: defaultValues?.Email || '',
      PhoneNumber: defaultValues?.PhoneNumber || '',
      Profile: defaultValues?.ProfilePicture || null,
      national_id: defaultValues?.national_id || null,
      insurance_document: defaultValues?.insurance_document || null,
      Role: OWNER_ROLE.toString(),
    },
  });

  console.log('defaultValues', defaultValues);

  const onSubmit = async (values: z.infer<typeof ownerSchema>) => {
    console.log('Form values:', values); // Debug log

    const formData = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      if (value instanceof File || typeof value === 'string') {
        formData.append(key, value);
      }
    });

    for (let pair of formData.entries()) {
      console.log(`${pair[0]}:`, pair[1]);
    } // Debug log

    toast.promise(
      (async () => {
        const data = await mutateAsync(formData);
        return data;
      })(),
      {
        loading: 'Creating owner...',
        success: 'Owner successfully created!',
        error: (error) =>
          error?.response?.data?.message || 'Submission failed.',
      }
    );
  };

  return (
    <>
      <Card className="py-8 px-4 w-full mx-2 flex flex-col items-center justify-center">
        <CardHeader className="w-full flex items-start justify-start">
          <CardTitle>Car Owner Registration</CardTitle>
        </CardHeader>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="grid grid-cols-1 gap-4 md:grid-cols-2"
          >
            <div className="col-span-full">
              <ImageUploader
                initialPreview={form.getValues('Profile')}
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

            <FormField
              control={form.control}
              name="insurance_document"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Insurance</FormLabel>
                  <ImageUploader
                    className="rounded-md"
                    initialPreview={form.getValues(field.name)}
                    onImageUpload={(file) => form.setValue(field.name, file)}
                    maxFileSize={5000000}
                    acceptedFormats={{
                      'image/png': [],
                      'image/jpg': [],
                      'image/jpeg': [],
                      'application/pdf': [],
                    }}
                  />
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="national_id"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>National Id</FormLabel>
                  <ImageUploader
                    className="rounded-md"
                    initialPreview={form.getValues(field.name)}
                    onImageUpload={(file) => form.setValue(field.name, file)}
                    maxFileSize={5000000}
                    acceptedFormats={{
                      'image/png': [],
                      'image/jpg': [],
                      'image/jpeg': [],
                      'application/pdf': [],
                    }}
                  />
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormItem className="col-span-full">
              <FormControl>
                <Button
                  type="submit"
                  disabled={form.formState.isSubmitting}
                  className="w-full rounded-lg mt-7"
                >
                  Submit
                </Button>
              </FormControl>
              <FormMessage />
            </FormItem>
          </form>
        </Form>
      </Card>
    </>
  );
}
