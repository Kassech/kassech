import api from '../api/axiosInstance';

export const fetchUserData = async () => {
  const response = await api.get('/user');
  return response.data;
};

export const updateUserData = async (userData: unknown) => {
  const response = await api.put('/user', userData);
  return response.data;
};
