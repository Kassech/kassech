import React from "react";
import CardWrapper from "./card-wrapper";
import { Button } from "../ui/button";

const LoginPage: React.FC = () => {
  return (
    <CardWrapper
      label="Welcome Back"
      title="Login"
      backButtonHref="/"
      backButtonLabel="Don't have an account? Sign Up"
    >
      <div className="space-y-6">
        <div className="space-y-4">
          <div className="space-y-2">
            <label htmlFor="phoneNumber" className="block text-sm font-medium">
              Phone Number
            </label>
            <input
              type="text"
              id="phoneNumber"
              placeholder="+2519.../ +2517.."
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
        </div>
        <Button type="button" className="w-full">
          Login
        </Button>
      </div>
    </CardWrapper>
  );
};

export default LoginPage;
