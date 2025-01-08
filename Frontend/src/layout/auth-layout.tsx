import { ModeToggle } from "@/components/mode-toggle";
import { Outlet } from "react-router-dom";
import { LanguageSelector } from "@/components/language-toggle";
import { Card } from "@/components/ui/card";

const AuthLayout = () => {

  return (
    <Card className={`flex items-center justify-center min-h-screen relative `}>
      <div className="w-full max-w-md mx-4 shadow-md rounded-md overflow-hidden">
        <Outlet />
      </div>
      <div className="absolute bottom-4 right-4 p-4 flex space-x-4">
        <LanguageSelector />
        <ModeToggle />
      </div>
    </Card>
  );
};

export default AuthLayout;
