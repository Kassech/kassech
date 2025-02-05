import { useQuery } from 'react-query';
import api from '../api/axiosInstance';

interface Log {
  ID: number;
  LoginTime: string;
  IP: string;
  UserAgent: string;
  UserID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
}



interface ApiResponse<T> {
  meta: {
    status: number;
    message: string;
  };
  data: T;
}

// Fetch total users
export const useTotalUsers = () => {
  return useQuery('totalUsers', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/total-users'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(response.data.meta.message || 'Error fetching total users');
  });
};

// Fetch active users
export const useActiveUsers = () => {
  return useQuery('activeUsers', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/active-users'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(
      response.data.meta.message || 'Error fetching active users'
    );
  });
};

// Fetch total drivers
export const useTotalDrivers = () => {
  return useQuery('totalDrivers', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/total-drivers'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(
      response.data.meta.message || 'Error fetching total drivers'
    );
  });
};

// Fetch total vehicles
export const useTotalVehicles = () => {
  return useQuery('totalVehicles', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/total-vehicles'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(
      response.data.meta.message || 'Error fetching total vehicles'
    );
  });
};

// Fetch active vehicles
export const useActiveVehicles = () => {
  return useQuery('activeVehicles', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/active-vehicles'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(
      response.data.meta.message || 'Error fetching active vehicles'
    );
  });
};

// Fetch total travel logs
export const useTotalTravelLogs = () => {
  return useQuery('totalTravelLogs', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/total-travel-logs'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(
      response.data.meta.message || 'Error fetching total travel logs'
    );
  });
};

// Fetch total routes
export const useTotalRoutes = () => {
  return useQuery('totalRoutes', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/total-routes'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(
      response.data.meta.message || 'Error fetching total routes'
    );
  });
};

// Fetch total stations
export const useTotalStations = () => {
  return useQuery('totalStations', async () => {
    const response = await api.get<ApiResponse<number>>(
      '/analysis/total-stations'
    );
    if (response.data.meta.status === 200) {
      return response.data.data;
    }
    throw new Error(
      response.data.meta.message || 'Error fetching total stations'
    );
  });
};

// Fetch login logs
export const useLoginLogs = () => {
  return useQuery<Log[]>('loginLogs', async () => {
    const response = await api.get<ApiResponse<Log[]>>('/analysis/login-logs');

    if (response.data.meta.status === 200) {
      return response.data.data;
    }

    throw new Error(response.data.meta.message || 'Error fetching login logs');
  });
};



