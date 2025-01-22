import { create } from 'zustand';

interface DialogState {
  isDialogOpen: boolean;
  isEditDialogOpen: boolean;
  selectedUser: number | null;
  setDialogOpen: (userId: number) => void;
  setDialogClose: () => void;
  setEditDialogOpen: (userId: number) => void;
  setEditDialogClose: () => void;
  confirmDelete: (userId: number) => void;
}

export const userManagingStore = create<DialogState>((set) => ({
  isDialogOpen: false,
  selectedUser: null,
  isEditDialogOpen: false,

  setDialogOpen: (userId) => set({ isDialogOpen: true, selectedUser: userId }),
  setDialogClose: () => set({ isDialogOpen: false, selectedUser: null }),
  setEditDialogOpen: (userId) =>
    set({ isEditDialogOpen: true, selectedUser: userId }),
  setEditDialogClose: () =>
    set({ isEditDialogOpen: false, selectedUser: null }),

  confirmDelete: (userId) => {
    console.log('User deleted:', userId);
    set({ isDialogOpen: false, selectedUser: null });
  },
}));
