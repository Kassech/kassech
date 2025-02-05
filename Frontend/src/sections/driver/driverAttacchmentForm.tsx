import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Button } from '@/components/ui/button';
import { toast } from 'sonner';
import { driverSchema, driverAttachmentSchema } from '@/types/schemas'; // Import attachment schema
import ImageUploader from '@/components/image-uploader';
import { Card } from '@/components/ui/card';
import { useDriverStore } from '@/store/driverStore';
import { useCreateUser, useUpdateUserData } from '@/services/userService';

export default function DriverAttachmentForm({
  switchTab,
  defaultValues=null,
}: {
  switchTab: (tab: string) => void;
  defaultValues?: Partial<z.infer<typeof driverAttachmentSchema>> | null;
}) {
  const { formData, setField } = useDriverStore();
const createUser = useCreateUser();
const updateUser = useUpdateUserData();
  const form = useForm<
    z.infer<typeof driverSchema & typeof driverAttachmentSchema>
  >({
    resolver: zodResolver(driverSchema.merge(driverAttachmentSchema)), // Merge driver schema with attachment schema
    mode: 'onBlur',
    defaultValues: {
      ...formData,
      driving_license:
        defaultValues?.driving_license || formData?.driving_license || null,
      national_id: defaultValues?.national_id || formData?.national_id || null,
      insurance_document:
        defaultValues?.insurance_document ||
        formData?.insurance_document ||
        null,
      other_file: defaultValues?.other_file || formData?.other_file || null,
    },
  });

  const onSubmit = async (
    values: z.infer<typeof driverSchema & typeof driverAttachmentSchema>
  ) => {
    const formData = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      if (value instanceof File || typeof value === 'string') {
        formData.append(key, value);
      }
    });
    const isEdit = !!defaultValues?.ID;
    console.log(formData);

    toast.promise(
      (async () => {
        if (isEdit) {
          return await updateUser.mutateAsync({
            id: defaultValues.ID!.toString(),
            userData: formData,
          });
        } else {
          return await createUser.mutateAsync(formData);
        }
      })(),
      {
        loading: isEdit ? 'Updating driver...' : 'Creating driver...',
        success: isEdit
          ? 'Driver successfully updated!'
          : 'Driver successfully created!',
        error: (error) =>
          error?.response?.data?.message || 'Submission failed.',
      }
    );
  };


  const handleBackClick = () => {
    switchTab('person'); // Switch back to the personal tab
  };

  return (
    <Card className="py-10 px-4 w-full mx-2 flex items-center justify-center">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-2 gap-4 md:grid-cols-2 space-x-5"
        >
          {[
            { label: 'Driving License', name: 'driving_license' as const },
            { label: 'National ID', name: 'national_id' as const },
            { label: 'Insurance Document', name: 'insurance_document' as const },
            { label: 'Others', name: 'other_file' as const},
          ].map((field) => (
            <div key={field.name}>
              <FormField
                control={form.control}
                name={field.name}
                render={() => (
                  <FormItem>
                    <FormLabel>{field.label}</FormLabel>
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
            </div>
          ))}

          <Button
            type="button" // Make this a button with type "button" to not trigger form submission
            onClick={handleBackClick} // Call the back handler
            className="w-full rounded-lg mt-auto"
            variant="outline"
          >
            Back
          </Button>
          <Button
            type="submit"
            disabled={form.formState.isSubmitting}
            className="w-full rounded-lg mt-auto"
          >
            Submit
          </Button>
        </form>
      </Form>
    </Card>
  );
}
