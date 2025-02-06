import { useQuery, useMutation, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';
import { useTotalUsers } from './dashboardService';

type SearchParams = {
  search?: string;
  role?: number;
  ID?:number;
};


// Fetch user data with react-query
export const useFetchUserData = ({ search = '', role , ID}: SearchParams = {}) => {
  const totalUsers = useTotalUsers()?.data;
  return useQuery(['userData', search, role,ID], async () => {
    const params: Record<string, string | number | undefined> = {
      search: search || '',
      page: 1,
      limit: totalUsers,
      role,
      ID,
    };
    const response = await api.get('/users', { params });
    return response.data;
  });
};

export const useGetUserById = (id: number | null) => {
  return useQuery(
    ['user', id],
    async () => {
      const response = await api.get(`/users/${id}`);
      return response.data;
    },
    {
      enabled: id !== null && id !== undefined,
    }
  );
};

// Update user data with react-query
export const useUpdateUserData = () => {
  return useMutation(async ({id,userData}:{id: string;userData:unknown}) => {
    const response = await api.put(`/users/${id}`, userData);
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
