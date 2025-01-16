import { create } from "zustand";
import * as L from "leaflet";

interface StationState {
  position: L.LatLng | null;
  LocationName: string;
  editingStationId: number | null;
  location1: string;
  location2: string;
  setPosition: (position: L.LatLng | null) => void;
  setLocationName: (name: string) => void;
  setEditingStationId: (id: number | null) => void;
  setLocation1: (location: string) => void;
  setLocation2: (location: string) => void;
  clearLocations: () => void;
}

export const useStationStore = create<StationState>((set) => ({
  position: null,
  LocationName: "",
  editingStationId: null,
  location1: "",
  location2: "",
  setPosition: (position) => set({ position }),
  setLocationName: (name) => set({ LocationName: name }),
  setEditingStationId: (id) => set({ editingStationId: id }),
  setLocation1: (location) => set({ location1: location }),
  setLocation2: (location) => set({ location2: location }),
  clearLocations: () => set({ location1: "", location2: "",position: null }),
}));
