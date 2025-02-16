import { getOrder } from "@/api/get-order";
import { useQuery } from "@tanstack/react-query";
import { Helmet } from "react-helmet-async";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useParams } from "react-router-dom";
import { ClipboardList, Info } from "lucide-react";
import { OrderDetailsSituation } from "./order-details-situation";

export function OrderDetails() {
  const { orderId } = useParams();

  const { data: order } = useQuery({
    queryKey: ["order", orderId],
    queryFn: () => getOrder({ orderId }),
  });

  return (
    <>
      <Helmet title="Pedido detalhes" />
      <div className="max-w-screen-2xl mx-auto w-full pb-10 -mt-24">
        {order && (
          <div className="flex flex-col md:flex-row gap-4 p-2 md:p-0">
            <Card className="border-none drop-shadow-sm md:w-1/2">
              <CardHeader className="gap-y-2 lg:flex-row lg:items-center lg:justify-between">
                <CardTitle className="text-xl line-clamp-1">
                  <div className="flex items-center gap-2">
                    <Info size={30} fill="#facc15" className="text-white" />
                    Dados
                  </div>
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="flex flex-col md:flex-row gap-4 md:gap-32 text-sm uppercase px-2">
                  <div className="flex flex-col gap-1 ">
                    <h3 className="font-medium">Destinatário</h3>
                    <p className="text-gray-500">{order.recipientName}</p>
                  </div>
                  <div className="flex flex-col gap-1">
                    <h3 className="font-medium">Endereço</h3>
                    <div className="flex gap-2">
                      <p className=" text-gray-500">{order.recipientAddress}</p>
                      <p className=" text-gray-500">{order.recipientZipcode}</p>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>

            <Card className="border-none drop-shadow-sm md:w-1/2">
              <CardHeader className="gap-y-2 lg:flex-row lg:items-center lg:justify-between">
                <CardTitle className="text-xl line-clamp-1">
                  <div className="flex items-center gap-2">
                    <ClipboardList
                      size={30}
                      fill="#facc15"
                      className="text-white"
                    />
                    Situação
                  </div>
                </CardTitle>
              </CardHeader>
              <CardContent>
                <OrderDetailsSituation
                  orderId={order.id}
                  status={order.status}
                  createdAt={order.createdAt}
                  deliveryAt={order.deliveryAt}
                  picknUpAt={order.picknUpAt}
                />
              </CardContent>
            </Card>
          </div>
        )}
      </div>
    </>
  );
}
