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
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { ColumnDef } from '@tanstack/react-table';
import { MoreHorizontal } from 'lucide-react';

export type Route = {
  id: string;
  StationA: {
    LocationName: string;
  };
  StationB: {
    LocationName: string;
  };
};

export const columns: ColumnDef<Route>[] = [
  {
    accessorKey: 'StationA',
    header: 'Location A',
    cell: ({ row }) => <div>{row.getValue('StationA')?.LocationName}</div>,
  },
  {
    accessorKey: 'StationB',
    header: 'Location B',
    cell: ({ row }) => <div>{row.getValue('StationB')?.LocationName}</div>,
  },

  {
    id: 'actions',
    enableHiding: false,
    cell: ({ row }) => {
      const route = row.original;
      console.log('🚀 ~ route:', route);

      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>
            <AlertDialog>
              <AlertDialogTrigger asChild>
                <Button variant="destructive">Delete</Button>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                  <AlertDialogDescription>
                    This action cannot be undone. This will permanently delete
                    your account and remove your data from our servers.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                  <AlertDialogAction>Continue</AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];

const handleDeleteRoute = (id: string) => {
  // logic to remove route
};
