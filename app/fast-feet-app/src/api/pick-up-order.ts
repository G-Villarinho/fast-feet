import { api } from "@/lib/axios"

export interface PickUpOrderResponse {
    picknUpAt: string
}

export interface PickUpOrderParam {
    orderId: string
}

export async function PickUpOrder({orderId}: PickUpOrderParam) {
    const response = await api.patch<PickUpOrderResponse>(`/orders/${orderId}/status/pick-up`)

    return response.data;
}