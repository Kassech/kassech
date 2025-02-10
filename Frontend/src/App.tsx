import {
  BrowserRouter as Router,
  Route,
  Routes,
  Navigate,
} from 'react-router-dom';
import { useUserStore } from '@/store/userStore';
import Login from '@/pages/authentication/login-form';
import SignUp from '@/pages/authentication/register-form';
import AuthLayout from './layout/auth-layout';
import DashboardLayout from './layout/dashboard-layout';
import { ThemeProvider } from '@/components/theme-provider';
import Dashboard from './pages/Dashboardd';
import NotFoundPage from './pages/error/404';
import Error500Page from './pages/error/500';
import VehicleRegistration from './pages/Vehicle Registration';
import CarOwnerRegistration from './pages/carOwnerRegistration';
import QueueManagerForm from './pages/QueueManagerForm';
import Users from './pages/Users';
import StationsPage from './pages/Station Managing';
import RolePermission from './pages/RolePermission';
import { Toaster } from 'sonner';
import DriverPage from './pages/driver';
import RoutePage from './pages/route';
import PathPage from './pages/path';
import VehicleTracking from './pages/VehicleTracking';
import QueueManagerPaths from './pages/QueueManagerPathsRegistration';
import Delegation from './pages/Delegation';
import TrackAllVehicle from './pages/TrackAllVehicle';
import TrackVehicle from './pages/TrackVehicle';
import TrackNearbyPage from './pages/TrackNearby';
import TrackPathPage from './pages/TrackByPath';
import NoPermissionPage from './pages/error/NoPermission';

/**
 * Role-based access control for routes
 */
const routePermissions: Record<string, string[]> = {
  '/user': ['ViewUser'],
  '/delegation': ['CreateDelegation'],
  '/route': ['ViewRoute', 'CreateRoute', 'EditRoute', 'DeleteRoute'],
  '/rolepermission': [
    'ViewRole',
    'CreateRole',
    'EditRole',
    'DeleteRole',
    'Assign Permission',
  ],
  '/vehicleTracking': ['VehicleTracking'],
};

/**
 * PrivateRoute: Restricts access based on user permissions
 */
function PrivateRoute({
  element,
  path,
}: {
  element: JSX.Element;
  path: string;
}) {
  const user = useUserStore((state) => state.user);

  if (!user) return <Navigate to="/login" />;

  const userPermissions = new Set(user.permissions || []);
  const requiredPermissions = routePermissions[path] || [];

  if (
    requiredPermissions.length === 0 ||
    requiredPermissions.some((perm) => userPermissions.has(perm))
  ) {
    return element;
  }

  return <Navigate to="/403" />;
}

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <Toaster />
      <Router>
        <Routes>
          {/* Redirect / to /dashboard */}
          <Route path="/" element={<Navigate to="/dashboard" />} />

          {/* Authentication routes */}
          <Route element={<AuthLayout />}>
            <Route path="/register" element={<SignUp />} />
            <Route path="/login" element={<Login />} />
          </Route>

          {/* Dashboard routes */}
          <Route element={<DashboardLayout />}>
            <Route path="/dashboard" element={<Dashboard />} />
            <Route
              path="/VehicleRegistration"
              element={<VehicleRegistration />}
            />
            <Route
              path="/carOwnerRegistration"
              element={<CarOwnerRegistration />}
            />
            <Route path="/driver" element={<DriverPage />} />
            <Route path="/queueManagerForm" element={<QueueManagerForm />} />
            <Route path="/stations" element={<StationsPage />} />
            <Route path="/paths" element={<PathPage />} />
            <Route path="/QueueManagerPaths" element={<QueueManagerPaths />} />
            <Route path="/trackall" element={<TrackAllVehicle />} />
            <Route path="/trackOne" element={<TrackVehicle />} />
            <Route path="/trackNearby" element={<TrackNearbyPage />} />
            <Route path="/trackByPath" element={<TrackPathPage />} />

            {/* Role-based protected routes */}
            <Route
              path="/user"
              element={<PrivateRoute path="/user" element={<Users />} />}
            />
            <Route
              path="/route"
              element={<PrivateRoute path="/route" element={<RoutePage />} />}
            />
            <Route
              path="/rolepermission"
              element={
                <PrivateRoute
                  path="/rolepermission"
                  element={<RolePermission />}
                />
              }
            />
            <Route
              path="/vehicleTracking"
              element={
                <PrivateRoute
                  path="/vehicleTracking"
                  element={<VehicleTracking />}
                />
              }
            />
            <Route
              path="/delegation"
              element={
                <PrivateRoute path="/delegation" element={<Delegation />} />
              }
            />
          </Route>

          {/* Error routes */}
          <Route path="*" element={<NotFoundPage />} />
          <Route path="/404" element={<NotFoundPage />} />
          <Route path="/500" element={<Error500Page />} />
          <Route path="/403" element={<NoPermissionPage />} />
        </Routes>
      </Router>
    </ThemeProvider>
  );
}

export default App;
