import React from 'react';
import { Route, Navigate } from 'react-router-dom';
import { useUserStore } from '@/store/userStore'; // Assuming you're using Zustand or other state management

// Define a ProtectedRoute that takes a permission prop and only allows access if the user has the required permission
interface ProtectedRouteProps {
  element: React.ReactNode;
  permission: string;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
  element,
  permission,
}) => {
  const user = useUserStore((state) => state.user);

  // Check if the user is logged in and if they have the required permission
  const hasPermission = user?.permissions.some(
    (perm: any) => perm.permission_name === permission
  );

  if (!user) {
    // Redirect to login if not logged in
    return <Navigate to="/login" />;
  }

  if (!hasPermission) {
    // Redirect to 403 if user doesn't have permission
    return <Navigate to="/403" />;
  }

  return <>{element}</>;
};

export default ProtectedRoute;
