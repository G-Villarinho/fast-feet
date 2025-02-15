import { OrderStatus } from "@/api/get-orders";
import { Progress } from "@/components/ui/progress";
import { cn } from "@/lib/utils";
import { CheckCircle, Clock, PackageCheck } from "lucide-react";

const orderStatusProgress = new Map<OrderStatus, number>([
  [OrderStatus.waiting, 20],
  [OrderStatus.picknUp, 50],
  [OrderStatus.done, 100],
]);

interface OrderTimelineProps {
  status: OrderStatus;
}

export function OrderTimeline({ status }: OrderTimelineProps) {
  return (
    <div className="flex flex-col gap-2">
      <Progress value={orderStatusProgress.get(status)} />
      <div className="flex items-center justify-between text-xs font-medium text-primary/50">
        <span className="flex items-center gap-1 text-green-500">
          <Clock className="size-4" /> Aguardando
        </span>
        <span
          className={cn(
            "flex items-center gap-1",
            status === OrderStatus.picknUp || status === OrderStatus.done
              ? "text-green-500"
              : "text-gray-400"
          )}
        >
          <PackageCheck className="size-4" /> Retirado
        </span>
        <span
          className={cn(
            "flex items-center gap-1",
            status === OrderStatus.done ? "text-green-500" : "text-gray-400"
          )}
        >
          <CheckCircle className="size-4" /> Entregue
        </span>
      </div>
    </div>
  );
}
