import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { Link } from "react-router-dom";

interface NavButtonProps {
  to: string;
  label: string;
  isActive: boolean;
}

export function NavButton({ to, label, isActive }: NavButtonProps) {
  return (
    <Button
      asChild
      size="sm"
      variant="outline"
      className={cn(
        "w-full lg:w-auto justify-between font-medium hover:bg-white/20 hover:text-white border-none focus-visible:ring-transparent outline-none text-white focus:bg-white/30 transtion",
        isActive ? "bg-white/10 text-white" : "bg-transparent"
      )}
    >
      <Link to={to}>{label}</Link>
    </Button>
  );
}
