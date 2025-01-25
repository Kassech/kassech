import { create } from 'zustand'

interface RoutePoint {
  name: string
  lat: number
  lng: number
}

interface RouteStore {
  selectedRoute: {
    start: RoutePoint
    end: RoutePoint
    pathName: string
  } | null
  distanceKm: number
  estimatedTime: string
  setSelectedRoute: (start: RoutePoint, end: RoutePoint, pathName: string) => void
  setCalculations: (distance: number, time: string) => void
  reset: () => void
}

export const useRouteStore = create<RouteStore>((set) => ({
  selectedRoute: null,
  distanceKm: 0,
  estimatedTime: '',
  setSelectedRoute: (start, end, pathName) => set({ selectedRoute: { start, end, pathName } }),
  setCalculations: (distanceKm, estimatedTime) => set({ distanceKm, estimatedTime }),
  reset: () => set({ selectedRoute: null, distanceKm: 0, estimatedTime: '' })
}))
