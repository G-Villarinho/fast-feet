import { api } from "@/lib/axios";

export interface CreateOrderRequest {
  recipientId: string;
  title: string;
}

export async function createOrder({ recipientId, title }: CreateOrderRequest) {
  await api.post("/orders", { recipientId, title });
}
