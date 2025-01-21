'use client';

import * as React from 'react';
import { Check, ChevronsUpDown } from 'lucide-react';
import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import { useSearchUsers } from '../../services/carOwnerService';

interface DriverSearchProps {
  onDriverSelect: (id: string, name: string) => void;
}

export default function DriverSearch({ onDriverSelect }: DriverSearchProps) {
  const [open, setOpen] = React.useState(false);
  const [search, setSearch] = React.useState('');
  const [value, setValue] = React.useState('');


  // Fetch users with role 1 and matching the search term
  const { data, isLoading, isError } = useSearchUsers({
    search, role:2
  });

  // Extract users from data
  const users = data?.users ?? [];

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
            ? users.find((user) => user.ID.toString() === value)?.FirstName ??
              'Select Driver'
            : 'Select Driver'}
          <ChevronsUpDown className="opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <Command>
          <CommandInput
            placeholder="Search Driver..."
            value={search}
            onValueChange={(value) => setSearch(value)}
          />
          <CommandList>
            {isLoading ? (
              <CommandEmpty>Loading...</CommandEmpty>
            ) : isError ? (
              <CommandEmpty>Error fetching driver.</CommandEmpty>
            ) : users.length === 0 ? (
              <CommandEmpty>No driver found.</CommandEmpty>
            ) : (
              <CommandGroup>
                {users.map((user) => (
                  <CommandItem
                    key={user.ID}
                    value={user.ID.toString()}
                    onSelect={(currentValue) => {
                      setValue(currentValue === value ? '' : currentValue);
                      onDriverSelect(
                        user.ID.toString(),
                        `${user.FirstName} ${user.LastName}`
                      );
                      setOpen(false);
                    }}
                  >
                    {user.FirstName} {user.LastName}
                    <Check
                      className={cn(
                        'ml-auto',
                        value === user.ID.toString()
                          ? 'opacity-100'
                          : 'opacity-0'
                      )}
                    />
                  </CommandItem>
                ))}
              </CommandGroup>
            )}
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
