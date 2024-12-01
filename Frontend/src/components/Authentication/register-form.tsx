import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "../ui/button";
import { useTranslation } from "react-i18next";
import CardWrapper from "./card-wrapper";
import { Label } from "../ui/label";
import { Input } from "../ui/input";

// Zod validation schema for registration
const registerSchema = z.object({
  name: z.string().min(3, { message: "Full name must be at least 3 characters." }),
  phoneNumber: z
    .string()
    .min(10, { message: "Phone number must be at least 10 characters." })
    .regex(/^\+251[0-9]{9}$/, { message: "Invalid phone number format." }),
  password: z.string().min(6, { message: "Password must be at least 6 characters." }),
  picture: z.any(),
});

type RegisterFormData = z.infer<typeof registerSchema>;

const RegisterForm: React.FC = () => {
  const { t } = useTranslation();
  const { register, handleSubmit, formState: { errors } } = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
  });

  const onSubmit = async (data: RegisterFormData) => {
    try {
      // const response = await axios.post('/api/register', data);
      console.log("Registration successful:", data);
    } catch (error) {
      console.error("Registration failed:", error);
    }
  };

  return (
    <CardWrapper
      label={t("Create an Account")}
      title={t("Sign Up")}
      backButtonHref="/login"
      backButtonLabel="Already have an account? Login"
    >
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
        <div className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name" >
              {t("Full Name")}
            </Label>
            <Input
              type="text"
              id="name"
              placeholder={t("Full Name")}
              className="w-full px-4 py-2 border rounded-md"
              {...register("name")}
            />
            {errors.name && (
              <p className="text-red-500 text-xs">{errors.name.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="phoneNumber" className="block text-sm font-medium">
              {t("Phone Number")}
            </Label>
            <Input
              type="text"
              id="phoneNumber"
              placeholder="+2519... / +2517.."
              className="w-full px-4 py-2 border rounded-md"
              {...register("phoneNumber")}
            />
            {errors.phoneNumber && (
              <p className="text-red-500 text-xs">{errors.phoneNumber.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="password" className="block text-sm font-medium">
              {t("Password")}
            </Label>
            <Input
              type="password"
              id="password"
              placeholder="******"
              className="w-full px-4 py-2 border rounded-md"
              {...register("password")}
            />
            {errors.password && (
              <p className="text-red-500 text-xs">{errors.password.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="picture" className="block text-sm font-medium">
              {t("Profile Picture")}
            </Label>
            <Input
              type="file"
              id="picture"
              className="w-full px-4 py-2 border rounded-md"
              {...register("picture")}
            />
          </div>
        </div>

        <Button type="submit" className="w-full">
          {t("Sign Up")}
        </Button>
      </form>
    </CardWrapper>
  );
};

export default RegisterForm;
