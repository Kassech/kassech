// components/VehicleRegistration.tsx
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import Header from '@/components/header';

// Validation Schema for Vehicle
const vehicleSchema = z.object({
  carType: z.string().min(1, { message: 'Car type is required' }),
  licenseNumber: z.string().min(1, { message: 'License number is required' }),
  vin: z.string().min(1, { message: 'VIN is required' }),
  make: z.string().min(1, { message: 'Make is required' }),
  year: z
    .string()
    .regex(/^\d{4}$/, { message: 'Year must be a valid 4-digit year' }),
  color: z.string().min(1, { message: 'Car color is required' }),
  carPicture: z
    .any()
    .refine((file) => file instanceof File && file.size > 0, {
      message: 'Car picture is required',
    })
    .refine(
      (file) => file instanceof File && file.size <= 5 * 1024 * 1024, // Max size: 5MB
      { message: 'Car picture must be less than 5MB' }
    )
    .optional(),
  bollo: z
    .any()
    .refine((file) => file instanceof File && file.size > 0, {
      message: 'Bollo document is required',
    })
    .refine((file) => file instanceof File && file.size <= 5 * 1024 * 1024, {
      message: 'Bollo document must be less than 5MB',
    }),
  insurance: z
    .any()
    .refine((file) => file instanceof File && file.size > 0, {
      message: 'Insurance document is required',
    })
    .refine((file) => file instanceof File && file.size <= 5 * 1024 * 1024, {
      message: 'Insurance document must be less than 5MB',
    }),
});

interface VehicleRegistrationProps {
  onSubmit: (data: any) => void;
}

export default function VehicleRegistration({
  onSubmit,
}: VehicleRegistrationProps) {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(vehicleSchema),
  });

  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Dashboard', href: '/b' },
  ];

  return (
    <>
      <Header paths={paths} />
      <Card>
        <CardHeader>
          <CardTitle>Vehicle Registration</CardTitle>
        </CardHeader>
        <CardContent className="space-y-2 lg:w-8/12 w-full">
          <form onSubmit={handleSubmit(onSubmit)}>
            {/* Car Type */}
            <div className="space-y-1">
              <Label htmlFor="carType">Car Type</Label>
              <Input
                id="carType"
                {...register('carType')}
                placeholder="Enter car type"
              />
              {errors.carType && typeof errors.carType.message === 'string' && (
                <p className="text-red-500 text-sm">{errors.carType.message}</p>
              )}
            </div>

            {/* License Number */}
            <div className="space-y-1">
              <Label htmlFor="licenseNumber">License Number</Label>
              <Input
                id="licenseNumber"
                {...register('licenseNumber')}
                placeholder="Enter license number"
              />
              {errors.licenseNumber &&
                typeof errors.licenseNumber.message === 'string' && (
                  <p className="text-red-500 text-sm">
                    {errors.licenseNumber.message}
                  </p>
                )}
            </div>

            {/* VIN */}
            <div className="space-y-1">
              <Label htmlFor="vin">VIN</Label>
              <Input id="vin" {...register('vin')} placeholder="Enter VIN" />
              {errors.vin && typeof errors.vin.message === 'string' && (
                <p className="text-red-500 text-sm">{errors.vin.message}</p>
              )}
            </div>

            {/* Make */}
            <div className="space-y-1">
              <Label htmlFor="make">Make</Label>
              <Input id="make" {...register('make')} placeholder="Enter make" />
              {errors.make && typeof errors.make.message === 'string' && (
                <p className="text-red-500 text-sm">{errors.make.message}</p>
              )}
            </div>

            {/* Year */}
            <div className="space-y-1">
              <Label htmlFor="year">Year</Label>
              <Input id="year" {...register('year')} placeholder="Enter year" />
              {errors.year && typeof errors.year.message === 'string' && (
                <p className="text-red-500 text-sm">{errors.year.message}</p>
              )}
            </div>

            {/* Car Color */}
            <div className="space-y-1">
              <Label htmlFor="color">Car Color</Label>
              <Input
                id="color"
                {...register('color')}
                placeholder="Enter car color"
              />
              {errors.color && typeof errors.color.message === 'string' && (
                <p className="text-red-500 text-sm">{errors.color.message}</p>
              )}
            </div>

            {/* Car Picture */}
            <div className="space-y-2">
              <Label htmlFor="carPicture">Car Picture</Label>
              <Input type="file" id="carPicture" {...register('carPicture')} />
              {errors.carPicture &&
                typeof errors.carPicture.message === 'string' && (
                  <p className="text-red-500 text-sm">
                    {errors.carPicture.message}
                  </p>
                )}
            </div>

            {/* Bollo */}
            <div className="space-y-2">
              <Label htmlFor="bollo">Bollo</Label>
              <Input type="file" id="bollo" {...register('bollo')} />
              {errors.bollo && typeof errors.bollo.message === 'string' && (
                <p className="text-red-500 text-sm">{errors.bollo.message}</p>
              )}
            </div>

            {/* Insurance */}
            <div className="space-y-2">
              <Label htmlFor="insurance">Insurance</Label>
              <Input type="file" id="insurance" {...register('insurance')} />
              {errors.insurance &&
                typeof errors.insurance.message === 'string' && (
                  <p className="text-red-500 text-sm">
                    {errors.insurance.message}
                  </p>
                )}
            </div>

            {/* Submit Button */}
            <CardFooter className="pt-4 w-48">
              <Button type="submit" className="w-full">
                Submit
              </Button>
            </CardFooter>
          </form>
        </CardContent>
      </Card>
    </>
  );
}
