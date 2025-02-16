import { GetUserResponse } from "@/api/get-user";
import { createContext } from "react";

interface UserContextType {
  user?: GetUserResponse | null;
  logout: () => void
  isFetchingUser: boolean;

}

export const UserContext = createContext({} as UserContextType);
