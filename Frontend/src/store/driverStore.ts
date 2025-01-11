import { create } from 'zustand';
import { persist } from 'zustand/middleware';

interface DriverStoreState {
  formData: Record<string, any>;
  setField: (field: string, value: any) => void;
  resetForm: () => void;
  clearStore: () => void;
}

export const useDriverStore = create<DriverStoreState>()(
  persist(
    (set) => ({
      formData: (() => {
        // Check if there is formData in localStorage and return it
        const savedData = typeof window !== 'undefined' ? localStorage.getItem('driver-store') : null;
        // If there is saved data, parse it; otherwise, return an empty object
        return savedData ? JSON.parse(savedData).formData : {};
      })(),

      setField: (field, value) =>
        set((state) => {
          const updatedFormData = { ...state.formData, [field]: value };

          // Stringify and save the updated formData to localStorage directly
          localStorage.setItem('driver-store', JSON.stringify({ formData: updatedFormData }));

          return { formData: updatedFormData };
        }),

      resetForm: () =>
        set(() => {
          // Remove formData from both the store and localStorage
          localStorage.removeItem('driver-store');
          return { formData: {} };
        }),

      clearStore: () =>
        set(() => {
          // Remove formData from both the store and localStorage
          localStorage.removeItem('driver-store');
          return { formData: {} };
        }),
    }),
    {
      name: 'driver-store', // Key for localStorage
      storage: {
        getItem: (key) => {
          if (typeof window !== 'undefined') {
            const storedValue = localStorage.getItem(key);
            // Parse and return the stored value if it exists
            return storedValue ? JSON.parse(storedValue) : null;
          }
          return null;
        },
        setItem: (key, value) => {
          if (typeof window !== 'undefined') {
            // Stringify the value before saving it to localStorage
            localStorage.setItem(key, JSON.stringify(value));
          }
        },
        removeItem: (key) => {
          if (typeof window !== 'undefined') {
            localStorage.removeItem(key); // Remove the item from localStorage
          }
        },
      },
    }
  )
);
