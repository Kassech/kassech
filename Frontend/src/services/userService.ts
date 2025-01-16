import { useQuery, useMutation } from 'react-query';
import api from '../api/axiosInstance';

// Fetch user data with react-query
export const useFetchUserData = () => {
  return useQuery('userData', async () => {
    const response = await api.get('/users');
    return response.data;
  });
};

// Update user data with react-query
export const useUpdateUserData = () => {
  return useMutation(
    async (userData: unknown) => {
      const response = await api.put('/users', userData);
      return response.data;
    }
  );
};

// Create a new user
export const useCreateUser = () => {
  return useMutation(
    async (newUser: unknown) => {
      const response = await api.post('/users', newUser);
      return response.data;
    }
  );
};
