'use client';

import * as React from 'react';
import { Check, ChevronsUpDown } from 'lucide-react';
import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import { useSearchUsers } from '../../services/carOwnerService';
import LoadingSpinner from '@/components/loading-spinner';

interface OwnerSearchProps {
  onOwnerSelect: (id: string, name: string) => void;
}

export function OwnerSearch({ onOwnerSelect }: OwnerSearchProps) {
  const [open, setOpen] = React.useState(false);
  const [search, setSearch] = React.useState('');
  const [value, setValue] = React.useState('');

  // Fetch users with role 1 and matching the search term
  const { data, isLoading, isError } = useSearchUsers({
    search: search || '',
  });

  if (!data) {
    return <LoadingSpinner />;
  }
  console.log(data);
  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {value
            ? data?.users.find((user) => user.ID.toString() === value)
                ?.FirstName ?? 'Select Owner'
            : 'Select Owner'}
          <ChevronsUpDown className="opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <div>
          <input
            type="text"
            placeholder="Search Owner..."
            value={search}
            onChange={(e) => {
              setSearch(e.target.value);
              console.log('Search:', search);
              console.log('Fetched Users:', data?.users, data?.users.length);
            }}
            className="w-full px-2 py-1 border border-gray-300 rounded"
          />
          <div>
            {isLoading ? (
              <div className="p-2 text-gray-500">Loading...</div>
            ) : isError ? (
              <div className="p-2 text-red-500">Error fetching users.</div>
            ) : !data?.users || data.users.length === 0 ? (
              <div className="p-2 text-gray-500">No owner found.</div>
            ) : (
              <ul>
                {data?.users.map((user) => {
                  console.log('Rendering User:', user);
                  return (
                    <li
                      key={user.ID}
                      className="p-2 cursor-pointer hover:bg-gray-100 text-sm flex"
                      onClick={() => {
                        setValue(user.ID.toString());
                        onOwnerSelect(
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
          </div>
        </div>
      </PopoverContent>
    </Popover>
  );
}
