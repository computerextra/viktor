import { Button } from "@/components/ui/button";
import { Suspense } from "react";
import { Link } from "react-router";
import LieferantenTable from "./_components/lieferanten-table";

export default function LieferantOverview() {
  return (
    <>
      <Button asChild className="ms-2 mb-2" variant={"outline"}>
        <Link to="/Lieferant/Neu">Neuen Lieferanten anlegen</Link>
      </Button>
      <h1 className="text-center mb-5">Lieferanten</h1>
      <Suspense fallback={<>Loading....</>}>
        <LieferantenTable />
      </Suspense>
    </>
  );
}
