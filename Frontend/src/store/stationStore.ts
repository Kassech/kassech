import {create} from "zustand";
import * as L from "leaflet";

interface Station {
  ID: number;
  LocationName: string;
  Latitude: number;
  Longitude: number;
}

interface StationState {
  position: L.LatLng | null;
  LocationName: string;
  editingStationId: number | null;
  setPosition: (position: L.LatLng | null) => void;
  setLocationName: (name: string) => void;
  setEditingStationId: (id: number | null) => void;
}

export const useStationStore = create<StationState>((set) => ({
  position: null,
  LocationName: "",
  editingStationId: null,
  setPosition: (position) => set({ position }),
  setLocationName: (name) => set({ LocationName: name }),
  setEditingStationId: (id) => set({ editingStationId: id }),
}));
