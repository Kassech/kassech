import axios from "axios";
import api from "../api/axiosInstance";

// Fetch all stations
export const fetchStations = async () => {
  try {
    const response = await api.get("/station/");
    return response.data; // Returns an array of stations
  } catch (error) {
    console.error("Error fetching stations:", error);
    throw error;
  }
};

// Create a new station
export const createStation = async (stationData: {
  LocationName: string;
  Latitude: number;
  Longitude: number;
}) => {
  try {
    const response = await api.post("/station/", stationData);
    return response.data; // Returns the created station
  } catch (error) {
    console.error("Error creating station:", error);
    throw error;
  }
};

// Update an existing station
export const updateStation = async (
  stationId: number,
  updatedData: {
    LocationName?: string;
    Latitude?: number;
    Longitude?: number;
  }
) => {
  try {
    const response = await api.put(`/station/${stationId}`, updatedData);
    return response.data; // Returns the updated station
  } catch (error) {
    console.error("Error updating station:", error);
    throw error;
  }
};

export const deleteStation = async (id: number) => {
  try {
    const response = await api.delete(`/station/${id}`);
    console.log("Deleted station:", response.data);
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error details:", error.response);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};


