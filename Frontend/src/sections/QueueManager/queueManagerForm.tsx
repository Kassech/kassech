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
import { queueManagerSchema } from '@/types/schemas';
import { Card, CardHeader, CardTitle } from '@/components/ui/card';
import ImageUploader from '@/components/image-uploader';
import { useCreateUser } from '@/services/userService';
import { QUEUE_MANAGER_ROLE } from '@/constants';

export default function QueueManagerForm() {
  const form = useForm<z.infer<typeof queueManagerSchema>>({
    resolver: zodResolver(queueManagerSchema),
    mode: 'onBlur',
    defaultValues: {
      FirstName: '',
      LastName: '',
      Email: '',
      PhoneNumber: '',
      national_id: null,
      Profile: null,
      Role: QUEUE_MANAGER_ROLE,
    },
  });
  const { mutateAsync } = useCreateUser(); // Use the custom mutation hook
  const onSubmit = async (values: z.infer<typeof queueManagerSchema>) => {
    console.log('Form values:', values); // Debug log

    const formData = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      if (value instanceof File || typeof value === 'string') {
        formData.append(key, value);
      }
    });

    console.log('Prepared form data for mutation:', formData); // Debug log

    toast.promise(
      (async () => {
        const data = await mutateAsync(formData);
        return data;
      })(),
      {
        loading: 'Creating queue manager...',
        success: 'Queue manager successfully created!',
        error: (error) =>
          error?.response?.data?.message || 'Submission failed.',
      }
    );
  };

  const { errors } = form.formState;
  console.log('🚀 ~ QueueManagerForm ~ errors:', errors);
  const errorMessage = Object.values(errors)
    .map((error) => error.message)
    .join(', ');

  return (
    <Card className="py-8 px-4 w-full mx-2 flex flex-col items-center justify-center">
      <CardHeader className="w-full flex items-start justify-start">
        <CardTitle>Queue Manager Registration</CardTitle>
      </CardHeader>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-1 gap-4 md:grid-cols-2"
        >
          <div className="col-span-full">
            <ImageUploader
              onImageUpload={(file: File | null) =>
                form.setValue('Profile', file)
              }
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
                  <Input {...field} placeholder="Enter first name" />
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
                  <Input {...field} placeholder="Enter last name" />
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
                  <Input {...field} type="email" placeholder="Enter email" />
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
                  <Input {...field} placeholder="Enter phone number" />
                </FormControl>
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
                <FormControl>
                  <Input
                    type="file"
                    onChange={(e) =>
                      form.setValue('national_id', e.target.files?.[0] || null)
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
            <FormMessage>{errorMessage}</FormMessage>
          </FormItem>
        </form>
      </Form>
    </Card>
  );
}
