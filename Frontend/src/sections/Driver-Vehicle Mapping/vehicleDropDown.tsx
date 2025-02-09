'use client';
import {useState} from 'react';
import { Check, ChevronsUpDown } from 'lucide-react';
import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandItem,
  CommandList,
} from '@/components/ui/command';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import { useGetAllVehicles } from '@/services/vehicleService';
import { Input } from '@/components/ui/input';
interface VehicleSearchProps {
  onVehicleSelect: (id: string, name: string) => void;
}

export default function VehicleSearch({ onVehicleSelect }: VehicleSearchProps) {
  const [open, setOpen] = useState(false);
  const [search, setSearch] = useState('');
  const [value, setValue] = useState('');

  // Fetch vehicles from API
  const { data: vehicles, isLoading, isError } = useGetAllVehicles(search);

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
            ? vehicles?.find((vehicle) => vehicle.id.toString() === value)
                ?.vin ?? 'Select Vehicle'
            : 'Select Vehicle'}
          <ChevronsUpDown className="opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <Command>
          <Input
            placeholder="Search Vehicle..."
            value={search}
            onChange={(e) => setSearch(e.target.value)} // Corrected event handler
          />
          <CommandList>
            {isLoading ? (
              <CommandEmpty>Loading...</CommandEmpty>
            ) : isError ? (
              <CommandEmpty>Error fetching vehicles.</CommandEmpty>
            ) : vehicles?.length === 0 ? (
              <CommandEmpty>No vehicles found.</CommandEmpty>
            ) : (
              <CommandGroup>
                {vehicles?.map((vehicle) => (
                  <CommandItem
                    key={vehicle.id}
                    value={vehicle.id.toString()}
                    onSelect={(currentValue) => {
                      setValue(currentValue === value ? '' : currentValue);
                      onVehicleSelect(vehicle.id.toString(), vehicle.vin);
                      setOpen(false);
                    }}
                  >
                    {vehicle.vin}
                    <Check
                      className={cn(
                        'ml-auto',
                        value === vehicle.id.toString()
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
