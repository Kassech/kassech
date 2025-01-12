import { useMutation, useQueryClient } from 'react-query';
import api from '../api/axiosInstance';

type Owner = {
  firstName: string;
  lastName: string;
  email: string;
  phoneNumber: string;
  profilePicture: File;
  kebeleId: File;
  insurance: File;
};

export const useCreateOwner = () => {
  const queryClient = useQueryClient();

  return useMutation(
    async (ownerData: Owner) => {
      // Create FormData to handle file uploads
      const formData = new FormData();
      formData.append('firstName', ownerData.firstName);
      formData.append('lastName', ownerData.lastName);
      formData.append('email', ownerData.email);
      formData.append('phoneNumber', ownerData.phoneNumber);
      formData.append('profilePicture', ownerData.profilePicture);
      formData.append('kebeleId', ownerData.kebeleId);
      formData.append('insurance', ownerData.insurance);

      const response = await api.post('/owners/register', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      return response.data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('owner');
      },
    }
  );
};
