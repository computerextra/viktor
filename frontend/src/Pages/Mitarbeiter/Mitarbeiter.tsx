import { Button } from "@/components/ui/button";
import { Suspense } from "react";
import { Link } from "react-router";
import MitarbeiterTable from "./_components/mitarbeiter-table";

export default function Mitarbeiter() {
  return (
    <>
      {/* TODO: BTN nur anzeigen, wenn Admin */}
      <Button asChild className="ms-2 mb-2" variant={"outline"}>
        <Link to="/Mitarbeiter/Neu">Neuen Mitarbeiter anlegen</Link>
      </Button>
      <h1 className="text-center">Mitarbeiter Übersicht</h1>
      <Suspense fallback={<>Loading....</>}>
        <MitarbeiterTable />
      </Suspense>
    </>
  );
}
