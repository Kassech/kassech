import { useNavigate } from 'react-router-dom';
import api from '../api/axiosInstance';
import {
  LoginCredentials,
  LoginSuccessResponse,
  LoginErrorResponse,
  ApiResponse,
} from '../types/api'; // Import the types
import { useCustomMutation } from './useQueryHelpers';
import { useUserStore } from '@/store/userStore';

export const useLogin = () => {
  const navigate = useNavigate();

  return useCustomMutation<
    LoginCredentials,
    LoginSuccessResponse,
    LoginErrorResponse
  >(async (credentials: LoginCredentials) => {
    const response = await api.post<LoginSuccessResponse>(
      '/login',
      credentials
    );
    console.log('🚀 ~ response.data:', response.data);
    const { accessToken, user }: LoginSuccessResponse = response.data;

    // Save accessToken and user in localStorage
    localStorage.setItem('accessToken', accessToken);
    localStorage.setItem('user', JSON.stringify(user));

    // Save accessToken and user in a global state using Zustand store
    useUserStore.getState().setUser(user);
    // Navigate to dashboard after successful login
    navigate('/dashboard');

    return response.data; // Return the entire API response object
  });
};

export const useLogout = () =>
  useCustomMutation<void, void, Error>(async () => {
    await api.post("/logout");
    // Remove access token and user from localStorage
    localStorage.removeItem('accessToken');
    localStorage.removeItem('user');
  });

export const useAuthCheck = () =>
  useCustomMutation<void, ApiResponse<LoginSuccessResponse>, Error>(
    async () => {
      const token = localStorage.getItem('accessToken');

      if (!token) {
        throw new Error('Token is missing');
      }

      const response = await api.post<ApiResponse<LoginSuccessResponse>>(
        '/validate'
      );

      return response.data; // Return the entire API response object
    }
  );
