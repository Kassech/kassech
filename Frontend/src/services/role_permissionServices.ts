import { useMutation, useQueryClient } from "react-query";
import api from "../api/axiosInstance";

type RolePermissionPayload = {
  RoleID: number;
  PermissionID: number;
};

export const useCreateRolePermission = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (newPermission: RolePermissionPayload) => {
      const response = await api.post("/role_permissions", newPermission);
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("role_permissions");
        queryClient.invalidateQueries("roles");
      },
    }
  );
};
