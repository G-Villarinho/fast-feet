import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Loader2, LogOut } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { logout } from "@/api/logout";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";

function getInitials(name: string): string {
  return name
    .split(" ")
    .map((word) => word.charAt(0).toUpperCase())
    .slice(0, 2)
    .join("");
}

interface UserButtonProps {
  userFullName?: string;
  isFetchingUser?: boolean;
}

export function UserButton({ userFullName, isFetchingUser }: UserButtonProps) {
  const navigate = useNavigate();

  const { mutateAsync: logoutFn, isPending } = useMutation({
    mutationFn: logout,
  });

  async function handleLogout() {
    try {
      await logoutFn();
      navigate("/login");
    } catch {
      toast.error(
        "Ocorreu um erro inesperado ao tentar realizar o login. Por favor, tente novamente."
      );
    }
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="outline-none">
        <Avatar className="bg-gray-100 border border-gray-300">
          {isFetchingUser ? (
            <AvatarFallback className="flex items-center justify-center bg-gray-200 text-gray-700">
              <Loader2 className="animate-spin w-5 h-5" />
            </AvatarFallback>
          ) : (
            <AvatarFallback className="bg-white text-violet-700 font-bold">
              {userFullName ? getInitials(userFullName) : "?"}
            </AvatarFallback>
          )}
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuItem
          onClick={handleLogout}
          disabled={isPending}
          className={isPending ? "opacity-50 cursor-not-allowed" : ""}
        >
          {isPending ? (
            <Loader2 className="mr-2 size-4 animate-spin" />
          ) : (
            <LogOut className="mr-2 size-4" />
          )}
          Sair
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
