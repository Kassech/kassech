import { ModeToggle } from "@/components/mode-toggle";
import { Outlet, useNavigate } from "react-router-dom";
import { LanguageSelector } from "@/components/language-toggle";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useAuthCheck } from "@/hooks/useAuth";
import { useEffect } from "react";

const DashboardLayout = () => {
  const navigate = useNavigate();
  
  // Handle logout by removing the accessToken from localStorage
  const handelLogout = () => {
    localStorage.removeItem("accessToken");
    navigate("/login"); // Navigate to login page on logout
  };

  const { mutate: checkAuth, isLoading, isError, error } = useAuthCheck();

  useEffect(() => {
    checkAuth();
  }, [checkAuth]);

  // Redirect to login page if authentication fails or there's no response
  useEffect(() => {
    if (isError) {
      console.error("Authentication failed:", error);
      navigate("/login");
    }
  }, [isError, error, navigate]);

  if (isLoading) {
    return <div>Loading...</div>; // Show loading state while checking auth
  }

  return (
    <Card className="flex items-center justify-center min-h-screen relative">
      <div className="w-full max-w-md mx-4 shadow-md rounded-md overflow-hidden">
        <Outlet />
      </div>
      <div className="absolute bottom-4 right-4 p-4 flex space-x-4">
        <Button onClick={handelLogout}>LOGOUT</Button>
        <LanguageSelector />
        <ModeToggle />
      </div>
    </Card>
  );
};

export default DashboardLayout;
