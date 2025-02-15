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
import { Progress } from "@/components/ui/progress";
import { cn } from "@/lib/utils";
import { Clock, PackageCheck, CheckCircle, ChevronRight } from "lucide-react";

interface OrderCardProps {
  order: {
    id: string;
    title: string;
    createdAt: string;
    status: OrderStatus;
  };
}

const orderStatusProgress = new Map<OrderStatus, number>([
  [OrderStatus.waiting, 20],
  [OrderStatus.picknUp, 50],
  [OrderStatus.done, 100],
]);

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
            <span>{order.title}</span>
          </div>
          <span className="text-xs text-gray-500">{formattedDate}</span>
        </div>
      </CardHeader>
      <CardContent>
        <div className="flex flex-col gap-2">
          <Progress value={orderStatusProgress.get(order.status)} />
          <div className="flex items-center justify-between text-xs font-medium text-primary/50">
            <span className="flex items-center gap-1 text-green-500">
              <Clock className="size-4" /> Aguardando
            </span>
            <span
              className={cn(
                "flex items-center gap-1",
                order.status === OrderStatus.picknUp ||
                  order.status === OrderStatus.done
                  ? "text-green-500"
                  : "text-gray-400"
              )}
            >
              <PackageCheck className="size-4" /> Retirado
            </span>
            <span
              className={cn(
                "flex items-center gap-1",
                order.status === OrderStatus.done
                  ? "text-green-500"
                  : "text-gray-400"
              )}
            >
              <CheckCircle className="size-4" /> Entregue
            </span>
          </div>
        </div>
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
