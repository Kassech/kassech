import { useMutation, useQuery, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

type Vehicle = {
  id: number;
  vin: string;
  make: string;
  year: string;
  color: string;
  carType: string;
  carPicture: File | null;
  bollo: File | null;
  insurance: File | null;
  libre: File | null;
  ownerId: number;
  createdAt: string;
  updatedAt: string;
};

// Fetch all vehicles
export const useGetAllVehicles = () => {
  return useQuery<Vehicle[]>('vehicles', async () => {
    const response = await api.get('/vehicles/');
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
