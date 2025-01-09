import { create } from "zustand";

// Types
interface Driver {
    id: number;
    name: string;
}

// Zustand Store
interface DriverStore {
    drivers: Driver[];
    setDrivers: (drivers: Driver[]) => void;
    addDriver: (driver: Driver) => void;
    removeDriver: (id: number) => void;
}

export const useDriverStore = create<DriverStore>((set) => ({
    drivers: [],
    setDrivers: (drivers) => set({ drivers }),
    addDriver: (driver) => set((state) => ({ drivers: [...state.drivers, driver] })),
    removeDriver: (id) => set((state) => ({ drivers: state.drivers.filter(driver => driver.id !== id) })),
}));
