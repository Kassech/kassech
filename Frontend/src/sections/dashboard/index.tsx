// pages/Dashboard.jsx

import Overview from "./Overview";
import DemandingRoutesGraph from "./DemandingRoutesGraph";
import LoginLogsTable from "./LoginLogsTable";

export default function Dashboard() {
  return (
    <>
      <Overview />
      <LoginLogsTable/>
      <DemandingRoutesGraph />

    </>
  );
}
