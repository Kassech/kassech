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
} from "@/components/ui/alert-dialog";
import useStore from "./store";
import { useState } from "react";
import { useToast } from "@/hooks/use-toast";


export default function RolePermission() {
    const { setDialogClose, setDialogOpen, isDialogOpen }=useStore();
  const { data: allRoles, isLoading: isRolesLoading } = useGetAllRole();
  const { data: allPermissions, isLoading: isPermissionsLoading } =
    useGetAllPermission();
     const [selectedRole, setSelectedRole] = useState<number | null>(null);
      const { mutate: deleteRole } = useDeleteRole();
       const { toast } = useToast();

    function handleEditRole(role: { ID: number; RoleName: string; Description: string; Permission: string; CreatedAt: string; UpdatedAt: string; DeletedAt?: string | null; }): void {
        
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
    }

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
                      <Button variant="ghost" className="px-2">
                        ...
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                      <DropdownMenuItem onClick={() => handleEditRole(role)}>
                        Edit
                      </DropdownMenuItem>
                      <DropdownMenuItem onClick={() =>{ setSelectedRole(role.ID);setDialogOpen()}}>
                        Delete
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </div>
              ))
            )}
          </div>
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
                    <p className="text-sm text-gray-500">
                      {permission.Description}
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
