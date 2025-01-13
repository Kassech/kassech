import { create } from "zustand";

// TypeScript interface for Route
interface Route {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
  locationA: number;
  locationB: number;
}

// Zustand store state and actions for routes
interface RouteStoreState {
  routes: Route[];
  addRoute: (route: Route) => void;
  updateRoute: (id: number, updatedRoute: Partial<Route>) => void;
  deleteRoute: (id: number) => void;
}

const useRouteStore = create<RouteStoreState>((set) => ({
  routes: [],
  addRoute: (route) =>
    set((state) => ({
      routes: [...state.routes, route],
    })),
  updateRoute: (id, updatedRoute) =>
    set((state) => ({
      routes: state.routes.map((route) =>
        route.id === id ? { ...route, ...updatedRoute } : route
      ),
    })),
  deleteRoute: (id) =>
    set((state) => ({
      routes: state.routes.filter((route) => route.id !== id),
    })),
}));

export default useRouteStore;
