// components/PathMultiSelect.tsx
import { MultiSelect } from '@/components/ui/multi-select';
import { useGetPathsByStationID } from '@/services/pathService';
import { useGetAllStations } from '@/services/stationService';
import { useFormStore } from '@/store/queueManagerPathStore';

export const PathMultiSelect = () => {
  const selectedStation = useFormStore((state) => state.selectedStation);
  console.log('🚀 ~ PathMultiSelect ~ selectedStation:', selectedStation);
  const {
    data: paths,
    isLoading,
    error,
  } = useGetPathsByStationID(selectedStation ? selectedStation : null);
  console.log('🚀 ~ PathMultiSelect ~ paths:', paths);
  const setSelectedPaths = useFormStore((state) => state.setSelectedPaths);

  if (!selectedStation) return null;
  if (error) return <div>Error loading paths</div>;
  if (isLoading || !paths) return <div>Loading paths...</div>;

  return (
    <MultiSelect
      options={paths?.map((path) => ({
        value: path.ID,
        label: path.path_name,
      }))}
      onValueChange={setSelectedPaths}
      placeholder="Select paths"
      variant="inverted"
      animation={2}
      maxCount={6}
    />
  );
};
