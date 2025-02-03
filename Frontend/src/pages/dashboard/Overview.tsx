'use client';

import {
  useGetTotalUsers,
  useGetActiveUsers,
  useGetTotalDrivers,
  useGetTotalVehicles,
  useGetActiveVehicles,
  useGetTotalTravelLogs,
  useGetTotalRoutes,
  useGetTotalStations,
  useGetLoginLogs,
} from '@/services/dashboardService';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'; // Importing ShadCN card components
import ProgresBar from './ProgresBar';


export default function Overview() {
  const totalDrivers = useGetTotalDrivers()?.data;
  const totalVehicles = useGetTotalVehicles()?.data;
  const totalTravelLogs = useGetTotalTravelLogs()?.data;
  const totalRoutes = useGetTotalRoutes()?.data;
  const totalStations = useGetTotalStations()?.data;
  const loginLogs = useGetLoginLogs()?.data;



  return (
    <div className="space-y-6 p-4">
      <div className="grid auto-rows-min md:grid-cols-3 gap-4">
        {[
          // { title: 'Active Users', value: activeUsers },
          { title: 'Total Drivers', value: totalDrivers },
          { title: 'Total Vehicles', value: totalVehicles },
          // { title: 'Active Vehicles', value: activeVehicles },
          // { title: 'Total Travel Logs', value: totalTravelLogs },
          { title: 'Total Routes', value: totalRoutes },
          { title: 'Total Stations', value: totalStations },
          // { title: 'Login Logs', value: loginLogs },
        ].map((item, index) => (
          <div className="aspect-video rounded-xl bg-muted/50 flex items-center justify-center">
            <Card key={index} className="w-full h-full flex flex-col">
              <CardHeader>
                <CardTitle>{item.title}</CardTitle>
              </CardHeader>
              <CardContent className="flex-1 flex items-center justify-center">
                {typeof item.value === 'number' ? item.value : 'Loading...'}
              </CardContent>
            </Card>
          </div>
        ))}
      </div>
      <div className=" ">
        <ProgresBar />
      </div>
    </div>
  );
}
