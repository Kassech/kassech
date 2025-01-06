"use client";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import {
    Form,
    FormField,
    FormItem,
    FormLabel,
    FormControl,
    FormMessage,
} from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { toast } from "sonner";
import { driverSchema } from "@/types/schemas";
import ImageUploader from "@/components/image-uploader";
import { Card } from "@/components/ui/card";
import { DRIVER_ROLE } from "@/constants";

export default function DriverForm() {
    const form = useForm<z.infer<typeof driverSchema>>({
        resolver: zodResolver(driverSchema),
        mode: "onBlur",
        defaultValues: {
            role: DRIVER_ROLE
        },
    });

    const onSubmit = (values: z.infer<typeof driverSchema>) => {
        console.log(values);
        toast.success(`Form submitted successfully 🎉`);
    };

    return (
        <Card className="py-10 px-4 w-[400px] mx-2 flex items-center justify-center">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="grid grid-cols-1 gap-4 md:grid-cols-2">
                    <div className="col-span-full">
                        <ImageUploader
                            onImageUpload={(file) => form.setValue("profile", file)}
                            maxFileSize={2000000} // Optional: Customize max size
                            acceptedFormats={{ "image/png": [], "image/jpg": [], "image/jpeg": [] }} // Optional
                        />
                    </div>

                    <FormField
                        control={form.control}
                        name="firstName"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>First Name</FormLabel>
                                <FormControl>
                                    <Input {...field} placeholder="Enter your first name" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="lastName"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Last Name</FormLabel>
                                <FormControl>
                                    <Input {...field} placeholder="Enter your last name" />
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
                                    <Input {...field} placeholder="Enter your email" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="phoneNumber"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Phone Number</FormLabel>
                                <FormControl>
                                    <Input {...field} placeholder="+251..." />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input type="password" {...field} placeholder="Enter your password" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormItem>
                                <FormLabel>Submit</FormLabel>
                                <FormControl>
                                <Button
                            type="submit"
                            disabled={form.formState.isSubmitting}
                            className="w-full rounded-lg"
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
