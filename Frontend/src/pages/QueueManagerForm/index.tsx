import React, { useState } from "react";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import Header from "@/components/header";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

// Validation Schemas
const ownerSchema = z.object({
  firstName: z.string().min(1, { message: "First name is required" }),
  lastName: z.string().min(1, { message: "Last name is required" }),
  email: z.string().email({ message: "Invalid email address" }),
  phoneNumber: z
    .string()
    .regex(/^\+251\d{9}$/, { message: "Invalid phone number format" }),
  profilePicture: z.any().optional(),
});

export default function DriverForm() {
  const paths = [
    { name: "Home", href: "/" },
    { name: "Dashboard", href: "/b" },
  ];

  const {
    register: registerOwner,
    handleSubmit: handleSubmitOwner,
    formState: { errors: ownerErrors },
  } = useForm({
    resolver: zodResolver(ownerSchema),
  });

  // Handlers for form submissions
  const handleOwnerSubmit = (data: any) => {
    console.log("Owner Form Data:", data);
  };

  return (
    <>
      <Header paths={paths} />
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0 w-full">
        <div className="grid auto-rows-min gap-4 lg:px-20 w-full">
          <Card>
            <CardHeader>
              <CardTitle>Queue Manager Registration</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2 lg:w-8/12 w-full">
              <form onSubmit={handleSubmitOwner(handleOwnerSubmit)}>
                <div className="space-y-1">
                  <Label htmlFor="firstName">First Name</Label>
                  <Input
                    id="firstName"
                    {...registerOwner("firstName")}
                    placeholder="Enter your first name"
                  />
                  {ownerErrors.firstName && (
                    <p className="text-red-500 text-sm">
                      {(ownerErrors.firstName as { message: string }).message}
                    </p>
                  )}
                </div>
                <div className="space-y-1">
                  <Label htmlFor="lastName">Last Name</Label>
                  <Input
                    id="lastName"
                    {...registerOwner("lastName")}
                    placeholder="Enter your last name"
                  />
                  {ownerErrors.lastName && (
                    <p className="text-red-500 text-sm">
                      {(ownerErrors.lastName as { message: string }).message}
                    </p>
                  )}
                </div>
                <div className="space-y-1">
                  <Label htmlFor="email">Email</Label>
                  <Input
                    id="email"
                    {...registerOwner("email")}
                    placeholder="Enter your email"
                  />
                  {ownerErrors.email && (
                    <p className="text-red-500 text-sm">
                      {(ownerErrors.email as { message: string }).message}
                    </p>
                  )}
                </div>
                <div className="space-y-1">
                  <Label htmlFor="phoneNumber">Phone Number</Label>
                  <Input
                    id="phoneNumber"
                    {...registerOwner("phoneNumber")}
                    placeholder="+251..."
                  />
                  {ownerErrors.phoneNumber && (
                    <p className="text-red-500 text-sm">
                      {(ownerErrors.phoneNumber as { message: string }).message}
                    </p>
                  )}
                </div>
                <div className="space-y-2">
                  <Label htmlFor="profilePicture">Profile Picture</Label>
                  <Input
                    type="file"
                    id="profilePicture"
                    {...registerOwner("profilePicture")}
                  />
                </div>
                <CardFooter className="pt-4 w-48">
                  <Button type="submit" className="w-full">
                    Next
                  </Button>
                </CardFooter>
              </form>
            </CardContent>
          </Card>
        </div>
      </div>
    </>
  );
}
