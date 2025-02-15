import { api } from "@/lib/axios";

export enum OrderStatus {
  waiting = "WAITING",
  picknUp = "PICKN_UP",
  done = "DONE",
}

export interface GetOrdersQuery {
  pageIndex?: number | null;
}

export interface GetOrdersResponse {
  data: {
    id: string;
    title: string;
    status: OrderStatus;
    createdAt: string;
  }[];
  total: number;
  totalPages: number;
  pageIndex: number;
  limit: number;
}

export async function GetOrders({ pageIndex }: GetOrdersQuery) {
  const response = await api.get<GetOrdersResponse>("/orders", {
    params: {
      pageIndex,
    },
  });

  return response.data;
}
