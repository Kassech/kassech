'use client';

import * as React from 'react';
import { Check, ChevronsUpDown } from 'lucide-react';
import { Button } from '@/components/ui/button';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import LoadingSpinner from '@/components/loading-spinner';
import { useFetchUserData } from '@/services/userService';
import { DRIVER_ROLE } from '@/constants';

interface DriverSearchProps {
  onDriverSelect: (id: string, name: string) => void;
}

interface Users {
  ID: number;
  FirstName: string;
  LastName: string;
}

export function DriverDropDown({ onDriverSelect }: DriverSearchProps) {
  const [open, setOpen] = React.useState(false);
  const [search, setSearch] = React.useState('');
  const [value, setValue] = React.useState('');

  // Fetch users with role 1 and matching the search term
  const { data, isLoading, isError } = useFetchUserData({
    search: search || '',
    role: DRIVER_ROLE,
  });
  const userList: Users[] = data?.data || [];

  if (!data) {
    return <LoadingSpinner />;
  }
  console.log(data);
  return (
    <Popover modal open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {value
            ? data?.data?.find((user: Users) => user.ID.toString() === value)
                ?.FirstName ?? 'Select Driver'
            : 'Select Driver'}

          <ChevronsUpDown className="opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
          <input
            type="text"
            placeholder="Search Driver..."
            value={search}
            onChange={(e) => {
              setSearch(e.target.value);
              console.log('Search:', search);
              console.log('Fetched Users:', data?.users, data?.users.length);
            }}
            className="w-full px-2 py-1 border border-gray-300 rounded"
          />
            {isLoading ? (
              <div className="p-2 text-gray-500">Loading...</div>
            ) : isError ? (
              <div className="p-2 text-red-500">Error fetching users.</div>
            ) : !data?.data || data.data.length === 0 ? (
              <div className="p-2 text-gray-500">No Driver found.</div>
            ) : (
              <ul>
                {userList.map((user: Users) => {
                  console.log('Rendering User:', user);
                  return (
                    <li
                      key={user.ID}
                      className="p-2 cursor-pointer hover:bg-gray-100 text-sm flex"
                      onClick={() => {
                        setValue(user.ID.toString());
                        onDriverSelect(
                          user.ID.toString(),
                          `${user.FirstName} ${user.LastName}`
                        );
                        setOpen(false);
                      }}
                    >
                      {user.FirstName} {user.LastName}
                      {value === user.ID.toString() && (
                        <Check className="ml-auto opacity-100 inline-flex" />
                      )}
                    </li>
                  );
                })}
              </ul>
            )}
      </PopoverContent>
    </Popover>
  );
}
