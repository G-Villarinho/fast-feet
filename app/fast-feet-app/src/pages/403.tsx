import forbidden from "@/assets/forbidden.svg";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";

export function Forbidden() {
  const navigate = useNavigate();

  return (
    <div className="flex flex-col items-center justify-center w-screen h-screen gap-12 py-8 ">
      <div className="flex flex-col items-center gap-4">
        <img src={forbidden} alt="Forbidden" className="animate-in" />
        <h1 className="text-3xl font-medium text-center">
          Você não está autorizado
        </h1>
        <p className="text-xl text-center ">
          Você tentou acessar uma página para a qual não tinha autorização
          prévia.
        </p>

        <Button
          onClick={() => navigate(-1)}
          variant="outline"
          className="mt-8 px-6 py-3 text-lg"
        >
          Voltar
        </Button>
      </div>
    </div>
  );
}
