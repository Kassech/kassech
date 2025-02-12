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
import { useCreateUser, useUpdateUserData } from '@/services/userService';
import { OWNER_ROLE } from '@/constants';

export default function CarOwnerForm({
  defaultValues = null,
}: {
  defaultValues?: Partial<z.infer<typeof ownerSchema>> | null;
}) {
  const createUser = useCreateUser();
  const updateUser = useUpdateUserData();

  const form = useForm<z.infer<typeof ownerSchema>>({
    resolver: zodResolver(ownerSchema),
    mode: 'onBlur',
    defaultValues: {
      first_name: defaultValues?.first_name || '',
      last_name: defaultValues?.last_name || '',
      email: defaultValues?.email || '',
      phone_number: defaultValues?.phone_number || '',
      profile_picture: defaultValues?.profile_picture || null,
      national_id: defaultValues?.national_id || null,
      insurance_document: defaultValues?.insurance_document || null,
      roles: [OWNER_ROLE.toString()],
      id: defaultValues?.id || '',
    },
  });

  console.log('defaultValues', defaultValues);

  const onSubmit = async (values: z.infer<typeof ownerSchema>) => {
    console.log('Form values:', values); // Debug log

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

    for (let pair of formData.entries()) {
      console.log(`${pair[0]}:`, pair[1]);
    } // Debug log
    const isEdit = !!defaultValues?.id;

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
        loading: isEdit ? 'Updating owner...' : 'Creating Owner...',
        success: isEdit
          ? 'Car Owner successfully updated!'
          : 'Car Owner successfully created!',
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
                initialPreview={form.getValues('profile_picture')}
                onImageUpload={(file) => form.setValue('profile_picture', file)}
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
                    <Input {...field} placeholder="Enter your first name" />
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
              name="phone_number"
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
