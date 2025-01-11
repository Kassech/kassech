import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@/components/ui/button";
import { useTranslation } from "react-i18next";
import CardWrapper from "./card-wrapper";
import { useLogin } from "@/hooks/useAuth";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import LoadingSpinner from "@/components/loading-spinner";

// Zod validation schema for login
const loginSchema = z.object({
  phoneNumber: z
    .string()
    .min(10, { message: "Phone number must be at least 10 characters." })
    .regex(/^\+251[0-9]{9}$/, { message: "Invalid phone number format." }),
  password: z
    .string()
    .min(6, { message: "Password must be at least 6 characters." }),
});

type LoginFormData = z.infer<typeof loginSchema>;

const LoginForm: React.FC = () => {
  const { t } = useTranslation();
  const { mutateAsync, isError, isLoading, error } = useLogin();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginFormData) => {
    await mutateAsync({
      email_or_phone: data.phoneNumber,
      password: data.password,
    });
  };


  if (isLoading) {
    return <LoadingSpinner />
  }
  return (
    <CardWrapper
      label={t("Welcome Back")}
      title={t("Login")}
      backButtonHref="/register"
      backButtonLabel="Don't have an account? Sign Up"
    >
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
        {/* Phone Number */}
        <div className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="phoneNumber" className="block text-sm font-medium">
              {t("Phone Number")}
            </Label>
            <Input
              type="text"
              id="phoneNumber"
              defaultValue="+251"
              placeholder="+2519... / +2517.."
              className="w-full px-4 py-2 border rounded-md"
              {...register("phoneNumber")}
            />
            {errors.phoneNumber && (
              <p className="text-red-500 font-extraBold text-xs">
                {errors.phoneNumber.message}
              </p>
            )}
          </div>

          {/* Password */}
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
        </div>

        {isError && (
          <div className="text-red-500 text-xs font-extraBold text-center">
            {error?.response?.data?.error || 'An unknown error occurred'}
          </div>
        )}

        {/* Submit Button */}
        <Button type="submit" className="w-full" disabled={isLoading}>
          {t("Login")}
        </Button>
      </form>
    </CardWrapper>
  );
};

export default LoginForm;
