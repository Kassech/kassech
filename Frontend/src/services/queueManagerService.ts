// services/queueManagerService.ts
import { useQuery, useMutation, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

export type QueueManagerPath = {
  id: number;
  userId: number;
  stationId: number;
  pathsIds: string[];
};

export const useFetchQueueManagerPaths = () => {
  return useQuery<QueueManagerPath[]>(
    ['queueManagerPaths'],
    async () => {
      const response = await api.get('/queue-manager-routes/');
      return response.data;
    }
  );
};

export const useCreateQueueManagerPath = () => {
  const queryClient = useQueryClient();

  return useMutation(
    (newPath: Omit<QueueManagerPath, 'id'>) =>
      api.post('/queue-manager-routes/', newPath),
    {
      onSuccess: () => {
        queryClient.invalidateQueries('queueManagerPaths');
      },
    }
  );
};

export const useDeleteQueueManagerRoute = () => {
    const queryClient = useQueryClient();

    return useMutation(
      (id: number) => api.delete(`/queue-manager-routes/${id}`),
      {
        onSuccess: () => {
          queryClient.invalidateQueries('queueManagerPaths');
        },
      }
    );
  };
