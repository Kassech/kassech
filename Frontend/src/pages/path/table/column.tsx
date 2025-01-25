import { Badge } from '@/components/ui/badge';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog';
import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { useDeletePath } from '@/services/pathService';
import { ColumnDef } from '@tanstack/react-table';
import { MoreHorizontal } from 'lucide-react';
import { Path } from '@/types/path';

export const columns: ColumnDef<Path>[] = [
  {
    accessorKey: 'path_name',
    header: 'Path Name',
    cell: ({ row }) => (
      <div className="font-medium">{row.getValue('path_name')}</div>
    ),
  },
  {
    accessorKey: 'route.StationA.LocationName',
    header: 'Start Station',
    cell: ({ row }) => (
      <div>{row.original.route?.StationA?.LocationName || 'N/A'}</div>
    ),
  },
  {
    accessorKey: 'route.StationB.LocationName',
    header: 'End Station',
    cell: ({ row }) => {
      console.log('🚀 ~ row.original.route:', row.original.route);
      return <div>{row.original.route?.StationB?.LocationName || 'N/A'}</div>;
    },
  },
  {
    accessorKey: 'distance_km',
    header: 'Distance',
    cell: ({ row }) => (
      <div className="text-right">
        {Number(row.getValue('distance_km')).toFixed(2)} km
      </div>
    ),
  },
  {
    accessorKey: 'estimated_time',
    header: 'Estimated Time',
    cell: ({ row }) => {
      const time = row.getValue('estimated_time') as string;
      const formattedTime = time.replace(/[^\dhm]/g, '');
      return <div>{formattedTime}</div>;
    },
  },
  {
    accessorKey: 'is_active',
    header: 'Status',
    cell: ({ row }) => (
      <Badge variant={row.getValue('is_active') ? 'default' : 'secondary'}>
        {row.getValue('is_active') ? 'Active' : 'Inactive'}
      </Badge>
    ),
  },
  {
    accessorKey: 'CreatedAt',
    header: 'Created At',
    cell: ({ row }) => (
      <div>{new Date(row.getValue('CreatedAt')).toLocaleDateString()}</div>
    ),
  },
  {
    id: 'actions',
    enableHiding: false,
    cell: ({ row }) => {
      const path = row.original;
      const deleteMutation = useDeletePath();

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
            <AlertDialog>
              <AlertDialogTrigger asChild>
                <Button variant="destructive" className="w-full">
                  Delete
                </Button>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>Confirm Deletion</AlertDialogTitle>
                  <AlertDialogDescription>
                    This will permanently delete the path between{' '}
                    {path.route?.StationA?.LocationName} and{' '}
                    {path.route?.StationB?.LocationName}. This action cannot be
                    undone.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                  <AlertDialogAction
                    onClick={() => deleteMutation.mutate(path.ID)}
                  >
                    Confirm Delete
                  </AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];
