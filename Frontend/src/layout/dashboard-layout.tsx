import { AppSidebar } from "@/components/app-sidebar";
import {
  SidebarInset,
  SidebarProvider,
} from "@/components/ui/sidebar";
import { Outlet, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { useAuthCheck } from "@/hooks/useAuth";
import LoadingSpinner from "@/components/loading-spinner";

export default function DashboardLayout() {
  const { mutate, isLoading, isError} = useAuthCheck(); // Using the hook
  const navigate = useNavigate(); // For redirection

  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);

  // Check if the user is authenticated when the component mounts
  useEffect(() => {
    // Trigger the auth check when the component is mounted
    const checkAuth = async () => {
      try {
        await mutate(); // This will trigger the auth check
        setIsAuthenticated(true); // If the request is successful, the user is authenticated
      } catch (err) {
        setIsAuthenticated(false); // If error occurs, user is not authenticated
      }
    };

    checkAuth();
  }, [mutate]);

  // Redirect to login page if not authenticated
  useEffect(() => {
    if (isAuthenticated === false) {
      navigate("/login"); // Redirect to login page if not authenticated
    }
  }, [isAuthenticated, navigate]);


  if (isLoading) {
    return <LoadingSpinner />
  }
  // Render the layout only if the user is authenticated
  if (isAuthenticated === null || isAuthenticated === false) {
    return null; // You can optionally return a loading screen or redirect before this point.
  }
  if (isError) {
    navigate("/login"); // Redirect to login page if not authenticated
  }
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <Outlet /> {/* Render the child route components */}
      </SidebarInset>
    </SidebarProvider>
  );
}
