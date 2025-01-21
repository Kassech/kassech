import { useQuery } from 'react-query';
import api from '../api/axiosInstance';

type User = {
  ID: number;
  FirstName: string;
  LastName: string;
  Email: string;
  PhoneNumber: string;
  IsOnline: boolean;
  ProfilePicture: string | null;
  IsVerified: boolean;
  LastLoginDate: string;
  roles: string;
};

type PaginatedResponse = {
  limit: number;
  page: number;
  total_count: number;
  users: User[];
};

type SearchParams = {
  search: string;
  role?: number;
};

export const useSearchUsers = ({ search, role }: SearchParams) => {
  return useQuery<PaginatedResponse>(
    ['users', { search, role }],
    async () => {
      const params: Record<string, string | number | undefined> = {
        search,
        page: 1,
        limit: 5,
        ...(role && { role }),
      };

      const response = await api.get('/users', { params });
      return response.data;
    },
    {
      enabled: Boolean(search), // Only fetch data if search is provided
      keepPreviousData: true, // Keep previous data while new request is loading
    }
  );
};

