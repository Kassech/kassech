import { useState, useEffect, useMemo } from 'react';
import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  Polyline,
  CircleMarker,
} from 'react-leaflet';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Calendar } from '@/components/ui/calendar';
import { Vehicle } from '@/types/vehicle';
import { Path } from '@/types/path';
import api from '@/api/axiosInstance';
import { MultiSelect } from '@/components/ui/multi-select';

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

    // Sort logs chronologically for each vehicle
    Object.values(groups).forEach((logs) =>
      logs.sort(
        (a, b) =>
          new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
      )
    );

    return groups;
  }, [gpsLogs]);

  useEffect(() => {
    fetchGPSData();
  }, [filters]);

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

  return (
    <div className="m-2 flex h-screen">
      {/* Filters Sidebar */}
      <Card className="w-150 p-4 overflow-y-auto space-y-4">
        <h2 className="text-xl font-bold">Vehicle Tracking</h2>

        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium mb-1">Vehicle</label>
            <Select onValueChange={(v) => handleFilterChange('vehicleId', v)}>
              <SelectTrigger>
                <SelectValue placeholder="Select vehicle" />
              </SelectTrigger>
              <SelectContent>
                {vehicles.map((vehicle) => (
                  <SelectItem key={vehicle.ID} value={vehicle.ID.toString()}>
                    {vehicle.LicenseNumber} ({vehicle.Type.TypeName})
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
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

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium mb-1">Latitude</label>
              <Input
                type="number"
                value={filters.lat}
                onChange={(e) => handleFilterChange('lat', e.target.value)}
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">
                Longitude
              </label>
              <Input
                type="number"
                value={filters.lon}
                onChange={(e) => handleFilterChange('lon', e.target.value)}
              />
            </div>
          </div>

          <Button onClick={fetchGPSData} className="w-full">
            Apply Filters
          </Button>
        </div>
      </Card>

      {/* Map */}
      <div className="flex-1">
        <MapContainer
          center={[9.145, 40.4897]}
          zoom={12}
          className="h-full w-full"
          attributionControl={false}
        >
          <TileLayer url="https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png" />

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

                {/* Start Marker */}
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

                {/* End Marker */}
                <CircleMarker
                  center={coordinates[coordinates.length - 1]}
                  radius={8}
                  fillColor={color}
                  color={color}
                  fillOpacity={1}
                >
                  <Popup className="vehicle-popup">
                    <div className="flex gap-4">
                      <img
                        src={vehicle.Driver.User.ProfilePicture}
                        alt="Driver"
                        className="w-16 h-16 rounded-lg object-cover"
                      />
                      <div className="space-y-2">
                        <div>
                          <h3 className="font-bold">
                            {vehicle.Make} {vehicle.Year}
                          </h3>
                          <p className="text-sm text-muted-foreground">
                            {vehicle.LicenseNumber}
                          </p>
                        </div>
                        <div className="space-y-1">
                          <p className="text-sm">
                            <span className="font-medium">Driver:</span>{' '}
                            {vehicle.Driver.User.FirstName}
                          </p>
                          <p className="text-sm">
                            <span className="font-medium">Status:</span>{' '}
                            {vehicle.Status}
                          </p>
                        </div>
                        <div className="text-xs text-gray-500 space-y-1">
                          <p>🟢 Start: {startTime}</p>
                          <p>🔴 End: {endTime}</p>
                        </div>
                      </div>
                    </div>
                  </Popup>
                </CircleMarker>

                {/* Intermediate Points */}
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
