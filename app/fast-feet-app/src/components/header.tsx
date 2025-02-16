import { HeaderLogo } from "@/components/header-logo";
import { Navigation } from "@/components/navigation";
import { WelcomeMessage } from "./welcome-message";
import { useContext } from "react";
import { UserContext } from "@/contexts/user/user-context";
import { UserButton } from "./user-button";

export function Header() {
  const { user, isFetchingUser, logout } = useContext(UserContext);

  return (
    <header className="bg-gradient-to-b from-violet-800 to-violet-600 px-4 py-8 lg:px-14 pb-36">
      <div className="max-w-screen-2xl mx-auto">
        <div className="w-full flex items-center justify-between mb-14">
          <div className="flex items-center lg:gap-x-16">
            <HeaderLogo />
            <Navigation />
          </div>
          <UserButton
            userFullName={user?.fullName}
            isFetchingUser={isFetchingUser}
            logoutUser={logout}
          />
        </div>
        <WelcomeMessage
          userFullName={user?.fullName}
          isFetchingUser={isFetchingUser}
          userRole={user?.role}
        />
      </div>
    </header>
  );
}
