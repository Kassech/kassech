import React from 'react';
import MapContainerComponent from './component/MapContainerComponent';
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable';
import Header from '@/components/header';
import { CreatePathForm } from './component/createForm';
import { useGetAllRoutes } from '@/services/routeService';
import LoadingSpinner from '@/components/loading-spinner';
import { DataTable } from './table/data-table';
import { columns } from './table/column';
import { useGetAllPaths } from '@/services/pathService';

const BreadcrumbPath = [
  { name: 'Home', href: '/' },
  { name: 'Paths', href: '/paths' },
];
const PathPage: React.FC = () => {
  const {
    data: routes,
    isLoading: routesLoading,
    error: routesError,
  } = useGetAllRoutes();
  const {
    data: paths,
    isLoading: stationsLoading,
    error: stationsError,
  } = useGetAllPaths();

  if (routesLoading || stationsLoading) return <LoadingSpinner />;
  if (routesError || stationsError) return <div>Error loading routes!</div>;

  return (
    <>
      <Header paths={BreadcrumbPath} />
      <div className="p-6 w-full h-full">
        <ResizablePanelGroup
          direction="horizontal"
          className="w-full rounded-lg border"
        >
          <ResizablePanel defaultSize={50}>
            <ResizablePanelGroup direction="vertical">
              <ResizablePanel defaultSize={60}>
                <CreatePathForm routes={routes} />
              </ResizablePanel>
              <ResizablePanel defaultSize={40}>
                <DataTable data={paths?.data ?? []} columns={columns} />
              </ResizablePanel>
            </ResizablePanelGroup>
          </ResizablePanel>
          <ResizableHandle />
          <ResizablePanel defaultSize={50}>
            <MapContainerComponent />
          </ResizablePanel>
          <ResizableHandle />
        </ResizablePanelGroup>
      </div>
    </>
  );
};

export default PathPage;
