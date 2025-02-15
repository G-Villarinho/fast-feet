import { api } from "@/lib/axios";

export interface CreateRecipientRequest {
  fullName: string;
  email: string;
  state: string;
  city: string;
  neighborhood: string;
  address: string;
  zipcode: number;
}

export async function createRecipient({
  fullName,
  email,
  state,
  city,
  neighborhood,
  address,
  zipcode,
}: CreateRecipientRequest) {
  await api.post("/recipients", {
    fullName,
    email,
    state,
    city,
    neighborhood,
    address,
    zipcode,
  });
}
