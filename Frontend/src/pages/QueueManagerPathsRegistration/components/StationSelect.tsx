// components/StationSelect.tsx
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from '@/components/ui/select';
import { useGetAllStations } from '@/services/stationService';
import { useFormStore } from '@/store/queueManagerPathStore';

export const StationSelect = () => {
  const selectedUserId = useFormStore((state) => state.selectedUserId);
  const { data: stations, isLoading, error } = useGetAllStations();
  console.log('🚀 ~ StationSelect ~ stations:', stations);
  const setSelectedStation = useFormStore((state) => state.setSelectedStation);

  if (!selectedUserId) return null;
  if (error) return <div>Error loading stations</div>;
  if (isLoading) return <div>Loading stations...</div>;

  return (
    <Select onValueChange={setSelectedStation}>
      <SelectTrigger className="w-[300px]">
        <SelectValue placeholder="Select a station" />
      </SelectTrigger>
      <SelectContent>
        {stations.map((location) => (
          <SelectItem key={location.ID} value={String(location.ID)}>
            {location.LocationName}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};
