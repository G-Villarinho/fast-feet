import { Skeleton } from "@/components/ui/skeleton";
import {
  Card,
  CardHeader,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ChevronRight } from "lucide-react";

export function OrderCardSkeleton() {
  return (
    <Card className="max-w-lg flex flex-col">
      <CardHeader>
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <Skeleton className="w-16 h-16 rounded-full" />
            <Skeleton className="w-32 h-4" />
          </div>
          <Skeleton className="w-20 h-4" />
        </div>
      </CardHeader>

      <CardContent>
        <Skeleton className="h-4 w-32" />
      </CardContent>

      <CardFooter className="p-0">
        <Button className="w-full py-3 rounded-t-none">
          <div className="w-full flex items-center justify-between">
            <Skeleton className="w-24 h-4" />
            <ChevronRight className="invisible" />{" "}
          </div>
        </Button>
      </CardFooter>
    </Card>
  );
}
