import { create } from 'zustand';

interface DriverStore {
  search: string;
  setSearch: (search: string) => void;
  selectedDriver: string;
  setSelectedDriver: (driver: string) => void;
  open: boolean;
  setOpen: (open: boolean) => void;
}

export const useDriverStore = create<DriverStore>((set) => ({
  search: '',
  setSearch: (search) => set({ search }),
  selectedDriver: '',
  setSelectedDriver: (driver) => set({ selectedDriver: driver }),
  open: false,
  setOpen: (open) => set({ open }),
}));
