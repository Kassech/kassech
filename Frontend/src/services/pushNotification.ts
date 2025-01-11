import { useMutation } from 'react-query';
import api from '../api/axiosInstance';

type notification = {
  token: string;
};

export const usePushNotificaiton = () => {
  return useMutation(
    async (newNotification: notification) => {
      const response = await api.post('/notification', newNotification);
      return response.data;
    },
  );
};
