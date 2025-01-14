import React from 'react';
import { MapContainer, TileLayer, Marker, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import * as L from 'leaflet';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import {
  useUpdateStation,
  useCreateStation,
  useGetAllStations,
} from '../../services/stationService';
import { useStationStore } from '../../store/stationStore';
import Header from '@/components/header';
import { DataTable } from './table/data-table';
import { columns } from './table/column';
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable';

const StationsPage: React.FC = () => {
  const {
    position,
    LocationName,
    editingStationId,
    setPosition,
    setLocationName,
    setEditingStationId,
    clearLocations,
  } = useStationStore();

  const { data: stations, refetch } = useGetAllStations();
  const createStationMutation = useCreateStation();
  const updateStationMutation = useUpdateStation();

  const fetchLocationName = async (lat: number, lng: number) => {
    try {
      const response = await fetch(
        `https://nominatim.openstreetmap.org/reverse?format=json&lat=${lat}&lon=${lng}`
      );
      const data = await response.json();
      if (data && data.display_name) {
        setLocationName(data.display_name);
      } else {
        setLocationName('Unknown location');
      }
    } catch (error) {
      console.error('Error fetching location name:', error);
      setLocationName('Error fetching location name');
    }
  };

  const MapClickHandler = () => {
    useMapEvents({
      click(event: L.LeafletMouseEvent) {
        const { lat, lng } = event.latlng;
        setPosition(new L.LatLng(lat, lng));
        fetchLocationName(lat, lng);
      },
    });
    return null;
  };

  const addStation = () => {
    if (position && LocationName) {
      createStationMutation.mutate(
        {
          LocationName,
          Latitude: position.lat,
          Longitude: position.lng,
        },
        {
          onSuccess: () => {
            refetch();
            setPosition(null); // Clear position
            setLocationName(''); // Clear location name
            setEditingStationId(null); // Reset editing ID (if needed)
          },
        }
      );
    }
  };

  const updateStationInfo = () => {
    if (position && LocationName && editingStationId !== null) {
      updateStationMutation.mutate(
        {
          id: editingStationId,
          updatedStation: {
            LocationName,
            Latitude: position.lat,
            Longitude: position.lng,
          },
        },
        {
          onSuccess: () => {
            refetch();
            setEditingStationId(null);
            clearLocations();
          },
        }
      );
    }
  };

  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Stations', href: '/stations' },
  ];

  return (
    <div className="flex flex-col h-screen">
      <Header paths={paths} />
      <ResizablePanelGroup
        direction="horizontal"
        className="h-screen rounded-lg border "
      >
        <ResizablePanel defaultSize={50}>
          <MapContainer
            center={[51.505, -0.09]}
            zoom={13}
            style={{ zIndex: 49, height: '100vh', width: '100%' }}
          >
            <TileLayer
              url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            />
            <MapClickHandler />
            {position && <Marker position={position} />}
          </MapContainer>
        </ResizablePanel>
        <ResizableHandle />
        <ResizablePanel defaultSize={50}>
          <ResizablePanelGroup direction="vertical">
            <ResizablePanel defaultSize={30}>
              <div className="p-4 bg-white h-full border-t shadow-md">
                <h3 className="text-lg font-bold">
                  {editingStationId ? 'Edit Station' : 'Add Station'}
                </h3>
                <>
                  <Label htmlFor="latitude">Latitude</Label>
                  <Input
                    id="latitude"
                    value={position?.lat.toFixed(4)}
                    readOnly
                    className="mb-2"
                  />
                  <Label htmlFor="longitude">Longitude</Label>
                  <Input
                    id="longitude"
                    value={position?.lng.toFixed(4)}
                    readOnly
                    className="mb-2"
                  />
                  <Label htmlFor="location">Location Name</Label>
                  <Input
                    id="location"
                    value={LocationName}
                    onChange={(e) => setLocationName(e.target.value)}
                    className="mb-4"
                  />
                  {editingStationId ? (
                    <Button onClick={updateStationInfo}>Update Station</Button>
                  ) : (
                    <Button onClick={addStation}>Add Station</Button>
                  )}
                </>
              </div>
            </ResizablePanel>
            <ResizableHandle />
            <ResizablePanel defaultSize={50}>
              <DataTable columns={columns} data={stations ?? []} />
            </ResizablePanel>
          </ResizablePanelGroup>
        </ResizablePanel>
      </ResizablePanelGroup>
    </div>
  );
};

export default StationsPage;
