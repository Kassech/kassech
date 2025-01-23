import { useQuery, useMutation, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

type SearchParams = {
  search?: string;
  role?: number;
};

// Fetch user data with react-query
export const useFetchUserData = ({ search = '', role }: SearchParams = {}) => {
  return useQuery(['userData', search, role], async () => {
    const params: Record<string, string | number | undefined> = {
      search: search || '',
      page: 1,
      limit: 5,
      role,
    };
    const response = await api.get('/users', { params });
    return response.data;
  });
};

// Update user data with react-query
export const useUpdateUserData = () => {
  return useMutation(async (userData: unknown) => {
    const response = await api.put('/users', userData);
    return response.data;
  });
};

// Create a new user
export const useCreateUser = () => {
  return useMutation(async (newUser: unknown) => {
    const response = await api.post('/users', newUser);
    return response.data;
  });
};

export const useDeleteUser = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: number) => {
      const response = await api.delete(`/users/${id}`);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('userData');
      },
    }
  );
};


export const useVerifyUser = () => {
  return useMutation((data: { id: number; state: boolean }) => {
    return api.get(`/users/verify/${data.id}?state=${data.state}`);
  });
};
