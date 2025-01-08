import { useDriverStore } from "@/store/driverStore";
import { useQuery, useMutation, useQueryClient } from "react-query";

// Types
interface Driver {
    id: number;
    name: string;
}

// API Functions
const fetchDrivers = async (): Promise<Driver[]> => {
    const response = await fetch("/api/drivers");
    if (!response.ok) throw new Error("Failed to fetch drivers");
    return response.json();
};

const addDriverApi = async (driver: Driver) => {
    const response = await fetch("/api/drivers", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(driver),
    });
    if (!response.ok) throw new Error("Failed to add driver");
    return response.json();
};

const removeDriverApi = async (id: number) => {
    const response = await fetch(`/api/drivers/${id}`, { method: "DELETE" });
    if (!response.ok) throw new Error("Failed to remove driver");
};

// Hooks
export const useDrivers = () => {
    const setDrivers = useDriverStore((state) => state.setDrivers);

    return useQuery("drivers", fetchDrivers, {
        onSuccess: (data) => setDrivers(data),
        onError: (error) => console.error(error),
    });
};

export const useAddDriver = () => {
    const queryClient = useQueryClient();
    const addDriver = useDriverStore((state) => state.addDriver);

    return useMutation(addDriverApi, {
        onMutate: async (newDriver: Driver) => {
            addDriver(newDriver); // Optimistic update
        },
        onSuccess: () => queryClient.invalidateQueries("drivers"),
        onError: (error) => console.error(error),
    });
};

export const useRemoveDriver = () => {
    const queryClient = useQueryClient();
    const removeDriver = useDriverStore((state) => state.removeDriver);

    return useMutation(removeDriverApi, {
        onMutate: async (id: number) => {
            removeDriver(id); // Optimistic update
        },
        onSuccess: () => queryClient.invalidateQueries("drivers"),
        onError: (error) => console.error(error),
    });
};
