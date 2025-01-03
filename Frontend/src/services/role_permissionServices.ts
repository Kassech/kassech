import { useMutation, useQuery, useQueryClient } from "react-query";
import api from "../api/axiosInstance";

type RolePermission = {
  id: number;
  PermissionName: string;
  description: string;
  Roles: string;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
};


// Fetch all role
export const useGetAllRole = () => {
  return useQuery<RolePermission[]>("role", async () => {
    const response = await api.get("/role_permissions/");
    return response.data;
  });
};
// Fetch a role by ID
export const useGetRoleById = (id: number) => {
  return useQuery<RolePermission>(["role", id], async () => {
    const response = await api.get(`/role_permissions/${id}`);
    return response.data;
  });
};

// Create a new role
export const useCreateRole = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (
      newRole: Omit<RolePermission, "id" | "createdAt" | "updatedAt" | "deletedAt">
    ) => {
      const response = await api.post("/role_permissions", newRole);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("role");
      },
    }
  );
};

// Update a role by ID
export const useUpdateRole = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async ({ id, updateRole }: { id: number; updateRole: Partial<RolePermission> }) => {
      const response = await api.put(`/roles/${id}`, updateRole);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("role");
      },
    }
  );
};

// Delete a role by ID
export const useDeleteRole = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (id: number) => {
      const response = await api.delete(
        `/role_permissions/role/${id}/permission/${id}`
      );
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("role");
      },
    }
  );
};
