import { Header } from "@/components/header";
import { UserProvider } from "@/contexts/user/user-provider";
import { api } from "@/lib/axios";
import { isAxiosError } from "axios";
import { useLayoutEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";

export function AppLayout() {
  const navigate = useNavigate();

  useLayoutEffect(() => {
    const interceptorId = api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (isAxiosError(error)) {
          const status = error.response?.status;
          const code = error.response?.data.code;

          if (status === 401 && code === "UNAUTHORIZED") {
            navigate("/login", {
              replace: true,
            });
          }
        }
        return Promise.reject(error);
      }
    );
    return () => {
      api.interceptors.response.eject(interceptorId);
    };
  }, [navigate]);

  return (
    <UserProvider>
      <div className="flex min-h-screen flex-col antialiased">
        <Header />
        <div className="flex flex-1 flex-col gap-4 ">
          <Outlet />
        </div>
      </div>
    </UserProvider>
  );
}
