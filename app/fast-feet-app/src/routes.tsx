import { createBrowserRouter } from "react-router-dom";
import AuthLayout from "@/pages/auth/auth-layout";
import Login from "@/pages/auth/login/login";
import { AppLayout } from "@/pages/app/app-layout";
import { NotFound } from "@/pages/404";
import { CreateOrder } from "@/pages/app/orders/create-order/create-order";
import { Recipients } from "@/pages/app/recipients/recipients";
import { CreateRecipient } from "@/pages/app/recipients/create-recipient/create-recipient";
import { Orders } from "@/pages/app/orders/orders/orders";

export const router = createBrowserRouter([
  {
    path: "/",
    errorElement: <NotFound />,
    element: <AppLayout />,
    children: [
      {
        path: "/orders",
        element: <Orders />,
      },
      {
        path: "/create-order",
        element: <CreateOrder />,
      },
      {
        path: "/recipients",
        element: <Recipients />,
      },
      {
        path: "/create-recipient",
        element: <CreateRecipient />,
      },
    ],
  },
  {
    path: "/",
    element: <AuthLayout />,
    children: [
      {
        path: "/login",
        element: <Login />,
      },
    ],
  },
]);
