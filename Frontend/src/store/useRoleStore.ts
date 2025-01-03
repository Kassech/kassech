import { create } from "zustand";

type Role = {
  id?: number;
  roleName: string;
  description: string;
  permission: string;
};

type RoleStore = {
  roles: Role[];
  currentRole: Role;
  isModalOpen: boolean;
  setRoles: (roles: Role[]) => void;
  setCurrentRole: (role: Role) => void;
  toggleModal: (open: boolean) => void;
};

export const useRoleStore = create<RoleStore>((set) => ({
  roles: [],
  currentRole: { roleName: "", description: "", permission: "" },
  isModalOpen: false,
  setRoles: (roles) => set({ roles }),
  setCurrentRole: (role) => set({ currentRole: role }),
  toggleModal: (open) => set({ isModalOpen: open }),
}));
