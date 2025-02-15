import { api } from "@/lib/axios";

export enum Role {
  owner = "OWNER",
  admin = "ADMIN",
  deliveryMan = "DELIVERY_MAN",
}

export interface GetUserResponse {
  id: string;
  fullName: string;
  email: string;
  role: Role;
}

export async function getUser() {
  const response = await api.get<GetUserResponse>("/users/me");

  return response.data;
}
