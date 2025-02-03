import { useQuery } from 'react-query';
import api from '../api/axiosInstance';

// Fetch total number of users
export const useGetTotalUsers = () => {
  return useQuery<number>('total-users', async () => {
    const response = await api.get('/analysis/total-users');
    return response.data;
  });
};

// Fetch total number of active users
export const useGetActiveUsers = () => {
  return useQuery<number>('active-users', async () => {
    const response = await api.get('/analysis/active-users');
    return response.data;
  });
};

// Fetch total number of drivers
export const useGetTotalDrivers = () => {
  return useQuery<number>('total-drivers', async () => {
    const response = await api.get('/analysis/total-drivers');
    return response.data;
  });
};

// Fetch total number of vehicles
export const useGetTotalVehicles = () => {
  return useQuery<number>('total-vehicles', async () => {
    const response = await api.get('/analysis/total-vehicles');
    return response.data;
  });
};

// Fetch total number of active vehicles
export const useGetActiveVehicles = () => {
  return useQuery<number>('active-vehicles', async () => {
    const response = await api.get('/analysis/active-vehicles');
    return response.data;
  });
};

// Fetch total number of travel logs
export const useGetTotalTravelLogs = () => {
  return useQuery<number>('total-travel-logs', async () => {
    const response = await api.get('/analysis/total-travel-logs');
    return response.data;
  });
};

// Fetch total number of routes
export const useGetTotalRoutes = () => {
  return useQuery<number>('total-routes', async () => {
    const response = await api.get('/analysis/total-routes');
    return response.data;
  });
};

// Fetch total number of stations
export const useGetTotalStations = () => {
  return useQuery<number>('total-stations', async () => {
    const response = await api.get('/analysis/total-stations');
    return response.data;
  });
};

// Fetch total number of login logs
export const useGetLoginLogs = () => {
  return useQuery<number>('login-logs', async () => {
    const response = await api.get('/analysis/login-logs');
    return response.data;
  });
};
