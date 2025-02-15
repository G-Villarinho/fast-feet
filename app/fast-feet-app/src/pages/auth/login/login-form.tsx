import google from "@/assets/google.svg";

import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { isValidCPF } from "@/utils/cpf";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { InputCPF } from "@/components/input-cpf";
import { Link, useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import { login } from "@/api/login";
import { useState } from "react";
import { isAxiosError } from "axios";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Loader2 } from "lucide-react";

const loginSchema = z.object({
  cpf: z.string().refine(isValidCPF, "CPF inválido"),
  password: z.string().min(8, "A senha deve ter 8 dígitos"),
});

type LoginSchema = z.infer<typeof loginSchema>;

export function LoginForm({
  className,
  ...props
}: React.ComponentPropsWithoutRef<"div">) {
  const navigate = useNavigate();
  const [message, setMessage] = useState<string | null>(null);
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginSchema>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      cpf: "",
      password: "",
    },
  });

  const { mutateAsync: loginFn, isSuccess } = useMutation({
    mutationFn: login,
  });

  async function handleLogin({ cpf, password }: LoginSchema) {
    try {
      await loginFn({ cpf, password });
      navigate("/orders");
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
    <div className={cn("flex flex-col gap-8", className)} {...props}>
      <Card className="border-none shadow-xl rounded-3xl p-2 md:p-4 w-full max-w-md mx-auto">
        <CardHeader className="text-center mb-6">
          <CardTitle className="text-2xl">Bem vindo(a) de volta</CardTitle>
          <CardDescription>
            Faça seu login para começar suas entregas
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(handleLogin)}>
            {isSuccess === false && message && (
              <Alert variant="destructive" className="mb-8">
                <AlertTitle></AlertTitle>
                <AlertDescription>
                  <p>{message}</p>
                </AlertDescription>
              </Alert>
            )}
            <div className="grid gap-8">
              <div className="grid gap-5">
                <div className="grid gap-2">
                  <Label htmlFor="cpf">CPF</Label>
                  <InputCPF
                    id="cpf"
                    type="text"
                    autoFocus
                    className="h-12"
                    placeholder="Ex: 999.999.999-99"
                    {...register("cpf")}
                  />
                  {errors.cpf && (
                    <p className="text-red-500 text-sm">{errors.cpf.message}</p>
                  )}
                </div>
                <div className="grid gap-2">
                  <Label htmlFor="password">Senha</Label>
                  <Input
                    id="password"
                    type="password"
                    placeholder="Informe sua senha"
                    className="h-12"
                    {...register("password")}
                  />
                  {errors.password && (
                    <p className="text-red-500 text-sm">
                      {errors.password.message}
                    </p>
                  )}
                </div>
                <div className="flex text-sm font-medium self-start my-1">
                  <Link to="/forget-password" className="hover:underline">
                    {" "}
                    Esqueceu sua senha?
                  </Link>
                </div>
                <Button
                  type="submit"
                  size="lg"
                  disabled={isSubmitting}
                  className="w-full text-sm py-3"
                >
                  {isSubmitting ? (
                    <Loader2 className="size-4 animate-spin" />
                  ) : (
                    "Acessar painel"
                  )}
                </Button>
              </div>
              <div className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
                <span className="relative z-10 bg-background px-2 text-muted-foreground">
                  Ou continue com
                </span>
              </div>
              <Button
                variant="outline"
                className="w-full flex items-center gap-2"
                disabled={isSubmitting}
                size="lg"
              >
                <img src={google} className="w-5 h-5" />
                <span className="font-bold text-black">Entrar com Google</span>
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
      <div className="text-balance text-center text-sm text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 [&_a]:hover:text-primary  ">
        Ao clicar em continuar, você concorda com os{" "}
        <a href="#">Termos de serviço</a> e{" "}
        <a href="#">Políticas de Privacidade</a>.
      </div>
    </div>
  );
}
