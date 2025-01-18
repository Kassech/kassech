import { z } from "zod";
export const userSchema = z.object({
    FirstName: z.string().min(3, { message: "First name is required and cannot be empty" }),
    LastName: z.string().min(3, { message: "Last name is required and cannot be empty" }),
    Email: z.string().email({ message: "Invalid email address format. Please enter a valid email" }),
    PhoneNumber: z
        .string()
        .regex(/^\+251\d{9}$/, { message: "Invalid phone number format. It should start with +251 followed by 9 digits" }),
    Password: z.string().min(6, { message: "Password must be at least 6 characters long" }),
    Role: z.number().min(1, { message: "Role is required and must be a positive number" }),
    Profile: z
        .instanceof(File).nullable()
        .refine((file) => file && file.size !== 0, { message: "Please upload an image file. The file cannot be empty" }),
});

export const driverSchema = userSchema.omit({ Password: true });

export const queueManagerSchema = userSchema.omit({ Password: true }).extend({
  national_id: z
    .instanceof(File).nullable()
    .refine((file) => file && file.size > 0, {
      message: 'Kebele id document is required',
    })
    .refine((file) => file && file.size <= 5 * 1024 * 1024, {
      message: 'Kebele id document must be less than 5MB',
    }),
});

export const driverAttachmentSchema = z.object({
    driving_license: z.instanceof(File).optional(),
    national_id: z.instanceof(File).optional(),
    insurance_document: z.instanceof(File).optional(),
    other_file: z.instanceof(File).optional(),
  });


export const vehicleSchema = z.object({
  carType: z.string().min(1, { message: 'Car type is required' }),
  // licenseNumber: z.string().min(1, { message: 'License number is required' }),
  vin: z.string().min(1, { message: 'VIN is required' }),
  make: z.string().min(1, { message: 'Make is required' }),
  year: z
    .string()
    .regex(/^\d{4}$/, { message: 'Year must be a valid 4-digit year' }),
  color: z.string().min(1, { message: 'Car color is required' }),

  // For carPicture: nullable and optional
  carPicture: z
    .instanceof(File)
    .nullable() // Allows the value to be null
    .refine((file) => file === null || (file && file.size > 0), {
      message: 'Car picture is required',
    })
    .optional(), // Makes the field optional

  // Bollo: nullable and optional
  bollo: z
    .instanceof(File)
    .nullable() // Allows null
    .refine((file) => file === null || file.size > 0, {
      message: 'Bollo document is required',
    })
    .refine((file) => file === null || file.size <= 5 * 1024 * 1024, {
      message: 'Bollo document must be less than 5MB',
    })
    .optional(),

  // Insurance: nullable and optional
  insurance: z
    .instanceof(File)
    .nullable() // Allows null
    .refine((file) => file === null || file.size > 0, {
      message: 'Insurance document is required',
    })
    .refine((file) => file === null || file.size <= 5 * 1024 * 1024, {
      message: 'Insurance document must be less than 5MB',
    })
    .optional(),

  // Libre: nullable and optional
  libre: z
    .instanceof(File)
    .nullable() // Allows null
    .refine((file) => file === null || file.size > 0, {
      message: 'Libre document is required',
    })
    .refine((file) => file === null || file.size <= 5 * 1024 * 1024, {
      message: 'Libre document must be less than 5MB',
    })
    .optional(),
  ownerID: z.object({
    id: z.string().min(1, { message: 'Owner ID is required' }),
  }),
});

export const ownerSchema = userSchema.omit({ Password: true }).extend({
  national_id: z
    .instanceof(File)
    .nullable()
    .refine((file) => file && file.size > 0, {
      message: 'Kebele id document is required',
    })
    .refine((file) => file && file.size <= 5 * 1024 * 1024, {
      message: 'Kebele id document must be less than 5MB',
    }),
  insurance_document: z
    .instanceof(File)
    .nullable()
    .refine((file) => file && file.size > 0, {
      message: 'Insurance document is required',
    })
    .refine((file) => file && file.size <= 5 * 1024 * 1024, {
      message: 'Insurance document must be less than 5MB',
    }),
});
