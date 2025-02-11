import { useQuery } from 'react-query';
import useWebSocketStore from '@/store/useWebSocketStore';

const useWebSocketData = () => {
  const { vehicles, paths, nearbyCars, allCars } = useWebSocketStore();

  return useQuery('websocketData', () => {
    return {
      vehicles,
      paths,
      nearbyCars,
      allCars,
    };
  });
};

export default useWebSocketData;
