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
import {
  queueManagerSchema,
  driverSchema,
  ownerSchema,
} from '@/types/schemas';
import { Card, CardHeader, CardTitle } from '@/components/ui/card';
import ImageUploader from '@/components/image-uploader';
import { useFetchUserData, useUpdateUserData } from '@/services/userService';
import { QUEUE_MANAGER_ROLE, DRIVER_ROLE, OWNER_ROLE } from '@/constants';

export default function EditUserForm({
  userId,
  userRole = QUEUE_MANAGER_ROLE,
}: {
  userId: string;
  userRole?: number;
}) {
  const formSchema =
    userRole === DRIVER_ROLE
      ? driverSchema
      : userRole === OWNER_ROLE
      ? ownerSchema
      : queueManagerSchema;

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    mode: 'onBlur',
  });

  const { data: userDetails, isLoading, error } = useFetchUserData(userId); // Fetch user data for editing
  const { mutateAsync } = useUpdateUserData(); // Function to update user data

  if (isLoading) return <p>Loading user details...</p>;
  if (error) return <p>Error fetching user details.</p>;

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const formData = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      if (value instanceof File || typeof value === 'string') {
        formData.append(key, value);
      }
    });

    toast.promise(
      (async () => {
        const data = await mutateAsync({ userId, formData });
        return data;
      })(),
      {
        loading: 'Updating user details...',
        success: 'User details successfully updated!',
        error: (error) => error?.response?.data?.message || 'Update failed.',
      }
    );
  };

  const { errors } = form.formState;
  const errorMessage = Object.values(errors)
    .map((error) => error?.message)
    .join(', ');

  return (
    <Card className="py-8 px-4 w-full mx-2 flex flex-col items-center justify-center">
      <CardHeader className="w-full flex items-start justify-start">
        <CardTitle>Edit {userRole} Details</CardTitle>
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
              defaultValue={userDetails?.Profile || null}
            />
          </div>

          {/* Common fields for all user types */}
          <FormField
            control={form.control}
            name="FirstName"
            defaultValue={userDetails?.FirstName || ''}
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
            defaultValue={userDetails?.LastName || ''}
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
            defaultValue={userDetails?.Email || ''}
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
            defaultValue={userDetails?.PhoneNumber || ''}
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

          {/* Dynamic Fields based on Role */}
          {userRole === DRIVER_ROLE && (
            <FormField
              control={form.control}
              name="DrivingLicense"
              defaultValue={userDetails?.DrivingLicense || ''}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Driving License</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="Enter driving license number"
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          )}

          {userRole === OWNER_ROLE && (
            <FormField
              control={form.control}
              name="CarDetails"
              defaultValue={userDetails?.CarDetails || ''}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Car Details</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder="Enter car details" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          )}

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
