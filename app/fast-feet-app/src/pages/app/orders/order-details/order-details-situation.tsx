import { GetOrderResponse } from "@/api/get-order";
import { OrderStatus } from "@/api/get-orders";
import { Role } from "@/api/get-user";
import { PickUpOrder } from "@/api/pick-up-order";
import { Button } from "@/components/ui/button";
import { UserContext } from "@/contexts/user/user-context";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { useContext } from "react";
import { Link } from "react-router-dom";
import { toast } from "sonner";
import { Loader2 } from "lucide-react";

interface OrderDetailsSituationProps {
  orderId: string;
  status: OrderStatus;
  createdAt: string;
  picknUpAt?: string | null;
  deliveryAt?: string | null;
}

const orderStatusMap = new Map<OrderStatus, string>([
  [OrderStatus.waiting, "Aguardando"],
  [OrderStatus.picknUp, "Retirado"],
  [OrderStatus.done, "Entregue"],
]);

export function OrderDetailsSituation({
  orderId,
  status,
  createdAt,
  deliveryAt,
  picknUpAt,
}: OrderDetailsSituationProps) {
  const { user } = useContext(UserContext);
  const queryClient = useQueryClient();

  function updateOrderOnCache(picknUpAt: string) {
    const orderCache = queryClient.getQueryData<GetOrderResponse>([
      "order",
      orderId,
    ]);

    if (orderCache) {
      queryClient.setQueryData<GetOrderResponse>(["order", orderId], {
        ...orderCache,
        status: OrderStatus.picknUp,
        picknUpAt: picknUpAt,
      });
    }
  }

  const { mutateAsync: pickUpOrderFn, isPending } = useMutation({
    mutationFn: PickUpOrder,
  });

  async function handlePickUpOrder() {
    try {
      const response = await pickUpOrderFn({ orderId });
      updateOrderOnCache(response.picknUpAt);
      toast.success("Encomenda retirada com sucesso.");
    } catch (error) {
      if (isAxiosError(error)) {
        const errorMessage =
          error.response?.data?.details ||
          "Ocorreu um erro inesperado ao tentar realizar a retirada do pacote. Por favor, tente novamente.";

        toast.warning(errorMessage);
        return;
      }

      toast.error(
        "Ocorreu um erro inesperado ao tentar realizar o login. Por favor, tente novamente."
      );
    }
  }

  const isDeliveryMan = user?.role === Role.deliveryMan;

  return (
    user && (
      <div className="grid grid-cols-2 gap-6 md:px-12">
        <div className="flex flex-col gap-2">
          <h3 className="text-sm font-medium text-gray-700 uppercase">
            Status
          </h3>
          <p className="text-sm text-gray-500">{orderStatusMap.get(status)}</p>
        </div>

        <div className="flex flex-col gap-2">
          <h3 className="text-sm font-medium text-gray-700 uppercase">
            Postado em
          </h3>
          <p className="text-sm text-gray-500">
            {format(new Date(createdAt), "dd/MM/yyyy", {
              locale: ptBR,
            })}
          </p>
        </div>

        <div className="flex flex-col gap-2">
          <h3 className="text-sm font-medium text-gray-700 uppercase">
            Data de retirada
          </h3>
          <p className="text-sm text-gray-500">
            {picknUpAt
              ? format(new Date(picknUpAt), "dd/MM/yyyy", {
                  locale: ptBR,
                })
              : "--/--/----"}
          </p>
        </div>

        <div className="flex flex-col gap-2">
          <h3 className="text-sm font-medium text-gray-700 uppercase">
            Data de entrega
          </h3>
          <p className="text-sm text-gray-500">
            {deliveryAt
              ? format(new Date(deliveryAt), "dd/MM/yyyy", {
                  locale: ptBR,
                })
              : "--/--/----"}
          </p>
        </div>

        <div className="col-span-2 flex justify-center mt-4">
          {status === OrderStatus.waiting && isDeliveryMan && (
            <Button
              className="w-full"
              onClick={handlePickUpOrder}
              disabled={isPending}
            >
              {isPending ? (
                <Loader2 className="size-4 animate-spin" />
              ) : (
                "Retirar pacote"
              )}
            </Button>
          )}

          {status === OrderStatus.picknUp && isDeliveryMan && (
            <Button className="w-full">
              <Link to={`/orders/${orderId}/delivery`}>Entregar</Link>
            </Button>
          )}
        </div>
      </div>
    )
  );
}
