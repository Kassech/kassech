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
