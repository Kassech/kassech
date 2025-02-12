import { create } from 'zustand';

interface DialogState {
  isDialogOpen: boolean;
  isEditDialogOpen: boolean;
  selectedUser: number | null;
  selectedUserRole: string ;
  setDialogOpen: (userId: number) => void;
  setDialogClose: () => void;
  setEditDialogOpen: (userId: number, userRole: string) => void;
  setEditDialogClose: () => void;
}


export const userManagingStore = create<DialogState>((set) => ({
  isDialogOpen: false,
  selectedUser: null,
  isEditDialogOpen: false,
  selectedUserRole: '',
  setDialogOpen: (userId) => set({ isDialogOpen: true, selectedUser: userId }),
  setDialogClose: () => set({ isDialogOpen: false, selectedUser: null }),
  setEditDialogOpen: (userId, userRole) =>
    set({ isEditDialogOpen: true, selectedUser: userId , selectedUserRole: userRole }),
  setEditDialogClose: () =>
    set({ isEditDialogOpen: false, selectedUser: null, selectedUserRole: '' }),


}));
