import { useMutation, useQuery, useQueryClient } from "react-query";
import api from "../api/axiosInstance";

type Permission = {
  ID: number;
  PermissionName: string;  
  Description: string;
  Role: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
};


// Fetch all stations
export const useGetAllPermission = () => {
  return useQuery<Permission[]>("permissions", async () => {
    const response = await api.get("/permissions");
    return response.data;
  });
};
// Fetch a station by ID
export const useGetRoleById = (id: number) => {
  return useQuery<Permission>(["permission", id], async () => {
    const response = await api.get(`/permission/${id}`);
    return response.data;
  });
};
