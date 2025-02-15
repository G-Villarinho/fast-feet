import { Helmet, HelmetProvider } from "react-helmet-async";
import { RouterProvider } from "react-router-dom";
import { Toaster } from "sonner";
import { queryClient } from "@/lib/react-query";
import { QueryClientProvider } from "@tanstack/react-query";
import { router } from "@/routes";

import "@/index.css";

export function App() {
  return (
    <HelmetProvider>
      <Helmet titleTemplate="%s | fast.feet" />
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
        <Toaster richColors />
      </QueryClientProvider>
    </HelmetProvider>
  );
}
