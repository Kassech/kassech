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
import { toast } from 'sonner';
import { driverSchema } from '@/types/schemas';
import ImageUploader from '@/components/image-uploader';
import { Card } from '@/components/ui/card';
import { DRIVER_ROLE } from '@/constants';
import { useDriverStore } from '@/store/driverStore';

export default function DriverAttachmentForm() {
  const { formData, setField, resetForm } = useDriverStore();
  console.log('🚀 ~ DriverAttachmentForm ~ formData:', formData);

  const form = useForm<z.infer<typeof driverSchema>>({
    resolver: zodResolver(driverSchema),
    mode: 'onBlur',
    defaultValues: {
      ...formData,
      role: DRIVER_ROLE,
    },
  });

  const onSubmit = (values: z.infer<typeof driverSchema>) => {
    console.groupCollapsed('DriverAttachmentForm.onSubmit');
    console.log('values:', values);
    Object.entries(values).forEach(([key, value]) => {
      console.log(`Setting field ${key} to`, value);
      setField(key, value); // Save each field to the store
    });
    console.groupEnd();
    toast.success('Form submitted successfully 🎉');
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
          ].map((field) => (
            <div key={field.name}>
              <FormField
                control={form.control}
                name={field.name}
                render={() => (
                  <FormItem>
                    <FormLabel>{field.label}</FormLabel>
                    <ImageUploader
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
