import { Button } from "@/components/ui/button";
import { OrderStatus } from "@/api/get-orders";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";

import packageImg from "@/assets/package.svg";
import { OrderTimeline } from "./order-timeline";
import { ChevronRight } from "lucide-react";

interface OrderCardProps {
  order: {
    id: string;
    title: string;
    createdAt: string;
    status: OrderStatus;
  };
}

export function OrderCard({ order }: OrderCardProps) {
  const formattedDate = format(new Date(order.createdAt), "dd/MM/yyyy", {
    locale: ptBR,
  });

  return (
    <Card className="max-w-lg flex flex-col">
      <CardHeader>
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <img src={packageImg} />
            {/* Apliquei a classe truncate no título */}
            <span className="truncate max-w-[200px]">{order.title}</span>
          </div>
          <span className="text-xs text-gray-500">{formattedDate}</span>
        </div>
      </CardHeader>
      <CardContent>
        <OrderTimeline status={order.status} />
      </CardContent>
      <CardFooter className="p-0">
        <Button className="w-full py-3 rounded-t-none">
          <div className="w-full flex items-center justify-between">
            Detalhes
            <ChevronRight />
          </div>
        </Button>
      </CardFooter>
    </Card>
  );
}
