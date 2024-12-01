import { useMutation } from 'react-query';
import api from '../api/axiosInstance';
import { useNavigate } from 'react-router-dom';

export const useLogin = () => {
    const navigate = useNavigate()
    return useMutation(async (credentials: { email_or_phone: string; password: string }) => {
        const response = await api.post('/login', credentials);
        const { accessToken, user } = response.data;
        localStorage.setItem('accessToken', accessToken);
        localStorage.setItem('user', JSON.stringify(user));
        navigate("/dashboard");
        return response.data.user;
    });
}

export const useLogout = () =>
  useMutation(async () => {
    await api.post('/logout');
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
  });

  export const useAuthCheck = () => {
    return useMutation(async () => {
      const token = localStorage.getItem('accessToken');
      if (!token) {
        throw new Error('Token is missing');
      }
      
      const response = await api.post('/validate');
      return response.data;
    });
  };
  