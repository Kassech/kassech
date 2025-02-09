'use client';
import { Check, ChevronsUpDown } from 'lucide-react';
import { Button } from '@/components/ui/button';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import { useVehicleStore } from '@/store/vehicleStore';
import { useGetAllVehicles } from '@/services/vehicleService';

interface VehicleSearchProps {
  onVehicleSelect: (id: string, name: string) => void;
}

interface Vehicle {
  ID: number;
  LicenseNumber: string;
  Make: string;
  VIN: string;
}

export function VehicleDropDown({ onVehicleSelect }: VehicleSearchProps) {
  // Use Zustand store
  const {
    search,
    setSearch,
    selectedVehicle,
    setSelectedVehicle,
    open,
    setOpen,
  } = useVehicleStore();

  // Fetch vehicles with role DRIVER_ROLE and search term
  const { data, isLoading, isError } = useGetAllVehicles(search);

  const vehicleList: Vehicle[] = data?.data || [];

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
            {selectedVehicle || 'Select Vehicle'}
            <ChevronsUpDown className="opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-full p-0">
          <input
            type="text"
            placeholder="Search Vehicle..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-full px-2 py-1 border border-gray-300 rounded"
          />
          {isLoading ? (
            <div className="p-2 text-gray-500 w-full">Loading...</div>
          ) : isError ? (
            <div className="p-2 text-red-500 w-full">
              Error fetching vehicles.
            </div>
          ) : vehicleList.length === 0 ? (
            <div className="p-2 text-gray-500 w-full">No Vehicle found.</div>
          ) : (
            <ul>
              {vehicleList.map((vehicle: Vehicle) => (
                <li
                  key={vehicle.ID}
                  className="p-2 cursor-pointer hover:bg-gray-100 text-sm flex w-full"
                  onClick={() => {
                    setSelectedVehicle(
                      `${vehicle.Make} ${vehicle.LicenseNumber}`
                    );
                    onVehicleSelect(
                      vehicle.ID.toString(),
                      `${vehicle.Make} ${vehicle.LicenseNumber}`
                    );
                    setOpen(false);
                  }}
                >
                  {vehicle.Make} {vehicle.LicenseNumber}
                  {selectedVehicle ===
                    `${vehicle.Make} ${vehicle.LicenseNumber}` && (
                    <Check className="ml-auto opacity-100 inline-flex" />
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
