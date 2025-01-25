import { create } from 'zustand';

interface FormState {
  selectedUserId: string | null;
  selectedStation: string | null;
  selectedPaths: string[];
  setSelectedUserId: (id: string | null) => void;
  setSelectedStation: (id: string | null) => void;
  setSelectedPaths: (paths: string[]) => void;
}

export const useFormStore = create<FormState>((set) => ({
  selectedUserId: null,
  selectedStation: null,
  selectedPaths: [],
  setSelectedUserId: (id) => set({ selectedUserId: id }),
  setSelectedStation: (id) => set({ selectedStation: id }),
  setSelectedPaths: (paths) => set({ selectedPaths: paths }),
}));
