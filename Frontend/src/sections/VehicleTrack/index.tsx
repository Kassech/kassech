import React, { useEffect, useState } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import { Card, CardContent } from '@/components/ui/card';
import { Input } from '@/components/ui/input';

const dummyDrivers = [
  {
    id: 1,
    name: 'John Doe',
    email: 'johndoe@example.com',
    phone: '+1234567890',
    status: 'Active',
    latitude: 9.03,
    longitude: 38.74,
    vehicleImage:
      'https://l450v.alamy.com/450v/d4akyn/typical-ethiopian-minibus-addis-ababa-africa-d4akyn.jpg',
    driverImage:
      'https://plus.unsplash.com/premium_photo-1689708721750-8a0e6dc14cee?q=80&w=1470&auto=format&fit=crop&ixlib=rb-4.0.3',
  },
  {
    id: 2,
    name: 'Jane Smith',
    email: 'janesmith@example.com',
    phone: '+0987654321',
    status: 'Idle',
    latitude: 9.05,
    longitude: 38.76,
    vehicleImage:
      'https://l450v.alamy.com/450v/d4akyn/typical-ethiopian-minibus-addis-ababa-africa-d4akyn.jpg',
    driverImage:
      'https://plus.unsplash.com/premium_photo-1689708721750-8a0e6dc14cee?q=80&w=1470&auto=format&fit=crop&ixlib=rb-4.0.3',
  },
];

export function VehicleTrack() {
  const [driverLocations, setDriverLocations] = useState(dummyDrivers);
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedDriver, setSelectedDriver] = useState<(typeof dummyDrivers)[number] | null>(null);



  useEffect(() => {
    const interval = setInterval(() => {
      setDriverLocations((prevDrivers) =>
        prevDrivers.map((driver) => ({
          ...driver,
          latitude: driver.latitude + (Math.random() - 0.5) * 0.001,
          longitude: driver.longitude + (Math.random() - 0.5) * 0.001,
        }))
      );
    }, 5000);
    return () => clearInterval(interval);
  }, []);

  const filteredDrivers = driverLocations.filter((driver) =>
    driver.name.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className="relative h-screen">
      <div className="absolute top-4 left-4 z-[1000] bg-white p-4 rounded-lg shadow-lg max-w-sm max-h-[60vh] overflow-y-auto w-64 space-y-3">
        <Input
          placeholder="Search driver..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="mb-3"
        />
        {filteredDrivers.map((driver) => (
          <Card
            key={driver.id}
            onClick={() => setSelectedDriver(driver)}
            className="cursor-pointer"
          >
            <CardContent className="flex items-center space-x-4 p-1">
              <img
                src={driver.driverImage}
                alt="Driver"
                className="w-12 h-12 rounded-full"
              />
              <div>
                <h3 className="text-md font-semibold">{driver.name}</h3>
                <p className="text-sm">{driver.status}</p>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
      <MapContainer
        center={[9.03, 38.74]}
        zoom={13}
        style={{ height: '100vh', width: '100%' }}
      >
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
        {filteredDrivers.map((driver) => (
          <Marker
            key={driver.id}
            position={[driver.latitude, driver.longitude]}
            icon={L.icon({
              iconUrl: driver.vehicleImage,
              iconSize: [40, 40],
              iconAnchor: [20, 20],
              popupAnchor: [0, -20],
            })}
          >
            <Popup>
              <div className="text-center">
                <img
                  src={driver.driverImage}
                  alt={driver.name}
                  className="w-20 h-20 rounded-full mx-auto"
                />
                <h3 className="text-lg font-semibold mt-2">{driver.name}</h3>
                <p>Email: {driver.email}</p>
                <p>Phone: {driver.phone}</p>
                <p>Status: {driver.status}</p>
              </div>
            </Popup>
          </Marker>
        ))}
      </MapContainer>
    </div>
  );
}
