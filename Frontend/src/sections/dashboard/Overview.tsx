'use client';

import {
  useTotalUsers,
  useTotalDrivers,
  useTotalVehicles,
  useTotalTravelLogs,
  useTotalRoutes,
  useTotalStations,
  useLoginLogs,
} from '@/services/dashboardService';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import ProgresBar from './ProgresBar';

export default function Overview() {
  const {
    data: totalDrivers,
    isLoading: loadingDrivers,
    error: errorDrivers,
  } = useTotalDrivers();
  const {
    data: totalVehicles,
    isLoading: loadingVehicles,
    error: errorVehicles,
  } = useTotalVehicles();
  const {
    data: totalTravelLogs,
    isLoading: loadingTravelLogs,
    error: errorTravelLogs,
  } = useTotalTravelLogs();
  const {
    data: totalRoutes,
    isLoading: loadingRoutes,
    error: errorRoutes,
  } = useTotalRoutes();
  const {
    data: totalStations,
    isLoading: loadingStations,
    error: errorStations,
  } = useTotalStations();
  const {
    data: loginLogs,
    isLoading: loadingLoginLogs,
    error: errorLoginLogs,
  } = useLoginLogs();

  // Extract count if loginLogs is an array
  const loginLogsCount = Array.isArray(loginLogs) ? loginLogs.length : 0;
console.log('login log',loginLogs)
  const stats = [
    {
      title: '🧑‍✈️ Total Drivers',
      value: totalDrivers,
      loading: loadingDrivers,
      error: errorDrivers,
    },
    {
      title: '🚗 Total Vehicles',
      value: totalVehicles,
      loading: loadingVehicles,
      error: errorVehicles,
    },
    {
      title: '📜 Total Travel Logs',
      value: totalTravelLogs,
      loading: loadingTravelLogs,
      error: errorTravelLogs,
    },
    {
      title: '🗺️ Total Routes',
      value: totalRoutes,
      loading: loadingRoutes,
      error: errorRoutes,
    },
    {
      title: '🏢 Total Stations',
      value: totalStations,
      loading: loadingStations,
      error: errorStations,
    },
    {
      title: '🔑 Login Logs',
      value: loginLogsCount,
      loading: loadingLoginLogs,
      error: errorLoginLogs,
    },
  ];

  return (
    <div className="space-y-6 p-4">
      <div className="grid auto-rows-min md:grid-cols-3 gap-4">
        {stats.map((item, index) => (
          <div
            key={index}
            className="aspect-video auto-rows-min rounded-xl bg-muted/50 flex items-center justify-center"
          >
            <Card className="w-full h-full flex flex-col">
              <CardHeader>
                <CardTitle>{item.title}</CardTitle>
              </CardHeader>
              <CardContent className="flex-1 flex items-center justify-center text-5xl text-[#3A7CA5] font-bold">
                {item.loading ? (
                  <span className="text-gray-500">Loading...</span>
                ) : item.error ? (
                  <span className="text-red-500">Error!</span>
                ) : (
                  item.value ?? 0
                )}
              </CardContent>
            </Card>
          </div>
        ))}
      </div>
      <div>
        <ProgresBar />
      </div>
    </div>
  );
}
