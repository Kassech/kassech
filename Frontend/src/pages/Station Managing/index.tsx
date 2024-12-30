import React, { useState, useEffect } from "react";
import { MapContainer, TileLayer, Marker, useMapEvents } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import * as L from "leaflet";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  fetchStations,
  createStation,
  updateStation,
  deleteStation,
} from "../../services/stationService";

interface Station {
  ID: number;
  LocationName: string;
  Latitude: number;
  Longitude: number;
}

const SearchMap: React.FC = () => {
  const [position, setPosition] = useState<L.LatLng | null>(null);
  const [locationName, setLocationName] = useState<string>("");
  const [stations, setStations] = useState<Station[]>([]);
  const [editingStationId, setEditingStationId] = useState<number | null>(null);

  const loadStations = async () => {
    try {
      const data = await fetchStations();
      console.log("Fetched stations:", data); // Debug log
      setStations(data); // Update the state with fetched stations
    } catch (error) {
      console.error("Failed to fetch stations:", error);
    }
  };

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

  const addStation = async () => {
    if (position && locationName) {
      const newStation = {
        LocationName: locationName,
        Latitude: position.lat,
        Longitude: position.lng,
      };
      try {
        const createdStation = await createStation(newStation);
        console.log("Created station:", createdStation); // Debug log
        setStations((prev) => [...prev, createdStation]); // Add to the local state
        setPosition(null); // Clear selection
        setLocationName("");
      } catch (error) {
        console.error("Failed to add station:", error);
      }
    }
  };

  const updateStationInfo = async () => {
    if (position && locationName && editingStationId !== null) {
      const updatedData = {
        LocationName: locationName,
        Latitude: position.lat,
        Longitude: position.lng,
      };
      try {
        const updatedStation = await updateStation(
          editingStationId,
          updatedData
        );
        console.log("Updated station:", updatedStation); // Debug log
        setStations((prev) =>
          prev.map((station) =>
            station.ID === editingStationId ? updatedStation : station
          )
        );
        setEditingStationId(null);
        setPosition(null);
        setLocationName("");
      } catch (error) {
        console.error("Failed to update station:", error);
      }
    }
  };

  const handleDeleteStation = async (id: number) => {
    try {
      console.log(`Deleting station with ID: ${id}`); // Debug log
      await deleteStation(id);
      setStations((prev) => prev.filter((station) => station.ID !== id)); // Remove from local state
    } catch (error) {
      console.error("Failed to delete station:", error);
    }
  };

  useEffect(() => {
    loadStations();
  }, []);

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
              value={locationName}
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
        {stations.length > 0 ? (
          <ul className="mt-2">
            {stations.map((station) => (
              <li
                key={station.ID}
                className="flex justify-between items-center p-2 bg-white shadow mb-2 rounded"
              >
                <div>
                  <p>
                    <strong>{station.LocationName}</strong>
                  </p>
                  <p>
                    Lat: {station.Latitude.toFixed(4)}, Lng:{" "}
                    {station.Longitude.toFixed(4)}
                  </p>
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
