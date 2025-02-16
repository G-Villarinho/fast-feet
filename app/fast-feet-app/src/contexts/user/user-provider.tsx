import { getUser, GetUserResponse } from "@/api/get-user";
import { useQuery } from "@tanstack/react-query";
import { ReactNode, useState } from "react";
import { UserContext } from "@/contexts/user/user-context";

interface UserProviderProps {
  children: ReactNode;
}

export function UserProvider({ children }: UserProviderProps) {
  const [user, setUser] = useState<GetUserResponse | null>(null);

  const { data: result, isFetching } = useQuery({
    queryKey: ["user"],
    queryFn: getUser,
    enabled: user === null || user === undefined,
  });

  if (result && user !== result) {
    setUser(result);
  }

  function logout() {
    setUser(null);
  }

  return (
    <UserContext.Provider
      value={{ user: user, isFetchingUser: isFetching, logout: logout }}
    >
      {children}
    </UserContext.Provider>
  );
}
