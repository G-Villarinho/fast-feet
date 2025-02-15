import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { createRecipient } from "@/api/create-recipient";
import { useMutation } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { useState } from "react";
import { Loader2 } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Link } from "react-router-dom";

const createRecipientSchema = z.object({
  fullName: z.string().nonempty("O nome completo é obrigatório."),
  email: z.string().email("Digite um e-mail válido."),
  state: z.string().nonempty("O estado é obrigatório."),
  city: z.string().nonempty("A cidade é obrigatória."),
  neighborhood: z.string().nonempty("O bairro é obrigatório."),
  address: z.string().nonempty("O endereço é obrigatório."),
  zipcode: z
    .string()
    .min(8, "O CEP deve ter pelo menos 8 dígitos.")
    .max(8, "O CEP deve ter no máximo 8 dígitos.")
    .regex(/^\d+$/, "O CEP deve conter apenas números."),
});

type CreateRecipientSchema = z.infer<typeof createRecipientSchema>;

export function CreateRecipientForm() {
  const [message, setMessage] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<CreateRecipientSchema>({
    resolver: zodResolver(createRecipientSchema),
  });

  const {
    mutateAsync: createRecipientFn,
    isPending,
    isSuccess,
  } = useMutation({
    mutationFn: createRecipient,
  });

  async function handleCreateRecipient(data: CreateRecipientSchema) {
    try {
      await createRecipientFn({
        fullName: data.fullName,
        address: data.address,
        city: data.city,
        email: data.email,
        neighborhood: data.neighborhood,
        state: data.state,
        zipcode: Number(data.zipcode),
      });
    } catch (err) {
      if (isAxiosError(err)) {
        const errorMessage =
          err.response?.data?.details ||
          "Ocorreu um erro inesperado ao tentar realizar o login. Por favor, tente novamente.";
        setMessage(errorMessage);
      }
    }
  }

  return (
    <form
      onSubmit={handleSubmit(handleCreateRecipient)}
      className="space-y-4 max-w-md mx-auto"
    >
      {isSuccess === false && message && (
        <Alert variant="destructive" className="mb-8">
          <AlertTitle></AlertTitle>
          <AlertDescription>
            <p>{message}</p>
          </AlertDescription>
        </Alert>
      )}

      {isSuccess === true && (
        <Alert variant="success" className="mb-8">
          <AlertTitle>Destinatário criado com sucesso!</AlertTitle>
          <AlertDescription>
            <p>Deseja criar uma encomenda para esse destinatário agora?</p>
            <Button className="mt-2" asChild>
              <Link to="/orders">Criar encomenda</Link>
            </Button>
          </AlertDescription>
        </Alert>
      )}

      <div>
        <label className="block text-sm font-medium mb-1">Nome Completo</label>
        <Input placeholder="Digite o nome completo" {...register("fullName")} />
        {errors.fullName && (
          <p className="text-red-500 text-sm">{errors.fullName.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium mb-1">E-mail</label>
        <Input
          type="email"
          placeholder="Digite o e-mail"
          {...register("email")}
        />
        {errors.email && (
          <p className="text-red-500 text-sm">{errors.email.message}</p>
        )}
      </div>

      <div className="flex gap-4">
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1">Estado</label>
          <Input placeholder="Ex: SP" {...register("state")} />
          {errors.state && (
            <p className="text-red-500 text-sm">{errors.state.message}</p>
          )}
        </div>

        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1">Cidade</label>
          <Input placeholder="Digite a cidade" {...register("city")} />
          {errors.city && (
            <p className="text-red-500 text-sm">{errors.city.message}</p>
          )}
        </div>
      </div>

      <div>
        <label className="block text-sm font-medium mb-1">Bairro</label>
        <Input placeholder="Digite o bairro" {...register("neighborhood")} />
        {errors.neighborhood && (
          <p className="text-red-500 text-sm">{errors.neighborhood.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium mb-1">Endereço</label>
        <Input placeholder="Digite o endereço" {...register("address")} />
        {errors.address && (
          <p className="text-red-500 text-sm">{errors.address.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium mb-1">CEP</label>
        <Input placeholder="Digite o CEP" {...register("zipcode")} />
        {errors.zipcode && (
          <p className="text-red-500 text-sm">{errors.zipcode.message}</p>
        )}
      </div>

      <Button
        type="submit"
        size="lg"
        disabled={isPending}
        className="w-full text-sm py-3"
      >
        {isPending ? (
          <Loader2 className="size-4 animate-spin" />
        ) : (
          "Acessar painel"
        )}
      </Button>
    </form>
  );
}
