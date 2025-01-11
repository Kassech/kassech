import { User } from '@/types/user';
import { create } from 'zustand';

// Define the store interface
interface UserStore {
    user: User | null;
    setUser: (user: User) => void;
    clearUser: () => void;
    updateUser: (user: Partial<User>) => void;
}

// Create the Zustand store with persistence
export const useUserStore = create<UserStore>((set) => {
    // Retrieve persisted user data from localStorage (if any)
    const persistedUser = localStorage.getItem('user');

    return {
        // Initialize state from localStorage or default to null
        user: persistedUser ? JSON.parse(persistedUser) : null,

        setUser: (user) => {
            // Save user data to localStorage when it's set
            localStorage.setItem('user', JSON.stringify(user));
            set({ user });
        },

        clearUser: () => {
            // Remove user data from localStorage when it's cleared
            localStorage.removeItem('user');
            set({ user: null });
        },

        updateUser: (user) => set((state) => {
            const updatedUser = state.user ? { ...state.user, ...user } : null;
            if (updatedUser) {
                // Save the updated user data to localStorage
                localStorage.setItem('user', JSON.stringify(updatedUser));
            } else {
                // Clear user data from localStorage if there's no user
                localStorage.removeItem('user');
            }
            return { user: updatedUser };
        })
    };
});
