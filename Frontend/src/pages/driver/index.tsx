// pages/driver/Driver.jsx

import Header from "@/components/header";
import Driver from "@/sections/driver";

export default function DriverPage() {
    const paths = [
        { name: "Home", href: "/" },
        { name: "Driver", href: "/driver" },
      ];
  return (
    <>
      <Header paths={paths} />
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        <div className="grid auto-rows-min gap-4 md:grid-cols-3">
          <div className="aspect-video rounded-xl" />
          <div className="aspect-video rounded-xl" />
          <div className="aspect-video rounded-xl" />
        </div>
        <div className="min-h-[100vh] flex-1 rounded-xl md:min-h-min" >
            <Driver />
        </div>
      </div>
    </>
  );
}
