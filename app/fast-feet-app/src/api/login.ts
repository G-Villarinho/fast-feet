import { api } from "@/lib/axios";

export interface LoginRequest {
  cpf: string;
  password: string;
}

export async function login({ cpf, password }: LoginRequest) {
  await api.post("/login", { cpf, password });
}
