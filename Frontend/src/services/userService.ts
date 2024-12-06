import api from '../api/axiosInstance';

export const fetchUserData = async () => {
  const response = await api.get('/users');
  return response.data;
};

export const updateUserData = async (userData: unknown) => {
  const response = await api.put('/users', userData);
  return response.data;
};
