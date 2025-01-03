import {create} from "zustand";

interface DialogState {
  isDialogOpen: boolean;
  setDialogOpen: () => void;
  setDialogClose: () => void;
  isAddRoleDialogOpen: boolean;
  setAddRoleDialogOpen: () => void;
  setAddRoleDialogClose: () => void;
  isEditRoleDialogOpen: boolean;
  setEditRoleDialogOpen: () => void;
  setEditRoleDialogClose: () => void;
}

const useStore = create<DialogState>((set) => ({
  isDialogOpen: false,
  setDialogOpen: () => set({ isDialogOpen: true }), 
  setDialogClose: () => set({ isDialogOpen: false }), 
}));
export default useStore;