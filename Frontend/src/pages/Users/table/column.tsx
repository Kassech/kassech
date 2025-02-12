/* eslint-disable react-hooks/rules-of-hooks */
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
import { User } from '@/types/user';
import { DataTableColumnHeader } from './data-table-column-header'; // Import your header component
import { userManagingStore } from '@/store/usersManagingStore';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { useDeleteUser } from '@/services/userService';

import QueueManagerForm from '@/sections/QueueManager/queueManagerForm';
import VerificationToggle from './verificationToggle';
import { toast } from 'sonner';
import DriverPage from '@/pages/driver';
import AdminForm from '@/sections/Admin';
import CarOwnerForm from '@/sections/owner/OwnerForm';

export const columns: ColumnDef<User>[] = [
  {
    accessorKey: 'first_name',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="First Name" />
    ),
  },
  {
    accessorKey: 'last_name',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Last Name" />
    ),
  },
  {
    accessorKey: 'email',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Email" />
    ),
  },
  {
    accessorKey: 'phone_number',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Phone Number" />
    ),
  },
  {
    accessorKey: 'is_online',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Online Status" />
    ),
    cell: ({ row }) => (row.original.IsOnline ? 'Online' : 'Offline'),
  },
  {
    accessorKey: 'is_verified',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Verified" />
    ),
    cell: ({ row }) => {
      console.log('is verified: ', row.original.is_verified);
      return (
        <VerificationToggle
          initialVerified={row.original.is_verified}
          userId={row.original.id}
        />
      );
    },
  },

  {
    id: 'actions',
    cell: ({ row }) => {
      console.log('🚀 ~ row:', row);
      const {
        isDialogOpen,
        selectedUser,
        selectedUserRole,
        setDialogOpen,
        setDialogClose,
        isEditDialogOpen,
        setEditDialogOpen,
        setEditDialogClose,
      } = userManagingStore();

      const { mutate: deleteUser } = useDeleteUser();
      const handleDelete = (userid: number) => {
        console.log(userid);
        console.log('🚀 ~ handleDelete ~ userid:', userid);
        deleteUser(userid, {
          onSuccess: () => {
            toast('User Deletion', {
              description: 'User deleted successfully.',
            });
            console.log(`User with id ${userid} deleted successfully.`);
            setDialogClose();
          },
          onError: (error) => {
            console.error(`Failed to delete user: ${error}`);
          },
        });
      };
      console.log('🚀 ~ row.original:', row.original);

      return (
        <>
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
              <DropdownMenuItem
                onClick={() =>
                  setEditDialogOpen(row.original.id, row.original.roles)
                }
              >
                Edit
              </DropdownMenuItem>
              <DropdownMenuItem onClick={() => setDialogOpen(row.original.id)}>
                Delete
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
          {console.log(selectedUser, row.original.id)}
          {console.log('🚀 ~ selectedUser:', selectedUser)}{' '}
          {console.log('🚀 ~ selectedUserRole:', selectedUserRole)}{' '}
          {console.log('🚀 ~ row.original.id:', row.original.id)}{' '}
          {selectedUser === row.original.id && (
            <AlertDialog open={isDialogOpen} onOpenChange={setDialogClose}>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>Confirm Delete</AlertDialogTitle>
                  <AlertDialogDescription>
                    Are you sure you want to delete this user? This action
                    cannot be undone.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel onClick={setDialogClose}>
                    Cancel
                  </AlertDialogCancel>

                  <AlertDialogAction
                    onClick={() => handleDelete(row.original.id)}
                  >
                    Delete
                  </AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          )}
          {selectedUser === row.original.id && (
            <AlertDialog
              open={isEditDialogOpen}
              onOpenChange={setEditDialogClose}
            >
              <AlertDialogContent className="overflow-y-auto max-h-[80vh]">
                <AlertDialogHeader>
                  <AlertDialogTitle>Edit User</AlertDialogTitle>
                  <AlertDialogDescription>
                    Update the details of the user below.
                    {Array.isArray(selectedUserRole) &&
                    selectedUserRole.length > 0 ? (
                      selectedUserRole[0] === 'Admin' ? (
                        <AdminForm defaultValues={row.original} />
                      ) : selectedUserRole[0] === 'Driver' ? (
                        <DriverPage defaultValues={row.original} />
                      ) : selectedUserRole[0] === 'QueueManager' ? (
                        <QueueManagerForm defaultValues={row.original} />
                      ) : selectedUserRole[0] === 'Owner' ? (
                        <CarOwnerForm defaultValues={row.original} />
                      ) : (
                        <p className="text-red-500">
                          Invalid role: {selectedUserRole[0]}
                        </p>
                      )
                    ) : (
                      <p className="text-red-500">
                        No role assigned or invalid role data.
                      </p>
                    )}
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          )}
        </>
      );
    },
    size: 80, // Define column width for actions
  },
];
