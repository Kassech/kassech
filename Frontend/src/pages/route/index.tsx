import React from 'react';
import MapContainerComponent from './component/MapContainerComponent';
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable';
import Header from '@/components/header';
import { CreateRouteForm } from './component/createForm';
import { useGetAllRoutes } from '@/services/routeService';
import LoadingSpinner from '@/components/loading-spinner';
import { useGetAllStations } from '@/services/stationService';
import { DataTable } from './table/data-table';
import { columns } from './table/column';

const paths = [
  { name: 'Home', href: '/' },
  { name: 'Routes', href: '/routes' },
];

const RoutePage: React.FC = () => {
  const { data: routes, isLoading, error } = useGetAllRoutes();
  const {
    data: stations,
    isLoading: isStationLoading,
    error: stationError,
  } = useGetAllStations();
  console.log('🚀 ~ stations:', stations);

  console.log('🚀 ~ routes:', routes);
  if (isLoading || isStationLoading) return <LoadingSpinner />;
  if (error || stationError) return <div>Error loading routes!</div>;

  return (
    <>
      <Header paths={paths} />
      <div className="p-6 w-full h-full    ">
        <ResizablePanelGroup
          direction="horizontal"
          className="w-full rounded-lg border "
        >
          <ResizablePanel defaultSize={40}>
            <ResizablePanelGroup direction="vertical">
              <ResizablePanel defaultSize={40}>
                <CreateRouteForm data={stations} />
              </ResizablePanel>
              <ResizableHandle />
              <ResizablePanel defaultSize={60}>
                <div className="-2">
                  <DataTable data={routes ?? []} columns={columns} />{' '}
                </div>
              </ResizablePanel>
            </ResizablePanelGroup>
          </ResizablePanel>
          <ResizableHandle />
          <ResizablePanel defaultSize={60}>
            <MapContainerComponent stations={stations ?? []} />
          </ResizablePanel>
          <ResizableHandle />
        </ResizablePanelGroup>
      </div>
    </>
  );
};

export default RoutePage;
