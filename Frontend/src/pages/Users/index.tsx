import React, { useEffect, useState } from "react";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
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
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import Header from "@/components/header";

// Define a type for user data
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
  const itemsPerPage = 10;

  // Fetch users from API
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

  // Filter and Search Logic
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

  // Pagination Logic
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
  
    console.log("Editing user:", user);

    // editUserAPI(user);
  };

  const handleDelete = (userId: number) => {
   
    console.log("Deleting user with ID:", userId);
    // deleteUserFromAPI(userId);
  
    // setUsers((prevUsers) => prevUsers.filter((u) => u.ID !== userId));
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

        {/* Search and Filter */}
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

        {/* Table or No User Found */}
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
                    {/* Actions */}
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button variant="ghost" className="h-8 w-8 p-0">
                          <span className="sr-only">Open menu</span>
                          <MoreHorizontal />
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        <DropdownMenuItem>Edit</DropdownMenuItem>
                        <DropdownMenuItem>Delete</DropdownMenuItem>
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

        {/* Pagination */}
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
    </>
  );
};

export default UserList;
