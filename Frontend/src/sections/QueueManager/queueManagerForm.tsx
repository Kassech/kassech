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
import { useCreateUser, useUpdateUserData } from '@/services/userService';
import { QUEUE_MANAGER_ROLE } from '@/constants';

export default function QueueManagerForm({
  defaultValues = null,
}: {
  defaultValues?: Partial<z.infer<typeof queueManagerSchema>> | null;
}) {
  const form = useForm<z.infer<typeof queueManagerSchema>>({
    resolver: zodResolver(queueManagerSchema),
    mode: 'onBlur',
    defaultValues: {
      first_name: defaultValues?.first_name || '',
      last_name: defaultValues?.last_name || '',
      email: defaultValues?.email || '',
      phone_number: defaultValues?.phone_number || '',
      national_id: defaultValues?.national_id || null,
      profile_picture: defaultValues?.profile_picture || null,
      roles: defaultValues?.roles ?? [QUEUE_MANAGER_ROLE.toString()],
      id: defaultValues?.id || '',
    },
  });
  console.log('profile type', typeof defaultValues?.profile_picture);

  const createUser = useCreateUser();
  const updateUser = useUpdateUserData();

  const onSubmit = async (values: z.infer<typeof queueManagerSchema>) => {
    console.log('Form values:', values);

    const formData = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      if (value instanceof File || typeof value === 'string') {
        let keyToUse = key;

        if (key === 'first_name') keyToUse = 'FirstName';
        if (key === 'last_name') keyToUse = 'LastName';
        if (key === 'email') keyToUse = 'Email';
        if (key === 'phone_number') keyToUse = 'PhoneNumber';
        if (key === 'roles') keyToUse = 'Role'; // Assuming roles should be passed as a single value.

        // Append the key-value pair to the FormData
        formData.append(keyToUse, value);
      }
    });

    const isEdit = !!defaultValues?.id;
    console.log('Prepared form data for mutation:', formData);

    toast.promise(
      (async () => {
        if (isEdit) {
          return await updateUser.mutateAsync({
            id: defaultValues.id!.toString(),
            userData: formData,
          });
        } else {
          return await createUser.mutateAsync(formData);
        }
      })(),
      {
        loading: isEdit
          ? 'Updating Queue Manager...'
          : 'Creating Queue Manager...',
        success: isEdit
          ? 'Queue Manager successfully updated!'
          : 'Queue Manager successfully created!',
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
                form.setValue('profile_picture', file)
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
            name="first_name"
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
            name="last_name"
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
            name="email"
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
            name="phone_number"
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
                <FormLabel>National id</FormLabel>
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
            <FormMessage>{errorMessage}</FormMessage>
          </FormItem>
        </form>
      </Form>
    </Card>
  );
}
