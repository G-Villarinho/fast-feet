import { Role } from "@/api/get-user";

interface WelcomeMessageProps {
  userFullName?: string;
  isFetchingUser: boolean;
  userRole?: Role;
}

export function WelcomeMessage({
  userFullName,
  isFetchingUser,
  userRole,
}: WelcomeMessageProps) {
  return (
    <div className="space-y-2 mb-4">
      <h2 className="text-2xl lg:text-4xl text-white font-medium">
        Seja bem-vindo(a) {isFetchingUser ? ", " : " "}
        {userFullName} 👋🏻
      </h2>
      <p className="text-sm lg:text-base text-gray-400">
        {isFetchingUser
          ? "..."
          : userRole === Role.admin
          ? "Gerencie destinatários, encomendas e entregadores de forma simples e eficiente."
          : "Confira suas encomendas e inicie suas entregas com rapidez e segurança!"}
      </p>
    </div>
  );
}
