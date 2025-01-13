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
import { vehicleSchema } from '@/types/schemas';
import ImageUploader from '@/components/image-uploader';
import { Card, CardHeader, CardTitle } from '@/components/ui/card';
import { OwnerSearch } from './autoCompleteSearch';

export default function VehicleForm() {
  const form = useForm<z.infer<typeof vehicleSchema>>({
    resolver: zodResolver(vehicleSchema),
    mode: 'onBlur',
    defaultValues: {
      carType: '', 
      vin: '',
      make: '',
      year: '',
      color: '#000000',
      bollo: null, 
      insurance: null,
      libre: null,
      carPicture: null,
    },
  });

  const onSubmit = (values: z.infer<typeof vehicleSchema>) => {
    console.log('Form values:', values); // Debug log
    toast.success(`Form submitted successfully 🎉`);
  };

  return (
    <Card className="py-8 px-4 w-full mx-2 flex flex-col items-center justify-center">
      <CardHeader className="w-full flex items-start justify-start">
        <CardTitle>Vehicle Registration</CardTitle>
      </CardHeader>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-1 gap-4 md:grid-cols-2"
        >
          <div className="col-span-full">
            <ImageUploader
              onImageUpload={(file) => form.setValue('carPicture', file)}
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
            name="carType"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Car Type</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="Enter car type" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="vin"
            render={({ field }) => (
              <FormItem>
                <FormLabel>VIN</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="Enter VIN" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="make"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Make</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="Enter the car make" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="year"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Year</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="Enter the year" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="color"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Color</FormLabel>
                <FormControl>
                  <Input
                    type="color"
                    {...field}
                    placeholder="Enter car color"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="bollo"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Bollo</FormLabel>
                <FormControl>
                  <Input
                    type="file"
                    onChange={(e) =>
                      form.setValue('bollo', e.target.files?.[0] || null)
                    }
                  />
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
            name="libre"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Libre</FormLabel>
                <FormControl>
                  <Input
                    type="file"
                    onChange={(e) =>
                      form.setValue('libre', e.target.files?.[0] || null)
                    }
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div>
            <OwnerSearch />
          </div>
          <FormItem className="col-span-full">
            <FormControl>
              <Button
                type="submit"
                disabled={
                  form.formState.isSubmitting || !form.formState.isValid
                }
                className="w-full rounded-lg"
              >
                Submit
              </Button>
            </FormControl>
          </FormItem>
        </form>
      </Form>
    </Card>
  );
}
