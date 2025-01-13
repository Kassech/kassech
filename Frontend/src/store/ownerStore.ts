import {create} from 'zustand';

interface OwnerStore {
  formData: {
    firstName: string;
    lastName: string;
    email: string;
    phoneNumber: string;
    profilePicture: File | null;
    kebeleId: File | null;
    insurance: File | null;
  };
  setField: (key: keyof OwnerStore['formData'], value: any) => void;
  resetForm: () => void;
}

export const useOwnerStore = create<OwnerStore>((set) => ({
  formData: {
    firstName: '',
    lastName: '',
    email: '',
    phoneNumber: '',
    profilePicture: null,
    kebeleId: null,
    insurance: null,
  },
  setField: (key, value) =>
    set((state) => ({
      formData: {
        ...state.formData,
        [key]: value,
      },
    })),
  resetForm: () =>
    set(() => ({
      formData: {
        firstName: '',
        lastName: '',
        email: '',
        phoneNumber: '',
        profilePicture: null,
        kebeleId: null,
        insurance: null,
      },
    })),
}));
