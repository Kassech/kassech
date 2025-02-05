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
import {
  PlusIcon,
  TruckIcon,
  UsersIcon,
  MapPinIcon,
  ArchiveBoxIcon,
  KeyIcon,
  DocumentTextIcon,
} from '@heroicons/react/24/solid';

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

  const {
    data: totalUsers,
    isLoading: loadingtotalUsers,
    error: errortotalUsers,
  } = useTotalUsers();

  // Extract count if loginLogs is an array
  const loginLogsCount = Array.isArray(loginLogs) ? loginLogs.length : 0;

  const stats = [
    {
      title: 'Total Drivers',
      value: totalDrivers,
      loading: loadingDrivers,
      error: errorDrivers,
      icon: <UsersIcon className="w-8 h-8 text-customBlue" />,
    },
    {
      title: 'Total Vehicles',
      value: totalVehicles,
      loading: loadingVehicles,
      error: errorVehicles,
      icon: <TruckIcon className="w-8 h-8 text-customBlue" />,
    },
    {
      title: 'Total Users',
      value: totalUsers,
      loading: loadingtotalUsers,
      error: errortotalUsers,
      icon: <UsersIcon className="w-8 h-8 text-customBlue" />,
    },
    {
      title: 'Total Routes',
      value: totalRoutes,
      loading: loadingRoutes,
      error: errorRoutes,
      icon: <MapPinIcon className="w-8 h-8 text-customBlue" />,
    },
    {
      title: 'Total Stations',
      value: totalStations,
      loading: loadingStations,
      error: errorStations,
      icon: <ArchiveBoxIcon className="w-8 h-8 text-customBlue" />,
    },
    {
      title: 'Login Logs',
      value: loginLogsCount,
      loading: loadingLoginLogs,
      error: errorLoginLogs,
      icon: <KeyIcon className="w-8 h-8 text-customBlue" />,
    },
    {
      title: 'Total Travel Logs',
      value: totalTravelLogs,
      loading: loadingTravelLogs,
      error: errorTravelLogs,
      icon: <DocumentTextIcon className="w-8 h-8 text-customBlue" />,
    },
  ];

  return (
    <div className="space-y-6 p-4">
      <div className="grid auto-rows-min md:grid-cols-4 gap-4">
        {stats.map((item, index) => (
          <div
            key={index}
            className="rounded-xl bg-muted/10 p-4 shadow-md hover:shadow-lg transition-all"
          >
            <Card className="w-full h-full flex flex-col">
              <CardHeader>
                <div className="flex items-center space-x-2">
                  {item.icon}
                  <CardTitle className="text-lg font-semibold">
                    {item.title}
                  </CardTitle>
                </div>
              </CardHeader>
              <CardContent className="flex-1 flex justify-center items-center text-3xl text-customBlue font-bold">
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
