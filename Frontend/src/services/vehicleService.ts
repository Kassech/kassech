import { useMutation, useQuery, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';
import { Vehicle } from '@/types/vehicle';

export const useGetAllVehicles = (search?: string) => {
  return useQuery(['vehicleData', search], async () => {
    const response = await api.get('/vehicles/', {
      params: search ? { search } : {},
    });
    console.log('from the function',response); 
    return response.data; 
  });
};


// Fetch a vehicle by ID
export const useGetVehicleById = (id: number | null) => {
  return useQuery<Vehicle>(
    ['vehicle', id],
    async () => {
      const response = await api.get(`/vehicles/${id}`);
      return response.data;
    },
    {
      enabled: id !== null && id !== undefined, // Fetch only if id is valid
    }
  );
};

export const useCreateVehicle = () => {
  return useMutation(async (newVehicle: unknown) => {
    const response = await api.post('/vehicles/', newVehicle);
    return response.data;
  });
};

export const useUpdateVehicle = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async ({ id, formData }: { id: number; formData: FormData }) => {
      const response = await api.put(`/vehicles/${id}`, formData, {
        headers: { 'Content-Type': 'multipart/form-data' },
      });
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('vehicles');
      },
    }
  );
};

export const useDeleteVehicle = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: number) => {
      const response = await api.delete(`/vehicles/${id}`);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('vehicles');
      },
    }
  );
};
