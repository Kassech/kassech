'use client';
import { Check, ChevronsUpDown } from 'lucide-react';
import { Button } from '@/components/ui/button';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import { useFetchUserData } from '@/services/userService';
import { DRIVER_ROLE } from '@/constants';
import { useDriverStore } from '@/store/useDriverStore';

interface DriverSearchProps {
  onDriverSelect: (id: string, name: string) => void;
}

interface Users {
  id: number;
  first_name: string;
  last_name: string;
}

export function DriverDropDown({ onDriverSelect }: DriverSearchProps) {
  // Use Zustand store
  const {
    search,
    setSearch,
    selectedDriver,
    setSelectedDriver,
    open,
    setOpen,
  } = useDriverStore();

  // Fetch users with role DRIVER_ROLE and search term
  const { data, isLoading, isError } = useFetchUserData({
    search,
    role: DRIVER_ROLE,
  });

  const userList: Users[] = data?.data || [];

  return (
    <div className="w-full">
      <Popover modal open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            className="w-full justify-between"
          >
            {selectedDriver || 'Select Driver'}
            <ChevronsUpDown className="opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-full p-0 max-h-60 overflow-y-auto">
          <input
            type="text"
            placeholder="Search Driver..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-full px-2 py-1 border border-gray-300 rounded sticky top-0 bg-white"
          />
          {isLoading ? (
            <div className="p-2 text-gray-500 w-full">Loading...</div>
          ) : isError ? (
            <div className="p-2 text-red-500 w-full">Error fetching users.</div>
          ) : userList.length === 0 ? (
            <div className="p-2 text-gray-500 w-full">No Driver found.</div>
          ) : (
            <ul className="divide-y">
              {userList.map((user: Users) => (
                <li
                  key={user.id}
                  className="p-2 cursor-pointer hover:bg-gray-100 text-sm flex  items-center justify-between"
                  onClick={() => {
                    setSelectedDriver(`${user.first_name} ${user.last_name}`);
                    onDriverSelect(
                      user.id.toString(),
                      `${user.first_name} ${user.last_name}`
                    );
                    setOpen(false);
                  }}
                >
                  {user.first_name} {user.last_name}
                  {selectedDriver ===
                    `${user.first_name} ${user.last_name}` && (
                    <Check className="h-4 w-4 ml-auto opacity-100 inline-flex" />
                  )}
                </li>
              ))}
            </ul>
          )}
        </PopoverContent>
      </Popover>
    </div>
  );
}
