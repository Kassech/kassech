import { create } from "zustand";

type Role = {
  ID?: number;
  RoleName: string;
  Description: string;
  Permission: string;
};

type RoleStore = {
  roles: Role[];
  currentRole: Role;
  selectedRole: number | null;
  setRoles: (roles: Role[]) => void;
  newRole: Role;
  setCurrentRole: (role: Role) => void;
  updateCurrentRoleField: (field: keyof Role, value: string) => void;
  setNewRole: (role: Role) => void;
  updateNewRoleField: (field: keyof Role, value: string) => void;
  setSelectedRole: (roleId: number | null) => void;
};

type DialogState = {
  isDialogOpen: boolean;
  setDialogOpen: () => void;
  setDialogClose: () => void;
  isAddRoleDialogOpen: boolean;
  setAddRoleDialogOpen: () => void;
  setAddRoleDialogClose: () => void;
  isEditRoleDialogOpen: boolean;
  setEditRoleDialogOpen: () => void;
  setEditRoleDialogClose: () => void;
};

interface CardStore {
  showCard: boolean;
  toggleCard: () => void;
}

export const useCardStore = create<CardStore>((set) => ({
  showCard: false,
  toggleCard: () => set((state) => ({ showCard: !state.showCard })),
}));

// Role Store
export const useRoleStore = create<RoleStore>((set) => ({
  roles: [],
  currentRole: {
    RoleName: "",
    Description: "",
    Permission: "",
  },
  newRole: {
    RoleName: "",
    Description: "",
    Permission:"",
  },
  selectedRole: null,
  setRoles: (roles) => set({ roles }),
  setCurrentRole: (role) => set({ currentRole: role }),
  updateCurrentRoleField: (field, value) =>
    set((state) => ({
      currentRole: { ...state.currentRole, [field]: value },
    })),
  setNewRole: (role) => set({ newRole: role }),
  updateNewRoleField: (field, value) =>
    set((state) => ({
      newRole: { ...state.newRole, [field]: value },
    })),
  setSelectedRole: (roleId) => set({ selectedRole: roleId }),
}));

// Dialog Store
export const useDialogStore = create<DialogState>((set) => ({
  isDialogOpen: false,
  setDialogOpen: () => set({ isDialogOpen: true }),
  setDialogClose: () => set({ isDialogOpen: false }),
  isAddRoleDialogOpen: false,
  setAddRoleDialogOpen: () => set({ isAddRoleDialogOpen: true }),
  setAddRoleDialogClose: () => set({ isAddRoleDialogOpen: false }),
  isEditRoleDialogOpen: false,
  setEditRoleDialogOpen: () => set({ isEditRoleDialogOpen: true }),
  setEditRoleDialogClose: () => set({ isEditRoleDialogOpen: false }),
}));

export default { useRoleStore, useDialogStore };
