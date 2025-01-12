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
import { useCreateOwner } from '@/services/OwnerRegistration';

export default function CarOwnerForm() {
  const { mutate } = useCreateOwner();

  const form = useForm<z.infer<typeof ownerSchema>>({
    resolver: zodResolver(ownerSchema),
    mode: 'onBlur',
    defaultValues: {
      firstName: '',
      lastName: '',
      email: '',
      phoneNumber: '',
      profilePicture: null,
      KebeleId: null,
      insurance: null,
    },
  });


  const onSubmit = (values: z.infer<typeof ownerSchema>) => {
    console.log('Form values:', values); // Debug log

    const ownerData = {
      ...values,
      profilePicture: values.profilePicture as File,
      kebeleId: values.KebeleId as File, // Ensure camelCase
      insurance: values.insurance as File,
    };

    console.log('Prepared data for mutation:', ownerData); // Debug log

    mutate(ownerData, {
      onSuccess: () => {
        console.log('Mutation successful'); // Debug log
        toast.success('Form submitted successfully!');
        form.reset();
      },
      onError: (error: any) => {
        console.error('Mutation error:', error); // Debug log
        toast.error(
          error?.response?.data?.message || 'Form submission failed.'
        );
      },
    });
  };

  return (
    <>
      <Card className="py-8 px-4 w-full mx-2 flex flex-col items-center justify-center">
        <CardHeader  className="w-full flex items-start justify-start">
          <CardTitle>Car Owner Registration</CardTitle>
        </CardHeader>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="grid grid-cols-1 gap-4 md:grid-cols-2"
          >
            <div className="col-span-full">
              <ImageUploader
                initialPreview={form.getValues('profilePicture')}
                onImageUpload={(file) => form.setValue('profilePicture', file)}
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
              name="firstName"
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
              name="lastName"
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
              name="email"
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
              name="phoneNumber"
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
              name="insurance"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Insurance</FormLabel>
                  <FormControl>
                    <Input
                      type="file"
                      onChange={(e) =>
                        form.setValue('insurance', e.target.files?.[0] || null)
                      }
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="KebeleId"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Kebele Id</FormLabel>
                  <FormControl>
                    <Input
                      type="file"
                      onChange={(e) =>
                        form.setValue('KebeleId', e.target.files?.[0] || null)
                      }
                    />
                  </FormControl>
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
