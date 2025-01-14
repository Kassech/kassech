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

export default function QueueManagerForm() {
  const form = useForm<z.infer<typeof queueManagerSchema>>({
    resolver: zodResolver(queueManagerSchema),
    mode: 'onBlur',
    defaultValues: {
      firstname: '',
      lastname: '',
      email: '',
      phonenumber: '',
      KebeleId: null,
      profile: null,
    },
  });

  const onSubmit = (values: z.infer<typeof queueManagerSchema>) => {
    console.log('Form values:', values); 

    const queueManagerData = {
      ...values,
      profile: values.profile as File, 
      kebeleId: values.KebeleId as File, 
    };

    console.log('Prepared data:', queueManagerData); 
    toast.success('Form data logged successfully!');
  };


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
                form.setValue('profile', file)
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
            name="firstname"
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
            name="lastname"
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
            name="phonenumber"
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
  );
}
