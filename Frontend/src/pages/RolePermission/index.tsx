import {
  useGetAllRole,
  useDeleteRole,
  useCreateRole,
  useUpdateRole,
} from "@/services/roleServices";
import { useGetAllPermission } from "@/services/permissionServices";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
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
} from "@/components/ui/alert-dialog";
import { useToast } from "@/hooks/use-toast";
import {
  useDialogStore,
  useRoleStore,
  useCardStore,
} from "@/store/useRoleStore";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function RolePermission() {
  const { toast } = useToast();
  const { toggleCard, showCard }=useCardStore();
  const {
    isDialogOpen,
    setDialogOpen,
    setDialogClose,
    isEditRoleDialogOpen,
    setEditRoleDialogOpen,
    setEditRoleDialogClose,
  } = useDialogStore();

  const { data: allRoles, isLoading: isRolesLoading } = useGetAllRole();
  const { data: allPermissions, isLoading: isPermissionsLoading } = useGetAllPermission();
  const { mutate: deleteRole } = useDeleteRole();
  const { mutate: editRole } = useUpdateRole();
  const { mutate: addRole } = useCreateRole();

  const {
    currentRole,
    selectedRole,
    newRole,
    setNewRole,
    updateNewRoleField,
    setCurrentRole,
    setSelectedRole,
    updateCurrentRoleField,
  } = useRoleStore();


  const handleAddRole = () =>{
     toggleCard();
 if (newRole.RoleName && newRole.Description) {
   addRole(
     {
       RoleName: newRole.RoleName,
       Description: newRole.Description,
     },
     {
       onSuccess: (data) => {
         console.log("API response data:", data);
         toast({
           description: "Role Added Successfully.",
         });
       },
       onError: (error) => {
         console.error(`Failed to add role: ${error}`);
       },
     }
   );
 } else {
   console.error("Role ID is undefined.");
 }
  }

  const confirmDelete = () => {
    if (selectedRole !== null) {
      deleteRole(selectedRole, {
        onSuccess: () => {
          toast({
            description: "Role Deleted Successfully.",
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
            console.log("API response data:", data);
            toast({
              description: "Role Edited Successfully.",
            });
          },
          onError: (error) => {
            console.error(`Failed to edit role: ${error}`);
          },
        }
      );
    } else {
      console.error("Role ID is undefined.");
    }
  };
        
          console.log("permissions",allPermissions);
      


  return (
    <>
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
                  <Button variant="outline" className="relative">
                    {role.RoleName}
                  </Button>
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" className="px-2 border-none">
                        ...
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
                            updateNewRoleField("RoleName", e.target.value)
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
                            updateNewRoleField("Description", e.target.value)
                          }
                        />
                      </div>
                    </div>
                    <AlertDialogFooter>
                      <AlertDialogCancel
                        onClick={() => {
                          console.log(
                            "Cancel button clicked, closing dialog..."
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
                      console.log("Cancel button clicked, closing dialog...");
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
                        updateCurrentRoleField("RoleName", e.target.value)
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
                        updateCurrentRoleField("Description", e.target.value)
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
                {allPermissions.map((permission) => (
                  <div
                    key={permission.ID ?? permission.PermissionName}
                    className="border rounded-lg p-4"
                  >
                    <h2 className="text-lg font-medium">
                      {permission.PermissionName}
                    </h2>
                    {/* <p className="text-sm text-gray-500">
                      {permission.Description}
                    </p> */}
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
