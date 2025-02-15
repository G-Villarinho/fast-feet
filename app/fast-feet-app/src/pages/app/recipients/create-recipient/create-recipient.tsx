import { InterceptedSheetContent } from "@/components/intercepted-sheet-content";
import { Sheet, SheetHeader, SheetTitle } from "@/components/ui/sheet";
import { CreateRecipientForm } from "./create-recipient-form";

export function CreateRecipient() {
  return (
    <Sheet defaultOpen>
      <InterceptedSheetContent className="bg-fast-feet-light-orange">
        <SheetHeader>
          <SheetTitle>Criar um destinatário</SheetTitle>
        </SheetHeader>

        <div className="pt-8">
          <CreateRecipientForm />
        </div>
      </InterceptedSheetContent>
    </Sheet>
  );
}
