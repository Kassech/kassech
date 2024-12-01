import Login from "@/components/Authentication/login-form";
import SignUp from "@/components/Authentication/register-form";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import AuthLayout from "./layout/auth-layout";
import { ThemeProvider } from "@/components/theme-provider"
import DashboardLayout from "./layout/dashboard-layout";

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <Router>
        <Routes>
          <Route element={<AuthLayout/>}>
            <Route path="/register" element={<SignUp />} />
            <Route path="/login" element={<Login />} />
          </Route>
          <Route element={<DashboardLayout/>}>
            <Route path="/dashboard" element={<h1>HAYE</h1>} />
          </Route>
        </Routes>
      </Router>
    </ThemeProvider>
  );
}

export default App;
