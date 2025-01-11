import api from '@/api/axiosInstance';
import { useCustomMutation, useCustomQuery } from '@/hooks/useQueryHelpers';

// CRUD Operations for Driver

// Create Driver
export const createDriver = async (data: any) => {
  const formData = new FormData();
  Object.keys(data).forEach(key => {
    formData.append(key, data[key]);
  });
  const response = await api.post('/users', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }); // Post request to create a driver
  return response.data; // Standardized response data
};

// Get Driver
export const getDriver = async (driverId: string) => {
  const response = await api.get(`/users/${driverId}`); // Get request to fetch a specific driver
  return response.data; // Standardized response data
};

// Update Driver
export const updateDriver = async (driverId: string, data: any) => {
  const response = await api.put(`/users/${driverId}`, data); // Put request to update driver
  return response.data; // Standardized response data
};

// Delete Driver
export const deleteDriver = async (driverId: string) => {
  const response = await api.delete(`/users/${driverId}`); // Delete request to remove a driver
  return response.data; // Standardized response data
};

// Custom Hook to handle Create Driver Mutation
export const useCreateDriver = () => {
  return useCustomMutation(createDriver, {
    onSuccess: (data) => {
      console.log('Driver created successfully:', data);
    },
    onError: (error) => {
      console.error('Error creating driver:', error);
    },
  });
};

// Custom Hook to handle Update Driver Mutation
export const useUpdateDriver = () => {
  return useCustomMutation(updateDriver, {
    onSuccess: (data) => {
      console.log('Driver updated successfully:', data);
    },
    onError: (error) => {
      console.error('Error updating driver:', error);
    },
  });
};

// Custom Hook to handle Delete Driver Mutation
export const useDeleteDriver = () => {
  return useCustomMutation(deleteDriver, {
    onSuccess: (data) => {
      console.log('Driver deleted successfully:', data);
    },
    onError: (error) => {
      console.error('Error deleting driver:', error);
    },
  });
};
