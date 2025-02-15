import { Button } from "@/components/ui/button";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { Plus } from "lucide-react";
import { Helmet } from "react-helmet-async";
import { Link } from "react-router-dom";

export function Recipients() {
  return (
    <>
      <Helmet title="Destinatários" />
      <div className="max-w-screen-2xl mx-auto w-full pb-10 -mt-24">
        <Card className="border-none drop-shadow-sm">
          <CardHeader className="gap-y-2 lg:flex-row lg:items-center lg:justify-between">
            <CardTitle className="text-xl line-clamp-1">
              Destinatários
            </CardTitle>
            <Button asChild size="sm">
              <Link to="/create-recipient">
                <Plus className="size-4 mr-2" />
                Criar destinatário
              </Link>
            </Button>
          </CardHeader>
        </Card>
      </div>
    </>
  );
}
