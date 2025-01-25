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
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { useCreateVehicle } from '@/services/vehicleService';

export default function VehicleForm() {
  const form = useForm<z.infer<typeof vehicleSchema>>({
    resolver: zodResolver(vehicleSchema),
    mode: 'onBlur',
    defaultValues: {
      carType: '',
      vin: '',
      make: '',
      year: '',
      color: '',
      bollo: null,
      insurance: null,
      libre: null,
      carPicture: null,
      ownerID: '',
    },
  });
   const { mutateAsync } = useCreateVehicle(); 
    const onSubmit = async (values: z.infer<typeof vehicleSchema>) => {
      console.log('Form values:', values); 

        const formData = new FormData();
        Object.entries(values).forEach(([key, value]) => {
          if (value instanceof File || typeof value === 'string') {
            console.log(key,value)
            formData.append(key, value);
          }
        });
    
      console.log('Prepared data:', formData);
      toast.promise(
        (async () => {
          const data = await mutateAsync(formData);
          return data;
        })(),
        {
          loading: 'Creating vehicle...',
          success: 'Vehicle successfully created!',
          error: (error) =>
            error?.response?.data?.message || 'Submission failed.',
        }
      );
    };
 
 const handleOwnerSelect = (id: string) => {
   form.setValue('ownerID', id); 
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
                    type="text"
                    {...field}
                    placeholder="Enter color name (e.g., red, blue, black)"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="carType"
            render={({ field }) => (
              <FormItem className="md:pt-8">
                <Select
                  onValueChange={(value) => form.setValue('carType', value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="Select vehicle type" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectGroup>
                      <SelectLabel>Car type</SelectLabel>
                      <SelectItem value="5">Mini Bus15</SelectItem>
                      <SelectItem value="6">Mini Bus16</SelectItem>
                      <SelectItem value="7">Mini Bus17</SelectItem>
                      <SelectItem value="8">Mini Bus18</SelectItem>
                      <SelectItem value="9">Mini Bus19</SelectItem>
                    </SelectGroup>
                  </SelectContent>
                </Select>
              </FormItem>
            )}
          />
          <div className="md:pt-8">
            <OwnerSearch onOwnerSelect={handleOwnerSelect} />
          </div>

          <FormField
            control={form.control}
            name="bollo"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Bollo</FormLabel>
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
            name="insurance"
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
            name="libre"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Libre</FormLabel>
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
  );
}
