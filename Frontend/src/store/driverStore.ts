// stores/driverStore.ts
import { create } from 'zustand';

interface DriverStore {
  formData: Record<string, any>;
  setField: (key: string, value: any) => void;
  resetForm: () => void;
}

export const useDriverStore = create<DriverStore>((set) => ({
  formData: {},

  setField: (key, value) =>
    set((state) => ({
      formData: {
        ...state.formData,
        [key]: value,
      },
    })),

  resetForm: () => set({ formData: {} }),
}));
