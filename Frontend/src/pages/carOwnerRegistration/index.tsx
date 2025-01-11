// components/OwnerRegistration.tsx
import React from 'react';
import { useState } from 'react';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { create } from 'zustand';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import Header from '@/components/header';

// Define the types for Zustand store state
interface OwnerStore {
  firstName: string;
  lastName: string;
  email: string;
  phoneNumber: string;
  profilePicture: File | null;
  kebeleId: File | null;
  insurance: File | null;
  libre: File | null;
  setOwnerData: (data: OwnerData) => void;
}

interface OwnerData {
  firstName: string;
  lastName: string;
  email: string;
  phoneNumber: string;
  profilePicture: File | null;
  kebeleId: File | null;
  insurance: File | null;
  libre: File | null;
}

// Zustand store for form data
const useOwnerStore = create<OwnerStore>((set) => ({
  firstName: '',
  lastName: '',
  email: '',
  phoneNumber: '',
  profilePicture: null,
  kebeleId: null,
  insurance: null,
  libre: null,
  setOwnerData: (data) => set(() => ({ ...data })),
}));

// Validation schema using zod
const ownerSchema = z.object({
  firstName: z.string().min(1, { message: 'First name is required' }),
  lastName: z.string().min(1, { message: 'Last name is required' }),
  email: z.string().email({ message: 'Invalid email address' }),
  phoneNumber: z
    .string()
    .regex(/^\+251\d{9}$/, { message: 'Invalid phone number format' }),
  profilePicture: z
    .instanceof(File)
    .refine((file) => file.size > 0, {
      message: 'Profile picture is required',
    }),
  kebeleId: z
    .instanceof(File)
    .refine((file) => file.size > 0, { message: 'Kebele ID is required' }),
  insurance: z
    .instanceof(File)
    .refine((file) => file.size > 0, { message: 'Insurance is required' }),
  libre: z
    .instanceof(File)
    .refine((file) => file.size > 0, { message: 'Libre is required' }),
});

export default function OwnerRegistration() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(ownerSchema),
  });

  const { setOwnerData } = useOwnerStore(); // Access Zustand's state setter

  // Form submission handler
  const handleOwnerSubmit = (data: any) => {
    console.log('Owner Form Data:', data);
    setOwnerData(data); // Save data to Zustand store
  };

  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Dashboard', href: '/b' },
  ];

  return (
    <>
      <Header paths={paths} />
      <Card>
        <CardHeader>
          <CardTitle>Owner Registration</CardTitle>
        </CardHeader>
        <CardContent className="space-y-2 lg:w-8/12 w-full">
          <form onSubmit={handleSubmit(handleOwnerSubmit)}>
            {/* First Name */}
            <div className="space-y-1">
              <Label htmlFor="firstName">First Name</Label>
              <Input
                id="firstName"
                {...register('firstName')}
                placeholder="Enter your first name"
              />
              {errors.firstName && (
                <p className="text-red-500 text-sm">
                  {(errors.firstName as any).message}
                </p>
              )}
            </div>

            {/* Last Name */}
            <div className="space-y-1">
              <Label htmlFor="lastName">Last Name</Label>
              <Input
                id="lastName"
                {...register('lastName')}
                placeholder="Enter your last name"
              />
              {errors.lastName && (
                <p className="text-red-500 text-sm">
                  {(errors.lastName as any).message}
                </p>
              )}
            </div>

            {/* Email */}
            <div className="space-y-1">
              <Label htmlFor="email">Email</Label>
              <Input
                id="email"
                {...register('email')}
                placeholder="Enter your email"
              />
              {errors.email && (
                <p className="text-red-500 text-sm">
                  {(errors.email as any).message}
                </p>
              )}
            </div>

            {/* Phone Number */}
            <div className="space-y-1">
              <Label htmlFor="phoneNumber">Phone Number</Label>
              <Input
                id="phoneNumber"
                {...register('phoneNumber')}
                placeholder="+251..."
              />
              {errors.phoneNumber && (
                <p className="text-red-500 text-sm">
                  {(errors.phoneNumber as any).message}
                </p>
              )}
            </div>

            {/* File Inputs */}
            {[
              { id: 'profilePicture', label: 'Profile Picture' },
              { id: 'kebeleId', label: 'Kebele ID' },
              { id: 'insurance', label: 'Insurance' },
              { id: 'libre', label: 'Libre' },
            ].map((fileField) => (
              <div key={fileField.id} className="space-y-2">
                <Label htmlFor={fileField.id}>{fileField.label}</Label>
                <Input
                  type="file"
                  id={fileField.id}
                  {...register(fileField.id)}
                />
                {errors[fileField.id] && (
                  <p className="text-red-500 text-sm">
                    {(errors[fileField.id] as any).message}
                  </p>
                )}
              </div>
            ))}

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
