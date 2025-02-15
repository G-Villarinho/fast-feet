import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Plus } from "lucide-react";
import { Helmet } from "react-helmet-async";
import { Link } from "react-router-dom";
import { OrderCard } from "./order-card";
import { OrderStatus } from "@/api/get-orders";

export const mockOrders = [
  {
    id: "1",
    title: "Encomenda 001",
    createdAt: "2025-02-13 10:00:00",
    status: OrderStatus.waiting,
  },
  {
    id: "2",
    title: "Encomenda 002",
    createdAt: "2025-02-13 12:00:00",
    status: OrderStatus.picknUp,
  },
  {
    id: "3",
    title: "Encomenda 003",
    createdAt: "2025-02-14 09:00:00",
    status: OrderStatus.done,
  },
];

export function Orders() {
  return (
    <>
      <Helmet title="pedidos" />
      <div className="max-w-screen-2xl mx-auto w-full pb-10 -mt-24">
        <Card className="border-none drop-shadow-sm">
          <CardHeader className="gap-y-2 lg:flex-row lg:items-center lg:justify-between">
            <CardTitle className="text-xl line-clamp-1">Encomendas</CardTitle>
            <Button asChild size="sm">
              <Link to="/create-order">
                <Plus className="size-4 mr-2" />
                Criar encomenda
              </Link>
            </Button>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
              {mockOrders.map((order) => (
                <OrderCard key={order.id} order={order} />
              ))}
            </div>
          </CardContent>
        </Card>
      </div>
    </>
  );
}
