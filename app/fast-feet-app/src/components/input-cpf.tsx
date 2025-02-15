import * as React from "react";
import { cn } from "@/lib/utils";
import { InputMask } from "@react-input/mask";

interface InputCPFProps extends React.ComponentProps<"input"> {
  mask?: string;
}

const InputCPF = React.forwardRef<HTMLInputElement, InputCPFProps>(
  ({ className, type = "text", mask = "___.___.___-__", ...props }, ref) => {
    return (
      <InputMask
        mask={mask}
        type={type}
        className={cn(
          "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          className
        )}
        replacement={{ _: /\d/ }}
        ref={ref}
        {...props}
      />
    );
  }
);

InputCPF.displayName = "InputCPF";

export { InputCPF };
