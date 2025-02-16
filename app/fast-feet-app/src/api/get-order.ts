import { api } from "@/lib/axios"
import { OrderStatus } from "./get-orders"

export interface GetOrderResponse {
    id: string
    status: OrderStatus
    recipientName: string
    recipientAddress: string
    recipientZipcode: string
    createdAt: string
    picknUpAt?: string | null
    deliveryAt?: string | null 
}

export interface GetOrderParam {
    orderId?: string | null
}

export async function getOrder({ orderId }: GetOrderParam) {
    const response = await api.get<GetOrderResponse>(`/orders/${orderId}`)

    return response.data
}