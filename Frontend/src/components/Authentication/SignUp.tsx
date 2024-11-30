import React from "react";
import CardWrapper from "./card-wrapper";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
import { Input } from "../ui/input";

const LoginPage: React.FC = () => {
  return (
    <CardWrapper
      label="Welcome Back"
      title="Create Account"
      backButtonHref="/login"
      backButtonLabel="Already have an account? Login."
    >
      <div className="space-y-6">
        <div className="space-y-4">
          <div className="space-y-2">
            <label htmlFor="name" className="block text-sm font-medium">
              Full Name
            </label>
            <input
              type="text"
              id="name"
              placeholder="full name"
              className="w-full px-4 py-2 border rounded-md"
            />
          </div>
          <div className="space-y-2">
            <label htmlFor="phoneNumber" className="block text-sm font-medium">
              Phone Number
            </label>
            <input
              type="text"
              id="phoneNumber"
              placeholder="+251.."
              className="w-full px-4 py-2 border rounded-md"
            />
          </div>
          <div className="space-y-2">
            <label htmlFor="password" className="block text-sm font-medium">
              Password
            </label>
            <input
              type="password"
              id="password"
              placeholder="******"
              className="w-full px-4 py-2 border rounded-md"
            />
          </div>
          <div className="space-y-2 ">
            <Label htmlFor="picture">Picture</Label>
            <Input id="picture" type="file" placeholder="Choose from file" />
          </div>
        </div>
        <Button type="button" className="w-full">
          SignUp
        </Button>
      </div>
    </CardWrapper>
  );
};

export default LoginPage;
