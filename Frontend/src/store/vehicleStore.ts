import { create } from 'zustand';

interface VehicleStore {
  search: string;
  setSearch: (search: string) => void;
  selectedVehicle: string;
  setSelectedVehicle: (driver: string) => void;
  open: boolean;
  setOpen: (open: boolean) => void;
}

export const useVehicleStore = create<VehicleStore>((set) => ({
  search: '',
  setSearch: (search) => set({ search }),
  selectedVehicle: '',
  setSelectedVehicle: (vehicle) => set({ selectedVehicle: vehicle }),
  open: false,
  setOpen: (open) => set({ open }),
}));
