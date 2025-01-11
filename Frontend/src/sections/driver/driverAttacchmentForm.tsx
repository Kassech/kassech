// pages/driver/DriverAttachmentForm.tsx
'use client';
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
import { useCreateDriver } from '@/services/driverService';

export default function DriverAttachmentForm({
  switchTab,
}: {
  switchTab: (tab: string) => void;
}) {
  const { formData, setField } = useDriverStore();
  const { mutateAsync } = useCreateDriver(); // Use the custom mutation hook to create a driver

  const form = useForm<
    z.infer<typeof driverSchema & typeof driverAttachmentSchema>
  >({
    resolver: zodResolver(driverSchema.merge(driverAttachmentSchema)), // Merge driver schema with attachment schema
    mode: 'onBlur',
    defaultValues: {
      ...formData,
      drivingLicense: formData.drivingLicense || null,
      nationalId: formData.nationalId || null,
      insuranceDocument: formData.insuranceDocument || null,
      others: formData.others || null,
    },
  });

  const onSubmit = (
    values: z.infer<typeof driverSchema & typeof driverAttachmentSchema>
  ) => {
    Object.entries(values).forEach(([key, value]) => {
      console.log(`Setting field ${key} to`, value);
      setField(key, value); // Save each field to the store
    });
    const promise = mutateAsync(values); // The promise returned by mutateAsync

    toast.promise(promise, {
      loading: 'Creating driver...',
      success: () => {
        return `Driver successfully created!`; // Success message
      },
      error: (error) => {
        return `Error: ${error}`; // Error message
      },
    });
  };

  const handleBackClick = () => {
    switchTab('person'); // Switch back to the personal tab
  };

  return (
    <Card className="py-10 px-4 w-full mx-2 flex items-center justify-center">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-2 gap-4 md:grid-cols-2"
        >
          {[
            { label: 'Driving License', name: 'drivingLicense' },
            { label: 'National ID', name: 'nationalId' },
            { label: 'Insurance Document', name: 'insuranceDocument' },
            { label: 'Others', name: 'others' },
          ].map((field) => (
            <div key={field.name}>
              <FormField
                control={form.control}
                name={field.name}
                render={() => (
                  <FormItem>
                    <FormLabel>{field.label}</FormLabel>
                    <ImageUploader
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
