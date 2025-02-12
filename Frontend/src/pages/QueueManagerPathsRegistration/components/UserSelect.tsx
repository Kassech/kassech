// components/UserSelect.tsx
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from '@/components/ui/select';
import { QUEUE_MANAGER_ROLE } from '@/constants';
import { useFetchUserData } from '@/services/userService';
import { useFormStore } from '@/store/queueManagerPathStore';

export const UserSelect = () => {
  const {
    data: users,
    isLoading,
    error,
  } = useFetchUserData({ role: QUEUE_MANAGER_ROLE });
  console.log('🚀 ~ UserSelect ~ users:', users);
  const setSelectedUserId = useFormStore((state) => state.setSelectedUserId);

  if (error) return <div>Error loading users</div>;
  if (isLoading) return <div>Loading users...</div>;

  return (
    <Select onValueChange={setSelectedUserId}>
      <SelectTrigger className="w-[300px]">
        <SelectValue placeholder="Select a Queue Manager" />
      </SelectTrigger>
      <SelectContent>
        {users?.data?.map((user) => (
          <SelectItem key={user.id} value={user.id.toString()}>
            {user.first_name} {user.last_name}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};
