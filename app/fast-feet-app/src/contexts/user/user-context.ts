import { GetUserResponse } from "@/api/get-user";
import { createContext } from "react";

interface UserContextType {
  user?: GetUserResponse | null;
  isFetchingUser: boolean;
}

export const UserContext = createContext({} as UserContextType);
