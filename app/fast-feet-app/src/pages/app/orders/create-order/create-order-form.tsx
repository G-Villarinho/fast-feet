import { useMutation } from "@tanstack/react-query";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Loader2 } from "lucide-react";
import { createOrder } from "@/api/create-order";
import { isAxiosError } from "axios";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { RecipientSelect } from "./recipient-select";

const createOrderSchema = z.object({
  title: z
    .string()
    .nonempty("O título da encomenda é obrigatório.")
    .max(255, "O título da encomenda não pode ultrapassar 255 caracteres."),
  recipientId: z.string().nonempty("O destinatário é obrigatório."),
});

type CreateOrderSchema = z.infer<typeof createOrderSchema>;

export function CreateOrderForm() {
  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<CreateOrderSchema>({
    resolver: zodResolver(createOrderSchema),
    defaultValues: {
      title: "",
      recipientId: "",
    },
  });

  const [message, setMessage] = useState<string | null>(null);

  const {
    mutateAsync: createOrderFn,
    isPending,
    isSuccess,
  } = useMutation({
    mutationFn: createOrder,
  });

  async function handleCreateOrder(data: CreateOrderSchema) {
    try {
      await createOrderFn({
        title: data.title,
        recipientId: data.recipientId,
      });
    } catch (err) {
      if (isAxiosError(err)) {
        const errorMessage =
          err.response?.data?.details ||
          "Ocorreu um erro inesperado. Por favor, tente novamente.";
        setMessage(errorMessage);
      }
    }
  }

  return (
    <form onSubmit={handleSubmit(handleCreateOrder)} className="space-y-6">
      {isSuccess === false && message && (
        <Alert variant="destructive" className="mb-8">
          <AlertDescription>
            <p>{message}</p>
          </AlertDescription>
        </Alert>
      )}

      {isSuccess === true && (
        <Alert variant="success" className="mb-8">
          <AlertDescription>Encomenda criada com sucesso</AlertDescription>
        </Alert>
      )}

      <div>
        <label htmlFor="title" className="block text-sm font-medium mb-1">
          Título da Encomenda
        </label>
        <Input
          id="title"
          placeholder="Digite o título"
          className="bg-zinc-100 h-11"
          {...register("title")}
        />
        {errors.title && (
          <p className="text-sm text-red-500 mt-1">{errors.title.message}</p>
        )}
      </div>

      <div className="flex flex-col">
        <label htmlFor="recipientId" className="block text-sm font-medium mb-1">
          Destinatário
        </label>
        <Controller
          name="recipientId"
          control={control}
          render={({ field }) => (
            <RecipientSelect value={field.value} onChange={field.onChange} />
          )}
        />
        {errors.recipientId && (
          <p className="text-sm text-red-500 mt-1">
            {errors.recipientId.message}
          </p>
        )}
      </div>

      <Button type="submit" className="w-full" disabled={isPending}>
        {isPending ? (
          <Loader2 className="h-4 w-4 animate-spin" />
        ) : (
          "Criar encomenda"
        )}
      </Button>
    </form>
  );
}
