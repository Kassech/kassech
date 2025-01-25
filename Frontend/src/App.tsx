import {
  BrowserRouter as Router,
  Route,
  Routes,
  Navigate,
} from 'react-router-dom';
import Login from '@/pages/authentication/login-form';
import SignUp from '@/pages/authentication/register-form';
import AuthLayout from './layout/auth-layout';
import DashboardLayout from './layout/dashboard-layout';
import { ThemeProvider } from '@/components/theme-provider';
import Dashboard from './pages/dashboard';
import NotFoundPage from './pages/error/404';
import Error500Page from './pages/error/500';
import VehicleRegistration from './pages/Vehicle Registration';
import CarOwnerRegistration from './pages/carOwnerRegistration';
import DriverForm from './pages/driver';
import QueueManagerForm from './pages/QueueManagerForm';
import Users from './pages/Users';
// import LoadingSpinner from "./components/loading-spinner";

import StationsPage from './pages/Station Managing';
import RolePermission from './pages/RolePermission';
import { Toaster } from 'sonner';
import DriverPage from './pages/driver';
import RoutePage from './pages/route';
import DriverVehicleMapping from './pages/Driver-VehicleMapping';
import VehicleTracking from './pages/VehicleTracking';

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
            <Route path="/user" element={<Users />} />
            <Route path="/stations" element={<StationsPage />} />
            <Route path="/route" element={<RoutePage />} />
            <Route path="/rolepermission" element={<RolePermission />} />
            <Route
              path="/driverVehicleMap"
              element={<DriverVehicleMapping />}
            />
            <Route path="/vehicleTracking" element={<VehicleTracking />} />
          </Route>

          {/* Catch-all route for 404 */}
          <Route path="*" element={<NotFoundPage />} />
          <Route path="/404" element={<NotFoundPage />} />
          <Route path="/500" element={<Error500Page />} />
        </Routes>
      </Router>
    </ThemeProvider>
  );
}

export default App;
