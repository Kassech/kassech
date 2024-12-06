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
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

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

const vehicleSchema = z.object({
  carType: z.string().min(1, { message: "Car type is required" }),
  licenseNumber: z.string().min(1, { message: "License number is required" }),
  vin: z.string().min(1, { message: "VIN is required" }),
  make: z.string().min(1, { message: "Make is required" }),
  year: z
    .string()
    .regex(/^\d{4}$/, { message: "Year must be a valid 4-digit year" }),
  color: z.string().min(1, { message: "Car color is required" }),
  carPicture: z.any().optional(),
});

export default function VehicleRegistration() {
  const paths = [
    { name: "Home", href: "/" },
    { name: "Dashboard", href: "/b" },
  ];

  const [activeTab, setActiveTab] = useState("owner");

  // React Hook Form setup for Owner Registration
  const {
    register: registerOwner,
    handleSubmit: handleSubmitOwner,
    formState: { errors: ownerErrors },
  } = useForm({
    resolver: zodResolver(ownerSchema),
  });

  // React Hook Form setup for Vehicle Registration
  const {
    register: registerVehicle,
    handleSubmit: handleSubmitVehicle,
    formState: { errors: vehicleErrors },
  } = useForm({
    resolver: zodResolver(vehicleSchema),
  });

  // Handlers for form submissions
  const handleOwnerSubmit = (data: any) => {
    console.log("Owner Form Data:", data);
    setActiveTab("vehicle"); // Switch to the vehicle tab on successful validation
  };

  const handleVehicleSubmit = (data: any) => {
    console.log("Vehicle Form Data:", data);
  };

  return (
    <>
      <Header paths={paths} />
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0 w-full">
        <div className="grid auto-rows-min gap-4 lg:px-20 w-full">
          <Tabs
            value={activeTab}
            onValueChange={setActiveTab}
            className="w-full"
          >
            <TabsList className="grid w-full grid-cols-2">
              <TabsTrigger value="owner">Owner</TabsTrigger>
              <TabsTrigger value="vehicle">Vehicle</TabsTrigger>
            </TabsList>

            {/* Owner Registration */}
            <TabsContent value="owner">
              <Card>
                <CardHeader>
                  <CardTitle>Owner Registration</CardTitle>
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
                          {
                            (ownerErrors.firstName as { message: string })
                              .message
                          }
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
                          {
                            (ownerErrors.lastName as { message: string })
                              .message
                          }
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
                          {
                            (ownerErrors.phoneNumber as { message: string })
                              .message
                          }
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
            </TabsContent>

            {/* Vehicle Registration */}
            <TabsContent value="vehicle">
              <Card>
                <CardHeader>
                  <CardTitle>Vehicle Registration</CardTitle>
                </CardHeader>
                <CardContent className="space-y-2 lg:w-8/12 w-full">
                  <form onSubmit={handleSubmitVehicle(handleVehicleSubmit)}>
                    <div className="space-y-1">
                      <Label htmlFor="carType">Car Type</Label>
                      <Input
                        id="carType"
                        {...registerVehicle("carType")}
                        placeholder="Enter car type"
                      />
                      {vehicleErrors.carType && (
                        <p className="text-red-500 text-sm">
                          {
                            (vehicleErrors.carType as { message: string })
                              .message
                          }
                        </p>
                      )}
                    </div>
                    <div className="space-y-1">
                      <Label htmlFor="licenseNumber">License Number</Label>
                      <Input
                        id="licenseNumber"
                        {...registerVehicle("licenseNumber")}
                        placeholder="Enter license number"
                      />
                      {vehicleErrors.licenseNumber && (
                        <p className="text-red-500 text-sm">
                          {
                            (vehicleErrors.licenseNumber as { message: string })
                              .message
                          }
                        </p>
                      )}
                    </div>
                    <div className="space-y-1">
                      <Label htmlFor="vin">VIN</Label>
                      <Input
                        id="vin"
                        {...registerVehicle("vin")}
                        placeholder="Enter VIN"
                      />
                      {vehicleErrors.vin && (
                        <p className="text-red-500 text-sm">
                          {(vehicleErrors.vin as { message: string }).message}
                        </p>
                      )}
                    </div>
                    <div className="space-y-1">
                      <Label htmlFor="make">Make</Label>
                      <Input
                        id="make"
                        {...registerVehicle("make")}
                        placeholder="Enter make"
                      />
                      {vehicleErrors.make && (
                        <p className="text-red-500 text-sm">
                          {(vehicleErrors.make as { message: string }).message}
                        </p>
                      )}
                    </div>
                    <div className="space-y-1">
                      <Label htmlFor="year">Year</Label>
                      <Input
                        id="year"
                        {...registerVehicle("year")}
                        placeholder="Enter year"
                      />
                      {vehicleErrors.year && (
                        <p className="text-red-500 text-sm">
                          {(vehicleErrors.year as { message: string }).message}
                        </p>
                      )}
                    </div>
                    <div className="space-y-1">
                      <Label htmlFor="color">Car Color</Label>
                      <Input
                        id="color"
                        {...registerVehicle("color")}
                        placeholder="Enter car color"
                      />
                      {vehicleErrors.color && (
                        <p className="text-red-500 text-sm">
                          {(vehicleErrors.color as { message: string }).message}
                        </p>
                      )}
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="carPicture">Car Picture</Label>
                      <Input
                        type="file"
                        id="carPicture"
                        {...registerVehicle("carPicture")}
                      />
                    </div>
                    <CardFooter className="pt-4 w-48">
                      <Button type="submit" className="w-full">
                        Submit
                      </Button>
                    </CardFooter>
                  </form>
                </CardContent>
              </Card>
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </>
  );
}
