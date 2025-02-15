import { api } from "@/lib/axios";

export interface GetRecipientsBasicInfoQuery {
  pageIndex?: number | null;
  q?: string | null;
}

export interface GetRecipientsBasicInfoResponse {
  data: {
    id: string;
    fullName: string;
    email: string;
  }[];
  total: number;
  totalPages: number;
  pageIndex: number;
  limit: number;
}

export async function GetRecipientsBasicInfo({
  pageIndex,
  q,
}: GetRecipientsBasicInfoQuery) {
  const response = await api.get<GetRecipientsBasicInfoResponse>(
    "/recipients/lite",
    {
      params: {
        pageIndex,
        q,
      },
    }
  );

  return response.data;
}
