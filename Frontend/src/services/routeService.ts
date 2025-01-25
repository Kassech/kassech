import { useMutation, useQuery, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

type Route = {
  id: number;
  locationA: number;
  locationB: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
};

// Fetch all routes
export const useGetAllRoutes = () => {
  return useQuery<Route[]>('routes', async () => {
    const response = await api.get('/routes/');
    return response.data;
  });
};

// Fetch a route by ID
export const useGetRouteById = (id: number) => {
  return useQuery<Route>(['route', id], async () => {
    const response = await api.get(`/routes/${id}`);
    return response.data;
  });
};

// Create a new route
export const useCreateRoute = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (newRoute: { locationA: number; locationB: number }) => {
      const response = await api.post('/routes/', newRoute);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('routes');
      },
    }
  );
};

// Update a route by ID
export const useUpdateRoute = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async ({ id, updatedRoute }: { id: number; updatedRoute: Partial<Route> }) => {
      const response = await api.put(`/routes/${id}`, updatedRoute);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('routes');
      },
    }
  );
};

// Delete a route by ID
export const useDeleteRoute = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: string) => {
      console.log("🚀 ~ id:", id)
      const response = await api.delete(`/routes/${id}`);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('routes');
      },
    }
  );
};
