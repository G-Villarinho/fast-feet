import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Plus } from "lucide-react";
import { Link, useSearchParams } from "react-router-dom";
import { OrderCard } from "./order-card";
import { getOrders } from "@/api/get-orders";
import { z } from "zod";
import { useQuery } from "@tanstack/react-query";

import packageImg from "@/assets/package.svg";
import { OrderCardSkeleton } from "./order-card-skeleton";
import { Helmet } from "react-helmet-async";
import { useContext } from "react";
import { UserContext } from "@/contexts/user/user-context";
import { Role } from "@/api/get-user";

const ORDERS_LIMIT_PAGE = 9;

export function Orders() {
  const { user } = useContext(UserContext);

  const [searchParams] = useSearchParams();

  const pageIndex = z.coerce.number().parse(searchParams.get("page") ?? "1");

  const { data: result, isLoading: isLoadingOrders } = useQuery({
    queryKey: ["orders", pageIndex],
    queryFn: () => getOrders({ pageIndex, limit: ORDERS_LIMIT_PAGE }),
  });

  const isAdmin = user?.role === Role.admin;

  return (
    <>
      <Helmet title="Pedidos" />
      <div className="max-w-screen-2xl mx-auto w-full pb-10 -mt-24">
        <Card className="border-none drop-shadow-sm">
          <CardHeader className="gap-y-2 lg:flex-row lg:items-center lg:justify-between">
            <CardTitle className="text-xl line-clamp-1">Encomendas</CardTitle>
            {isAdmin && (
              <Button asChild size="sm">
                <Link to="/create-order">
                  <Plus className="size-4 mr-2" />
                  Criar encomenda
                </Link>
              </Button>
            )}
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
              {isLoadingOrders &&
                !result &&
                Array.from({ length: ORDERS_LIMIT_PAGE }).map((_, index) => (
                  <OrderCardSkeleton key={index} />
                ))}

              {result &&
                result.data.map((order) => {
                  return <OrderCard key={order.id} order={order} />;
                })}
            </div>
            {result && result.data.length === 0 && (
              <div className="flex flex-col items-center justify-center py-10">
                <img src={packageImg} className="w-16 h-16 mb-4" />
                <h3 className="text-lg font-semibold">
                  Nenhuma encomenda encontrada
                </h3>
                <p className="text-sm text-gray-500 mb-4">
                  Ainda não há pedidos. Que tal criar um agora?
                </p>
                <Button asChild size="sm">
                  <Link to="/create-order">
                    <Plus className="size-4 mr-2" />
                    Criar encomenda
                  </Link>
                </Button>
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </>
  );
}
