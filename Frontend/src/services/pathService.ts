import { useMutation, useQuery, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';
import { Path } from '@/types/path';


// Fetch all path
export const useGetAllPaths = () => {
  return useQuery<Path[]>('path', async () => {
    const response = await api.get('/path/');
    return response.data;
  });
};

// Fetch a path by ID
export const useGetPathById = (id: number) => {
  return useQuery<Path>(['path', id], async () => {
    const response = await api.get(`/path/${id}`);
    return response.data;
  });
};

// Create a new path
export const useCreatePath = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (newPath: {
      route_id: number;
      path_name: string;
      distance_km: number;
      estimated_time: string;
      is_active: boolean;
    }) => {
      const response = await api.post('/path/', newPath);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('path');
      },
    }
  );
};

// Update a path by ID
export const useUpdatePath = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async ({
      id,
      updatedPath,
    }: {
      id: number;
      updatedPath: Partial<Path>;
    }) => {
      const response = await api.put(`/path/${id}`, updatedPath);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('path');
      },
    }
  );
};

// Delete a path by ID
export const useDeletePath = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: string) => {
      const response = await api.delete(`/path/${id}`);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('path');
      },
    }
  );
};
