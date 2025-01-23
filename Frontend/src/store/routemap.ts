import { create } from 'zustand';

const useFormStore = create((set) => ({
  location1: '',
  location2: '',
  loading: false,
  setLocation1: (location) => set({ location1: location }),
  setLocation2: (location) => set({ location2: location }),
  setLoading: (loading) => set({ loading }),
  resetForm: () => set({ location1: '', location2: '' }),
}));

export default useFormStore;
