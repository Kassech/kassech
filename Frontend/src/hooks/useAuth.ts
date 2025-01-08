import { useNavigate } from "react-router-dom";
import api from "../api/axiosInstance";
import {
  LoginCredentials,
  LoginResponse,
  ApiResponse,
} from "../types/api"; // Import the types
import { useCustomMutation } from "./useQueryHelpers";

export const useLogin = () => {
  const navigate = useNavigate();

  return useCustomMutation<
    LoginCredentials,
    ApiResponse<LoginResponse>,
    string
  >(
    async (credentials: LoginCredentials) => {
      const response = await api.post<ApiResponse<LoginResponse>>(
        "/login",
        credentials
      );
      const { accessToken, user } = response.data!;

      // Save accessToken and user in localStorage
      localStorage.setItem("accessToken", accessToken);
      localStorage.setItem("user", JSON.stringify(user));

      // Navigate to dashboard after successful login
      navigate("/dashboard");

      return response.data; // Return the entire API response object
    },
  );
};

export const useLogout = () =>
  useCustomMutation<void, void, Error>(
    async () => {
      // await api.post("/logout");
      // Remove access token and user from localStorage
      localStorage.removeItem("accessToken");
      localStorage.removeItem("user");
      
    },
  );

export const useAuthCheck = () =>
  useCustomMutation<void, ApiResponse<LoginResponse>, Error>(
    async () => {
      const token = localStorage.getItem("accessToken");

      if (!token) {
        throw new Error("Token is missing");
      }

      const response = await api.post<ApiResponse<LoginResponse>>("/validate");

      return response.data; // Return the entire API response object
    },
  );
