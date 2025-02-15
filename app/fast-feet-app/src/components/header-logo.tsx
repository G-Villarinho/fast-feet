import logo from "@/assets/logo.svg";

import { Link } from "react-router-dom";

export function HeaderLogo() {
  return (
    <Link to="/orders" className="flex items-center gap-2">
      <img src={logo} />
      <span className="text-2xl font-semibold text-white">Fast Feet</span>
    </Link>
  );
}
