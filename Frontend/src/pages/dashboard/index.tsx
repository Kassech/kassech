// pages/Dashboard.jsx

import Header from "@/components/header";
import Overview from "./overview";
import DemandingRoutesGraph from "./DemandingRoutesGraph";

export default function Dashboard() {
    const paths = [
        { name: "Home", href: "/" },
        { name: "Dashboard", href: "/dashboard" },
      ];
  return (
    <>
      <Header paths={paths} />
      <Overview />
      <DemandingRoutesGraph />

    </>
  );
}
