import { useMutation, useQuery, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

type VehicleType = {
  ID: number;
  TypeName: string;
  Capacity: number;
  Description: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
};

// Fetch all vehicle types
export const useGetAllVehicleTypes = () => {
  return useQuery<VehicleType[]>('vehicle_types', async () => {
    const response = await api.get('/vehicle_types/');
    return response.data;
  });
};

// Fetch a vehicle type by ID
export const useGetVehicleTypeById = (id: number | null) => {
  return useQuery<VehicleType>(
    ['vehicle_type', id],
    async () => {
      const response = await api.get(`/vehicle_types/${id}`);
      return response.data;
    },
    {
      enabled: id !== null && id !== undefined,
    }
  );
};

// Create a new vehicle type
export const useCreateVehicleType = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (
      newVehicleType: Omit<
        VehicleType,
        'ID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt'
      >
    ) => {
      const response = await api.post('/vehicle_types/', newVehicleType);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('vehicle_types');
      },
    }
  );
};

// Update a vehicle type by ID
export const useUpdateVehicleType = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async ({
      id,
      updateData,
    }: {
      id: number;
      updateData: Partial<VehicleType>;
    }) => {
      const response = await api.put(`/vehicle_types/${id}`, updateData);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('vehicle_types');
      },
    }
  );
};

// Delete a vehicle type by ID
export const useDeleteVehicleType = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: number) => {
      const response = await api.delete(`/vehicle_types/${id}`);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('vehicle_types');
      },
    }
  );
};
