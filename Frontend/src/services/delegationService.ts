import { useMutation } from 'react-query';
import api from '../api/axiosInstance';

interface DelegationData {
  driver_id: number;
  vehicle_id: number;
}

export const useAssignDriverToVehicle = () => {
  return useMutation(async (data: DelegationData) => {
    const response = await api.post('/driver-deligation/assign', data);

    if (response.status === 200) {
      return response.data;
    }

    throw new Error('Error assigning driver to vehicle');
  });
};
