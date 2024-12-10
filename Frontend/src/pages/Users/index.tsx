import React, { useEffect, useState } from "react";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { MoreHorizontal } from "lucide-react";

import { fetchUserData } from "../../services/userService"; // Replace with your actual service
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import Header from "@/components/header";
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

interface User {
  ID: number;
  FirstName: string;
  LastName: string;
  Email: string;
  PhoneNumber: string;
  IsOnline: boolean;
  LastLoginDate: string;
}

const UserList: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [search, setSearch] = useState<string>("");
  const [filter, setFilter] = useState<string>("all");
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [isDialogOpen, setDialogOpen] = useState(false);
  const [userToDelete, setUserToDelete] = useState<User | null>(null);
  const [userToEdit, setUserToEdit] = useState<User | null>(null);
  const itemsPerPage = 10;

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const data = await fetchUserData();
        setUsers(data.users || []);
      } catch (error) {
        console.error("Failed to fetch users:", error);
      }
    };

    fetchUsers();
  }, []);

  const filteredUsers = users.filter((user) => {
    const matchesSearch = `${user.FirstName} ${user.LastName} ${user.Email}`
      .toLowerCase()
      .includes(search.toLowerCase());
    const matchesFilter =
      filter === "all" ||
      (filter === "online" && user.IsOnline) ||
      (filter === "offline" && !user.IsOnline);

    return matchesSearch && matchesFilter;
  });

  const totalItems = filteredUsers.length;
  const totalPages = Math.ceil(totalItems / itemsPerPage);
  const paginatedUsers = filteredUsers.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage
  );

  const handlePageChange = (page: number) => {
    if (page >= 1 && page <= totalPages) {
      setCurrentPage(page);
    }
  };

  const handleEdit = (user: User) => {
    setUserToEdit(user);
  };

  const saveEditedUser = () => {
    if (userToEdit) {
      console.log("Saving edited user:", userToEdit);
      setUsers((prevUsers) =>
        prevUsers.map((u) => (u.ID === userToEdit.ID ? userToEdit : u))
      );
      setUserToEdit(null);
    }
  };

  const handleDeleteClick = (user: User) => {
    setUserToDelete(user);
    setDialogOpen(true);
  };

  const confirmDelete = () => {
    if (userToDelete) {
      console.log("Deleting user with ID:", userToDelete.ID);
      setUsers((prevUsers) =>
        prevUsers.filter((u) => u.ID !== userToDelete.ID)
      );
      setDialogOpen(false);
    }
  };

  const paths = [
    { name: "Home", href: "/" },
    { name: "Dashboard", href: "/b" },
  ];

  return (
    <>
      <Header paths={paths} />
      <div className="p-6">
        <h1 className="text-xl font-semibold mb-4">User Management</h1>
        <div className="flex gap-4 mb-6">
          <Input
            type="text"
            placeholder="Search users..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-full max-w-md"
          />
          <Select
            onValueChange={(value) => setFilter(value)}
            defaultValue="all"
          >
            <SelectTrigger className="w-[180px]">
              <span>Filter by Status</span>
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">All</SelectItem>
              <SelectItem value="online">Online</SelectItem>
              <SelectItem value="offline">Offline</SelectItem>
            </SelectContent>
          </Select>
        </div>
        {filteredUsers.length > 0 ? (
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>ID</TableHead>
                <TableHead>Name</TableHead>
                <TableHead>Email</TableHead>
                <TableHead>Phone</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {paginatedUsers.map((user) => (
                <TableRow key={user.ID}>
                  <TableCell>{user.ID}</TableCell>
                  <TableCell>
                    {user.FirstName} {user.LastName}
                  </TableCell>
                  <TableCell>{user.Email}</TableCell>
                  <TableCell>{user.PhoneNumber}</TableCell>
                  <TableCell>
                    <span
                      className={
                        user.IsOnline ? "text-green-600" : "text-gray-500"
                      }
                    >
                      {user.IsOnline ? "Online" : "Offline"}
                    </span>
                  </TableCell>
                  <TableCell>
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button variant="ghost" className="h-8 w-8 p-0">
                          <span className="sr-only">Open menu</span>
                          <MoreHorizontal />
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        <DropdownMenuItem onClick={() => handleEdit(user)}>
                          Edit
                        </DropdownMenuItem>
                        <DropdownMenuItem
                          onClick={() => handleDeleteClick(user)}
                        >
                          Delete
                        </DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        ) : (
          <div className="text-center text-gray-600 mt-6">
            No users found matching your criteria.
          </div>
        )}
        <div className="mt-4 flex justify-center">
          <Button
            variant="outline"
            onClick={() => handlePageChange(currentPage - 1)}
            disabled={currentPage === 1}
          >
            Previous
          </Button>
          <span className="mx-4">
            Page {currentPage} of {totalPages}
          </span>
          <Button
            variant="outline"
            onClick={() => handlePageChange(currentPage + 1)}
            disabled={currentPage === totalPages}
          >
            Next
          </Button>
        </div>
      </div>
      {/* Edit User Sheet */}
      {userToEdit && (
        <Sheet
          open={Boolean(userToEdit)}
          onOpenChange={() => setUserToEdit(null)}
        >
          <Sheet
            open={Boolean(userToEdit)}
            onOpenChange={() => setUserToEdit(null)}
          >
            <SheetTrigger asChild>
            </SheetTrigger>
            <SheetContent>
              <SheetHeader>
                <SheetTitle>Edit User</SheetTitle>
                <SheetDescription>
                  Modify the details of the selected user. Save the changes when
                  you're done.
                </SheetDescription>
              </SheetHeader>
              <div className="grid gap-4 py-4">
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="firstName" className="text-right">
                    First Name
                  </Label>
                  <Input
                    id="firstName"
                    value={userToEdit?.FirstName || ""}
                    onChange={(e) =>
                      setUserToEdit((prev) =>
                        prev ? { ...prev, FirstName: e.target.value } : null
                      )
                    }
                    className="col-span-3"
                  />
                </div>
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="lastName" className="text-right">
                    Last Name
                  </Label>
                  <Input
                    id="lastName"
                    value={userToEdit?.LastName || ""}
                    onChange={(e) =>
                      setUserToEdit((prev) =>
                        prev ? { ...prev, LastName: e.target.value } : null
                      )
                    }
                    className="col-span-3"
                  />
                </div>
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="email" className="text-right">
                    Email
                  </Label>
                  <Input
                    id="email"
                    value={userToEdit?.Email || ""}
                    onChange={(e) =>
                      setUserToEdit((prev) =>
                        prev ? { ...prev, Email: e.target.value } : null
                      )
                    }
                    className="col-span-3"
                  />
                </div>
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="phone" className="text-right">
                    Phone Number
                  </Label>
                  <Input
                    id="phone"
                    value={userToEdit?.PhoneNumber || ""}
                    onChange={(e) =>
                      setUserToEdit((prev) =>
                        prev ? { ...prev, PhoneNumber: e.target.value } : null
                      )
                    }
                    className="col-span-3"
                  />
                </div>
              </div>
              <SheetFooter>
                <SheetClose asChild>
                  <Button type="submit" onClick={saveEditedUser}>
                    Save Changes
                  </Button>
                </SheetClose>
                <SheetClose asChild>
                  <Button variant="outline" onClick={() => setUserToEdit(null)}>
                    Cancel
                  </Button>
                </SheetClose>
              </SheetFooter>
            </SheetContent>
          </Sheet>
        </Sheet>
      )}
      {/* Alert Dialog for Deletion */}
      {isDialogOpen && (
        <AlertDialog open={isDialogOpen} onOpenChange={setDialogOpen}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Confirm Delete</AlertDialogTitle>
              <AlertDialogDescription>
                Are you sure you want to delete this user? This action cannot be
                undone.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setDialogOpen(false)}>
                Cancel
              </AlertDialogCancel>
              <AlertDialogAction onClick={confirmDelete}>
                Delete
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      )}
    </>
  );
};

export default UserList;
