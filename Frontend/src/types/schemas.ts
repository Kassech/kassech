import { z } from "zod";
export const driverSchema = z.object({
    FirstName: z.string().min(3, { message: "First name is required and cannot be empty" }),
    LastName: z.string().min(3, { message: "Last name is required and cannot be empty" }),
    Email: z.string().email({ message: "Invalid email address format. Please enter a valid email" }),
    PhoneNumber: z
        .string()
        .regex(/^\+251\d{9}$/, { message: "Invalid phone number format. It should start with +251 followed by 9 digits" }),
    Password: z.string().min(6, { message: "Password must be at least 6 characters long" }),
    Role: z.number().min(1, { message: "Role is required and must be a positive number" }),
    Profile: z
        .instanceof(File)
        .refine((file) => file.size !== 0, { message: "Please upload an image file. The file cannot be empty" }),
});


export const driverAttachmentSchema = z.object({
    drivingLicense: z.instanceof(File).optional(),
    nationalId: z.instanceof(File).optional(),
    insuranceDocument: z.instanceof(File).optional(),
    others: z.instanceof(File).optional(),
  });
export const vehicleSchema = z.object({
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
  profile: z.instanceof(File).refine((file) => file.size !== 0, {
    message: 'Please upload an image file. The file cannot be empty',
  }),
  libre: z.any().refine((file) => file instanceof File && file.size > 0, {
    message: 'Libre document is required',
  }),
});

export const ownerSchema = z.object({
  firstName: z.string().min(1, { message: 'First name is required' }),
  lastName: z.string().min(1, { message: 'Last name is required' }),
  email: z.string().email({ message: 'Invalid email address' }),
  phoneNumber: z
    .string()
    .regex(/^\+251\d{9}$/, { message: 'Invalid phone number format' }),
  profilePicture: z
    .instanceof(File)
    .nullable()
    .refine((file) => file && file.size > 0, {
      message: 'Profile picture is required',
    }),
  KebeleId: z
    .instanceof(File)
    .nullable()
    .refine((file) => file && file.size > 0, {
      message: 'Kebele id document is required',
    })
    .refine((file) => file && file.size <= 5 * 1024 * 1024, {
      message: 'Kebele id document must be less than 5MB',
    }),
  insurance: z
    .instanceof(File)
    .nullable()
    .refine((file) => file && file.size > 0, {
      message: 'Insurance document is required',
    })
    .refine((file) => file && file.size <= 5 * 1024 * 1024, {
      message: 'Insurance document must be less than 5MB',
    }),
});

