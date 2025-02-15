import { Outlet } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Globe } from "lucide-react";
import Logo from "@/assets/logo.svg";

export default function AuthLayout() {
  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-muted py-6  lg:px-56  antialiased w-full">
      <nav className="w-full flex items-center justify-between px-6 md:px-10">
        <div className="flex items-center gap-2">
          <img src={Logo} alt="Logo" className="h-10" />
          <span className="text-xl font-bold">Fast Feet</span>
        </div>
        <div className="hidden md:flex items-center gap-6">
          <Button variant="ghost" size="icon">
            <Globe className="h-[1.5rem] w-[1.5rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
          </Button>

          <Button size="lg">Solicitar Demo</Button>
        </div>
      </nav>

      <div className="flex flex-1 w-full flex-col gap-6 items-center justify-center">
        <Outlet />
      </div>
    </div>
  );
}
