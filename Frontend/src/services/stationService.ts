import { useMutation, useQuery, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

type Station = {
  id: number;
  locationName: string;
  latitude: number;
  longitude: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
};

// Fetch all stations
export const useGetAllStations = () => {
  return useQuery<Station[]>('stations', async () => {
    const response = await api.get('/station/');
    return response.data;
  });
};
// Fetch a station by ID
export const useGetStationById = (id: number) => {
  return useQuery<Station>(['station', id], async () => {
    const response = await api.get(`/station/${id}`);
    return response.data;
  });
};

// Create a new station
export const useCreateStation = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (newStation: Omit<Station, 'id' | 'createdAt' | 'updatedAt' | 'deletedAt'>) => {
      const response = await api.post('/station', newStation);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('stations');
      },
    }
  );
};

// Update a station by ID
export const useUpdateStation = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async ({ id, updatedStation }: { id: number; updatedStation: Partial<Station> }) => {
      const response = await api.put(`/station/${id}`, updatedStation);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('stations');
      },
    }
  );
};

// Delete a station by ID
export const useDeleteStation = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: number) => {
      const response = await api.delete(`/station/${id}`);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('stations');
      },
    }
  );
};
