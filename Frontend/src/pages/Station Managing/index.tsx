import React, { useEffect } from "react";
import { MapContainer, TileLayer, Marker, useMapEvents } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import * as L from "leaflet";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  useDeleteStation,
  useUpdateStation,
  useCreateStation,
  useGetAllStations,
} from "../../services/stationService";
import { useStationStore } from "../../store/stationStore";

const SearchMap: React.FC = () => {
  const {
    position,
    LocationName,
    editingStationId,
    setPosition,
    setLocationName,
    setEditingStationId,
  } = useStationStore();

  const { data: stations, refetch } = useGetAllStations();
  const createStationMutation = useCreateStation();
  const updateStationMutation = useUpdateStation();
  const deleteStationMutation = useDeleteStation();

  const fetchLocationName = async (lat: number, lng: number) => {
    try {
      const response = await fetch(
        `https://nominatim.openstreetmap.org/reverse?format=json&lat=${lat}&lon=${lng}`
      );
      const data = await response.json();
      if (data && data.display_name) {
        setLocationName(data.display_name);
      } else {
        setLocationName("Unknown location");
      }
    } catch (error) {
      console.error("Error fetching location name:", error);
      setLocationName("Error fetching location name");
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
            setPosition(null);
            setLocationName("");
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
            setPosition(null);
            setLocationName("");
          },
        }
      );
    }
  };

  const handleDeleteStation = (id: number) => {
    deleteStationMutation.mutate(id, {
      onSuccess: () => {
        refetch();
      },
    });
  };

  return (
    <div>
      <MapContainer
        center={[51.505, -0.09]}
        zoom={13}
        style={{ height: "50vh", width: "100%" }}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />
        <MapClickHandler />
        {position && <Marker position={position} />}
      </MapContainer>

      <div className="p-4 bg-white border-t shadow-md">
        <h3 className="text-lg font-bold">
          {editingStationId ? "Edit Station" : "Add Station"}
        </h3>
        {position ? (
          <>
            <Label htmlFor="latitude">Latitude</Label>
            <Input
              id="latitude"
              value={position.lat.toFixed(4)}
              readOnly
              className="mb-2"
            />
            <Label htmlFor="longitude">Longitude</Label>
            <Input
              id="longitude"
              value={position.lng.toFixed(4)}
              readOnly
              className="mb-2"
            />
            <Label htmlFor="location">Location Name</Label>
            <Input
              id="location"
              value={LocationName}
              readOnly
              className="mb-4"
            />
            {editingStationId ? (
              <Button onClick={updateStationInfo}>Update Station</Button>
            ) : (
              <Button onClick={addStation}>Add Station</Button>
            )}
          </>
        ) : (
          <p>Click on the map to select a location.</p>
        )}
      </div>

      <div className="p-4 bg-gray-100">
        <h3 className="text-lg font-bold">Managed Stations</h3>
        {stations && stations.length > 0 ? (
          <ul className="mt-2">
            {stations.map((station) => (
              <li
                key={station.ID}
                className="flex justify-between items-center p-2 bg-white shadow mb-2 rounded"
              >
                <div>
                  {station ? (
                    <>
                      <p>
                        <strong>
                          {station.LocationName || "Unknown Location"}
                        </strong>
                      </p>
                      {station.Latitude != null && station.Longitude != null ? (
                        <p>
                          Lat: {(station.Latitude ?? 0).toFixed(4)}, Lng:{" "}
                          {(station.Longitude ?? 0).toFixed(4)}
                        </p>
                      ) : (
                        <p>Coordinates not available</p>
                      )}
                    </>
                  ) : (
                    <p>Station data not available</p>
                  )}
                </div>
                <div className="flex gap-2">
                  <Button
                    onClick={() => {
                      setEditingStationId(station.ID);
                      setPosition(
                        new L.LatLng(station.Latitude, station.Longitude)
                      );
                      setLocationName(station.LocationName);
                    }}
                    variant="secondary"
                    size="sm"
                  >
                    Edit
                  </Button>
                  <Button
                    onClick={() => handleDeleteStation(station.ID)}
                    variant="destructive"
                    size="sm"
                  >
                    Remove
                  </Button>
                </div>
              </li>
            ))}
          </ul>
        ) : (
          <p>No stations added yet.</p>
        )}
      </div>
    </div>
  );
};

export default SearchMap;
