import { useState, useMemo } from 'react';
import {
  MapContainer,
  TileLayer,
  Circle,
  Polyline,
  CircleMarker,
  useMapEvents,
  Popup,
} from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import { Card } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Calendar } from '@/components/ui/calendar';
import { Vehicle } from '@/types/vehicle';
import { Path } from '@/types/path';
import api from '@/api/axiosInstance';
import { MultiSelect } from '@/components/ui/multi-select';
import { Input } from '@/components/ui/input';

interface MapDashboardProps {
  vehicles: Vehicle[];
  paths: Path[];
}

const colors = [
  '#FF6B6B',
  '#4ECDC4',
  '#45B7D1',
  '#96CEB4',
  '#FFEEAD',
  '#D4A5A5',
  '#88D8B0',
  '#FF9999',
  '#A44A3F',
  '#F19A3E',
  '#3E5C76',
  '#2F4858',
];

const getVehicleColor = (vehicleId: number) =>
  colors[vehicleId % colors.length];

function MapClickHandler({
  onMapClick,
}: {
  onMapClick: (lat: number, lon: number) => void;
}) {
  useMapEvents({
    click: (e) => {
      onMapClick(e.latlng.lat, e.latlng.lng);
    },
  });
  return null;
}

export function VehicleMapDashboard({ vehicles, paths }: MapDashboardProps) {
  const [filters, setFilters] = useState({
    vehicleId: '',
    startTime: '',
    endTime: '',
    pathId: '',
    radius: '',
    lat: '',
    lon: '',
  });

  const [gpsLogs, setGpsLogs] = useState<any[]>([]);

  const groupedLogs = useMemo(() => {
    const groups: Record<number, any[]> = {};
    gpsLogs.forEach((log) => {
      const vehicleId = log.vehicle_id;
      groups[vehicleId] = [...(groups[vehicleId] || []), log];
    });

    Object.values(groups).forEach((logs) =>
      logs.sort(
        (a, b) =>
          new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
      )
    );

    return groups;
  }, [gpsLogs]);

  const fetchGPSData = async () => {
    try {
      const params = new URLSearchParams();
      Object.entries(filters).forEach(([key, value]) => {
        if (value) params.append(key, value.toString());
      });

      const response = await api.get(`/vehicles/gps-logs?${params}`);
      setGpsLogs(response.data.data);
    } catch (error) {
      console.error('Error fetching GPS data:', error);
    }
  };

  const handleFilterChange = (
    name: string,
    value: string | Date | string[]
  ) => {
    setFilters((prev) => ({
      ...prev,
      [name]: Array.isArray(value) ? value.join(',') : value.toString(),
    }));
  };

  const handleMapClick = (lat: number, lon: number) => {
    setFilters((prev) => ({
      ...prev,
      lat: lat.toString(),
      lon: lon.toString(),
    }));
  };

  return (
    <div className="m-2 flex h-screen">
      <Card className="w-150 p-4 overflow-y-auto space-y-4">
        <h2 className="text-xl font-bold">Vehicle Tracking</h2>

        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium mb-1">Vehicles</label>
            <MultiSelect
              options={vehicles.map((vehicle) => ({
                value: vehicle.ID.toString(),
                label: `${vehicle.LicenseNumber} (${vehicle.Type.TypeName})`,
              }))}
              selected={filters.vehicleId.split(',')}
              onChange={(values) => handleFilterChange('vehicleId', values)}
              placeholder="Select vehicles..."
            />
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">Paths</label>
            <MultiSelect
              options={paths.map((path) => ({
                value: path.ID.toString(),
                label: path.path_name,
              }))}
              selected={filters.pathId.split(',')}
              onChange={(values) => handleFilterChange('pathId', values)}
              placeholder="Select paths..."
            />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium mb-1">
                Start Time
              </label>
              <Calendar
                mode="single"
                selected={
                  filters.startTime ? new Date(filters.startTime) : undefined
                }
                onSelect={(date) =>
                  date && handleFilterChange('startTime', date.toISOString())
                }
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">End Time</label>
              <Calendar
                mode="single"
                selected={
                  filters.endTime ? new Date(filters.endTime) : undefined
                }
                onSelect={(date) =>
                  date && handleFilterChange('endTime', date.toISOString())
                }
              />
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">
              Radius (meters)
            </label>
            <Input
              type="number"
              value={filters.radius}
              onChange={(e) => handleFilterChange('radius', e.target.value)}
            />
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">
              Selected Location
            </label>
            <div className="text-sm">
              Lat: {filters.lat || 'N/A'}, Lon: {filters.lon || 'N/A'}
            </div>
          </div>

          <Button onClick={fetchGPSData} className="w-full">
            Apply Filters
          </Button>
        </div>
      </Card>

      <div className="flex-1">
        <MapContainer
          center={[9.145, 40.4897]}
          zoom={12}
          className="h-full w-full"
          attributionControl={false}
        >
          <TileLayer url="https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png" />
          <MapClickHandler onMapClick={handleMapClick} />

          {filters.radius && filters.lat && filters.lon && (
            <Circle
              center={[parseFloat(filters.lat), parseFloat(filters.lon)]}
              radius={parseFloat(filters.radius)}
              color="blue"
              fillColor="blue"
              fillOpacity={0.1}
            />
          )}

          {Object.entries(groupedLogs).map(([vehicleId, logs]) => {
            const vehicle = vehicles.find((v) => v.ID === Number(vehicleId));
            if (!vehicle) return null;

            const coordinates = logs.map((log: any) => [log.lat, log.lon]);
            const color = getVehicleColor(vehicle.ID);
            const startTime = new Date(logs[0].created_at).toLocaleString();
            const endTime = new Date(
              logs[logs.length - 1].created_at
            ).toLocaleString();

            return (
              <div key={vehicleId}>
                <Polyline
                  positions={coordinates}
                  color={color}
                  weight={3}
                  opacity={0.7}
                />

                <CircleMarker
                  center={coordinates[0]}
                  radius={6}
                  fillColor={color}
                  color={color}
                  fillOpacity={1}
                >
                  <Popup>
                    <div className="space-y-1">
                      <p className="font-medium">Start Point</p>
                      <p className="text-sm">{startTime}</p>
                    </div>
                  </Popup>
                </CircleMarker>

                <CircleMarker
                  center={coordinates[coordinates.length - 1]}
                  radius={8}
                  fillColor={color}
                  color={color}
                  fillOpacity={1}
                >
                  <Popup className="vehicle-popup">
                    {/* Popup content remains same */}
                  </Popup>
                </CircleMarker>

                {coordinates.slice(1, -1).map((pos, idx) => (
                  <CircleMarker
                    key={idx}
                    center={pos}
                    radius={3}
                    fillColor={color}
                    color={color}
                    fillOpacity={0.8}
                  />
                ))}
              </div>
            );
          })}
        </MapContainer>
      </div>
    </div>
  );
}
