import React, { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { ContactRound } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
import { GetRecipientsBasicInfo } from "@/api/get-recipient-basic-info";
import { useDebounce } from "@/hooks/use-debounce";

interface RecipientSelectProps {
  value: string;
  onChange: (value: string) => void;
}

export function RecipientSelect({ value, onChange }: RecipientSelectProps) {
  const [search, setSearch] = useState("");
  const [pageIndex, setPageIndex] = useState(1);
  const debouncedSearch = useDebounce(search, 500);

  const {
    data: recipients,
    isLoading,
    isFetching,
  } = useQuery({
    queryKey: ["recipients", debouncedSearch, pageIndex],
    queryFn: async () => {
      const response = await GetRecipientsBasicInfo({
        pageIndex,
        q: debouncedSearch,
      });
      return response!;
    },
  });

  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
    setPageIndex(1);
  };

  const loadMoreRecipients = () => {
    if (!isFetching) {
      setPageIndex((prev) => prev + 1);
    }
  };

  return (
    <Select onValueChange={onChange} defaultValue={value}>
      <SelectTrigger className="w-full bg-zinc-100 h-11">
        <SelectValue placeholder="Selecione um destinatário" />
      </SelectTrigger>
      <SelectContent className="w-[338px]">
        <div className="p-2">
          <Input
            placeholder="Pesquisar destinatário..."
            value={search}
            onChange={handleSearchChange}
            className="mb-2"
          />
        </div>

        {isLoading && !recipients ? (
          <div className="flex justify-center py-4">
            <span className="animate-spin h-5 w-5 border-2 border-gray-400 border-t-transparent rounded-full"></span>
          </div>
        ) : recipients?.data && recipients.data.length > 0 ? (
          recipients.data.map((recipient) => (
            <SelectItem
              key={recipient.id}
              value={recipient.id}
              className="flex items-center gap-3 p-3"
            >
              <div className="flex items-center gap-2">
                <ContactRound className="h-5 w-5 text-gray-500" />
                <span className="font-medium text-gray-900 truncate max-w-[140px]">
                  {recipient.fullName}
                </span>
                <span className="text-xs text-gray-500 truncate max-w-[140px]">
                  ({recipient.email})
                </span>
              </div>
            </SelectItem>
          ))
        ) : (
          <div className="p-2 text-center">
            <p className="text-gray-600">Nenhum destinatário encontrado.</p>
            <Button
              type="button"
              className="mt-2 w-full"
              size="xs"
              variant="outline"
              asChild
            >
              <Link to="/recipients">Criar Destinatário</Link>
            </Button>
          </div>
        )}

        {recipients?.data &&
          recipients.data.length > 0 &&
          pageIndex < (recipients?.totalPages || 1) && (
            <Button
              type="button"
              onClick={loadMoreRecipients}
              className="w-full mt-2"
              size="xs"
              variant="outline"
              disabled={isFetching}
            >
              {isFetching ? "Carregando..." : "Carregar mais"}
            </Button>
          )}
      </SelectContent>
    </Select>
  );
}
