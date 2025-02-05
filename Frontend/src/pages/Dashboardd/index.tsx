// pages/Dashboard.jsx
import Header from "@/components/header";
import Report from "@/sections/dashboard";

export default function Dashboard() {
    const paths = [
        { name: "Home", href: "/" },
        { name: "Dashboard", href: "/dashboard" },
      ];
  return (
    <>
      <Header paths={paths} />
      <Report />
    </>
  );
}
