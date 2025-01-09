import {
  useGetAllRole,
  useDeleteRole,
  useCreateRole,
  useGetRoleById,
  useUpdateRole,
} from '@/services/roleServices';
import { useCreateRolePermission } from '@/services/role_permissionServices';
import { useGetAllPermission } from '@/services/permissionServices';
import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
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
import { useToast } from '@/hooks/use-toast';
import {
  useDialogStore,
  useRoleStore,
  useCardStore,
} from '@/store/useRoleStore';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { ChevronDownCircle, ChevronRight } from 'lucide-react';
import { useEffect, useState } from 'react';
import Header from '@/components/header';

export default function RolePermission() {
  const { toast } = useToast();
  const { toggleCard, showCard } = useCardStore();
  const {
    isDialogOpen,
    setDialogOpen,
    setDialogClose,
    isEditRoleDialogOpen,
    setEditRoleDialogOpen,
    setEditRoleDialogClose,
  } = useDialogStore();

  const { data: allRoles, isLoading: isRolesLoading } = useGetAllRole();
  const { data: allPermissions, isLoading: isPermissionsLoading } =
    useGetAllPermission();
  const { mutate: deleteRole } = useDeleteRole();
  const { mutate: editRole } = useUpdateRole();
  const { mutate: addRole } = useCreateRole();
  const { mutate: addPermissionToRole } = useCreateRolePermission();

  const {
    currentRole,
    selectedRole,
    newRole,
    setRolePermissions,
    rolePermissions,
    updateNewRoleField,
    setCurrentRole,
    setSelectedRole,
    updateCurrentRoleField,
  } = useRoleStore();

  const {
    data: selectedRoleData,
    isLoading,
    error,
  } = useGetRoleById(selectedRole);

  const handleAddRole = () => {
    toggleCard();
    if (newRole.RoleName && newRole.Description) {
      addRole(
        {
          RoleName: newRole.RoleName,
          Description: newRole.Description,
        },
        {
          onSuccess: (data) => {
            console.log('API response data:', data);
            toast({
              description: 'Role Added Successfully.',
            });
          },
          onError: (error) => {
            console.error(`Failed to add role: ${error}`);
          },
        }
      );
    } else {
      console.error('Role ID is undefined.');
    }
  };

  const confirmDelete = () => {
    if (selectedRole !== null) {
      deleteRole(selectedRole, {
        onSuccess: () => {
          toast({
            description: 'Role Deleted Successfully.',
          });
          console.log(`Role with ID ${selectedRole} deleted successfully.`);
          setDialogClose();
        },
        onError: (error) => {
          console.error(`Failed to delete role: ${error}`);
        },
      });
    }
  };

  const handleSaveChange = () => {
    setEditRoleDialogClose();
    if (currentRole?.ID !== undefined) {
      editRole(
        {
          id: currentRole.ID,
          updateRole: {
            RoleName: currentRole.RoleName,
            Description: currentRole.Description,
          },
        },
        {
          onSuccess: (data) => {
            console.log('API response data:', data);
            toast({
              description: 'Role Edited Successfully.',
            });
          },
          onError: (error) => {
            console.error(`Failed to edit role: ${error}`);
          },
        }
      );
    } else {
      console.error('Role ID is undefined.');
    }
  };

  useEffect(() => {
    if (isLoading) {
      console.log('Loading role data...');
    }

    if (error) {
      console.error('Error fetching role data:', error);
    }

    if (selectedRoleData) {
      setRolePermissions(selectedRoleData.Permissions || []);
    }
  }, [selectedRoleData, isLoading, error]);

  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Dashboard', href: '/b' },
  ];

  interface Permission {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    PermissionName: string;
    Description: string;
  }
  const [highlightedPermissions, setHighlightedPermissions] = useState<
    number[]
  >([]);

  const handlePermissionClick = async (permission: Permission) => {
    const isPermissionAssigned = rolePermissions.some(
      (rp) => rp.ID === permission.ID
    );

    // Check if the selected role is valid (not null)
    if (selectedRole === null) {
      console.error('No role selected!');
      return; // Exit early if no role is selected
    }

    const roleId = selectedRole; // Ensure roleId is a valid number

    if (isPermissionAssigned) {
      // Remove permission from role
      // setRolePermissions(rolePermissions.filter((rp) => rp.ID !== permission.ID));
      // Call API to remove the permission for the role
      // useDeleteRolePermission({ roleId, permissionId: permission.ID });
      console.log('already assigned');
    } else {
      addPermissionToRole(
        {
          RoleID: roleId,
          PermissionID: permission.ID,
        },
        {
          onSuccess: (data) => {
            console.log('API response data:', data);
            setHighlightedPermissions((prev) => [...prev, permission.ID]);
          },
          onError: (error) => {
            console.error(`Failed to add role: ${error}`);
          },
        }
      );
    }
  };

  return (
    <>
      <Header paths={paths} />
      <div className="h-screen flex flex-col">
        <div className=" border-b-2 p-6">
          <h1 className="text-xl font-semibold mb-4">Roles</h1>
          <div className="flex flex-wrap gap-4">
            {isRolesLoading ? (
              <p>Loading roles...</p>
            ) : (
              allRoles &&
              allRoles.map((role) => (
                <div
                  key={role.ID ?? role.RoleName}
                  className="flex items-center gap-2"
                >
                  <Button
                    key={role.ID}
                    variant="outline"
                    // className="relative"
                    onClick={() => {
                      const prev = selectedRole;
                      setSelectedRole(role.ID);
                      if (prev != role.ID) {
                        setHighlightedPermissions([]);
                      }
                    }} // Update state with the role's id
                    className={`transition-colors duration-200 ease-in-out ${
                      selectedRole === role.ID
                        ? 'bg-blue-500 text-white border-blue-700'
                        : 'bg-white text-black'
                    }`}
                  >
                    {role.RoleName}
                  </Button>
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="icon">
                        <ChevronDownCircle />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                      <DropdownMenuItem
                        onClick={() => {
                          setCurrentRole(role);
                          setEditRoleDialogOpen();
                        }}
                      >
                        Edit
                      </DropdownMenuItem>
                      <DropdownMenuItem
                        onClick={() => {
                          setSelectedRole(role.ID);
                          setDialogOpen();
                        }}
                      >
                        Delete
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </div>
              ))
            )}
            <div className=" right-0">
              <AlertDialog>
                <AlertDialogTrigger>
                  <Button onClick={toggleCard}>Add New Role</Button>
                </AlertDialogTrigger>
                {showCard && (
                  <AlertDialogContent>
                    <AlertDialogHeader>
                      <AlertDialogTitle>Create New Role</AlertDialogTitle>
                    </AlertDialogHeader>
                    <div className="grid gap-4 py-4">
                      <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="name" className="text-right">
                          Role Name
                        </Label>
                        <Input
                          id="name"
                          defaultValue="rolename"
                          className="col-span-3"
                          value={newRole.RoleName}
                          onChange={(e) =>
                            updateNewRoleField('RoleName', e.target.value)
                          }
                        />
                      </div>
                      <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="username" className="text-right">
                          Description
                        </Label>
                        <Input
                          id="username"
                          defaultValue="description"
                          className="col-span-3"
                          value={newRole.Description}
                          onChange={(e) =>
                            updateNewRoleField('Description', e.target.value)
                          }
                        />
                      </div>
                    </div>
                    <AlertDialogFooter>
                      <AlertDialogCancel
                        onClick={() => {
                          console.log(
                            'Cancel button clicked, closing dialog...'
                          );
                          toggleCard();
                          console.log(isDialogOpen);
                        }}
                      >
                        Cancel
                      </AlertDialogCancel>
                      <AlertDialogAction onClick={handleAddRole}>
                        Continue
                      </AlertDialogAction>
                    </AlertDialogFooter>
                  </AlertDialogContent>
                )}
              </AlertDialog>
            </div>
          </div>
          {/* Form for role addition */}

          {/* Alert Dialog for Deletion */}
          {isDialogOpen && selectedRole !== null && (
            <AlertDialog open={isDialogOpen} onOpenChange={(open) => open}>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>Confirm Delete</AlertDialogTitle>
                  <AlertDialogDescription>
                    Are you sure you want to delete this user? This action
                    cannot be undone.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel
                    onClick={() => {
                      console.log('Cancel button clicked, closing dialog...');
                      setDialogClose();
                      console.log(isDialogOpen);
                    }}
                  >
                    Cancel
                  </AlertDialogCancel>
                  <AlertDialogAction onClick={confirmDelete}>
                    Delete
                  </AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          )}

          {/* Alert Dialog for Edit */}
          {isEditRoleDialogOpen && (
            <AlertDialog
              open={isEditRoleDialogOpen}
              onOpenChange={(open) => {
                if (open) {
                  setEditRoleDialogOpen();
                } else {
                  setEditRoleDialogClose();
                }
              }}
            >
              <AlertDialogContent className="sm:max-w-[425px]">
                <AlertDialogHeader>
                  <AlertDialogTitle>Edit Role</AlertDialogTitle>
                  <AlertDialogDescription>
                    Make changes to your profile here. Click save when you're
                    done.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <div className="grid gap-4 py-4">
                  <div className="grid grid-cols-4 items-center gap-4">
                    <Label htmlFor="name" className="text-right">
                      Role Name
                    </Label>
                    <Input
                      id="name"
                      defaultValue={currentRole.RoleName}
                      value={currentRole.RoleName}
                      className="col-span-3"
                      onChange={(e) =>
                        updateCurrentRoleField('RoleName', e.target.value)
                      }
                    />
                  </div>
                  <div className="grid grid-cols-4 items-center gap-4">
                    <Label htmlFor="username" className="text-right">
                      Description
                    </Label>
                    <Input
                      id="username"
                      defaultValue={currentRole.Description}
                      value={currentRole.Description}
                      className="col-span-3"
                      onChange={(e) =>
                        updateCurrentRoleField('Description', e.target.value)
                      }
                    />
                  </div>
                </div>
                <AlertDialogFooter>
                  <Button type="submit" onClick={handleSaveChange}>
                    Save changes
                  </Button>
                  <Button
                    onClick={() => {
                      setEditRoleDialogClose();
                    }}
                  >
                    Cancel
                  </Button>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          )}
        </div>
        {/* Permissions Section */}
        <div className=" p-6">
          <h1 className="text-xl font-semibold mb-4">Permissions</h1>

          {isPermissionsLoading ? (
            <p>Loading permissions...</p>
          ) : (
            allPermissions &&
            allPermissions.length > 0 && (
              <div className="grid grid-cols-3 gap-4">
                {allPermissions.map((permissions) => (
                  <div
                    key={permissions.ID ?? permissions.PermissionName}
                    className={`border rounded-lg p-4 ${
                      rolePermissions?.some((rp) => rp.ID === permissions.ID) ||
                      highlightedPermissions.includes(permissions.ID)
                        ? 'bg-blue-100 text-blue-900 border-blue-500'
                        : 'bg-gray-100 text-gray-700'
                    }`}
                    onClick={() => handlePermissionClick(permissions)}
                  >
                    <h2 className="text-lg font-medium">
                      {permissions.PermissionName}
                    </h2>
                    <p className="text-sm text-gray-500">
                      {permissions.Description}
                    </p>
                  </div>
                ))}
              </div>
            )
          )}
        </div>
      </div>
    </>
  );
}
