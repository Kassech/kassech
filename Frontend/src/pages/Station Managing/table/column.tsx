/* eslint-disable react-hooks/rules-of-hooks */
import { Station } from '@/types/api';
import { ColumnDef } from '@tanstack/react-table';
import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { MoreHorizontal } from 'lucide-react';
import { useDeleteStation } from '@/services/stationService';
import { useStationStore } from '@/store/stationStore';
import { DataTableColumnHeader } from './data-table-column-header';
import L from 'leaflet';

export const columns: ColumnDef<Station>[] = [
  {
    accessorKey: 'LocationName',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="LocationName" />
    ),
  },
  {
    accessorKey: 'Latitude',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Latitude" />
    ),
  },
  {
    accessorKey: 'Longitude',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Longitude" />
    ),
  },
  {
    id: 'actions',
    cell: ({ row }) => {
      const station = row.original;
      const deleteStationMutation = useDeleteStation();
      const { setPosition, setLocationName, setEditingStationId } =
        useStationStore();

      const handleEditStation = () => {
        setEditingStationId(station.ID);
        setPosition(new L.LatLng(station.Latitude, station.Longitude));
        setLocationName(station.LocationName);
      };

      const handleDeleteStation = () => {
        deleteStationMutation.mutate(station.ID);
      };

      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem onClick={handleEditStation}>
              Edit
            </DropdownMenuItem>
            <DropdownMenuItem onClick={handleDeleteStation}>
              Delete
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
    size: 80,
  },
];
