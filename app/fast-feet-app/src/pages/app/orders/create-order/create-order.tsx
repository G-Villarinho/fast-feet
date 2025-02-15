import { InterceptedSheetContent } from "@/components/intercepted-sheet-content";
import { Sheet, SheetHeader, SheetTitle } from "@/components/ui/sheet";
import { CreateOrderForm } from "./create-order-form";

export function CreateOrder() {
  return (
    <Sheet defaultOpen>
      <InterceptedSheetContent className="bg-fast-feet-light-orange">
        <SheetHeader>
          <SheetTitle>Criar um encomenda</SheetTitle>
        </SheetHeader>

        <div className="pt-8">
          <CreateOrderForm />
        </div>
      </InterceptedSheetContent>
    </Sheet>
  );
}
