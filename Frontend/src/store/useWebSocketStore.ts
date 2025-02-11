import {create} from 'zustand';

interface WebSocketState {
  socket: WebSocket | null;
  vehicles: any[];
  paths: any[];
  nearbyCars: any[];
  allCars: any[];
  connect: (token: string) => void;
  subscribe: (message: any) => void;
  disconnect: () => void;
}

const useWebSocketStore = create<WebSocketState>((set, get) => ({
  socket: null,
  vehicles: [],
  paths: [],
  nearbyCars: [],
  allCars: [],
  connect: (token) => {
    const socket = new WebSocket(
      `ws://localhost:5000/ws/location?token=${token}`
    );
    socket.onopen = () => {
      console.log('WebSocket connected');
      set({ socket });
    };
    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      switch (data.type) {
        case 'vehicle':
          set({ vehicles: [...get().vehicles, data] });
          break;
        case 'path':
          set({ paths: [...get().paths, data] });
          break;
        case 'nearby':
          set({ nearbyCars: [...get().nearbyCars, data] });
          break;
        case 'all':
          set({ allCars: [...get().allCars, data] });
          break;
        default:
          break;
      }
    };
    socket.onclose = () => {
      console.log('WebSocket disconnected');
      set({ socket: null });
    };
  },
  subscribe: (message) => {
    const { socket } = get();
    if (socket) {
      socket.send(JSON.stringify(message));
    }
  },
  disconnect: () => {
    const { socket } = get();
    if (socket) {
      socket.close();
      set({ socket: null });
    }
  },
}));

export default useWebSocketStore;
