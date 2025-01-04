import { useMutation, useQuery, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

type Role = {
  ID: number;
  RoleName: string;
  Description: string;
  Permission: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
};

// Fetch all role
export const useGetAllRole = () => {
  return useQuery<Role[]>('role', async () => {
    const response = await api.get('/roles/');
    return response.data;
  });
};
// Fetch a role by ID
export const useGetRoleById = (id: number) => {
  return useQuery<Role>(['role', id], async () => {
    const response = await api.get(`/roles/${id}`);
    return response.data;
  });
};

// Create a new role
export const useCreateRole = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (pendingRole: Omit<Role, 'ID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | 'Permission'>) => {
      const response = await api.post('/roles/', pendingRole);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('role');
      },
    }
  );
};

// Update a role by ID
export const useUpdateRole = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async ({ id, updateRole }: { id: number; updateRole: Partial<Role> }) => {
      console.log(id,updateRole)
      const response = await api.put(`/roles/${id}`, updateRole);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('role');
      },
    }
  );
};

// Delete a role by ID
export const useDeleteRole = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: number) => {
      const response = await api.delete(`/roles/${id}`);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('role');
      },
    }
  );
};
