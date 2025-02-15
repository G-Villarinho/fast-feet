import { useLocation } from "react-router-dom";
import { NavButton } from "@/components/nav-button";

const routes = [
  {
    to: "/orders",
    label: "Encomendas",
  },
  {
    to: "/admins",
    label: "Admins",
  },
  {
    to: "/deliveries",
    label: "Entregadores",
  },
  {
    to: "/recipients",
    label: "Destinatários",
  },
  {
    to: "/settings",
    label: "Configurações",
  },
];

export function Navigation() {
  const location = useLocation();

  const pathname = location.pathname;

  return (
    <nav className="hidden lg:flex items-center gap-x-2 overflow-x-auto">
      {routes.map((route) => (
        <NavButton
          key={route.to}
          to={route.to}
          label={route.label}
          isActive={pathname === route.to}
        />
      ))}
    </nav>
  );
}
