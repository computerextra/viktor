import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { Mitarbeiter } from "@api/db";
import type { db } from "@wails/go/models";
import { AlertCircle } from "lucide-react";
import { Suspense, useEffect, useState } from "react";
import { NavLink } from "react-router";
import Einkaufsliste from "../Einkauf/Einkaufsliste";
import { columns } from "./_components/columns";
import { DataTable } from "./_components/data-table";

export default function Home() {
  const [liste, setListe] = useState<db.Geburtstagsliste | undefined>(
    undefined
  );
  const [alles, setalles] = useState<db.Mitarbeiter[] | undefined>(undefined);

  useEffect(() => {
    async function x() {
      const res = await Mitarbeiter.Geburtstag();
      const a: db.Mitarbeiter[] = [];
      if (res.Vergangen?.length > 0) {
        res.Vergangen.forEach((e) => {
          if (e.Geburtstag.Valid && e.Geburtstag.Time.length > 0) {
            a.push(e);
          }
        });
      }
      if (res.Heute?.length > 0) {
        res.Heute.forEach((e) => {
          if (e.Geburtstag.Valid && e.Geburtstag.Time.length > 0) {
            a.push(e);
          }
        });
      }
      if (res.Zukunft?.length > 0) {
        res.Zukunft.forEach((e) => {
          if (e.Geburtstag.Valid && e.Geburtstag.Time.length > 0) {
            a.push(e);
          }
        });
      }
      setListe(res);
      setalles(a);
    }

    x();
  }, []);

  return (
    <>
      <h1 className="text-center print:hidden">Einkauf</h1>
      <h1 className="hidden print:block text-center my-2">
        An Post / Milch und Kaffee denken!
      </h1>
      <div className="container mx-auto grid grid-cols-4 my-5 gap-4 print:hidden">
        <Button asChild variant={"outline"}>
          <NavLink to="/Eingabe">Eingabe</NavLink>
        </Button>
        <Button variant={"outline"} onClick={window.print}>
          Liste Drucken
        </Button>
        <Button asChild variant={"outline"}>
          <a
            href="https://www.edeka.de/markt-id/10001842/prospekt/"
            target="_blank"
            rel="noopener noreferrer"
          >
            Edeka Blättchen
          </a>
        </Button>
        <Button asChild variant={"outline"}>
          <NavLink to="/Abrechnung">Paypal Abrechnung</NavLink>
        </Button>
      </div>
      <Suspense>
        <Einkaufsliste />
      </Suspense>

      <h1 className="text-center mt-5 print:hidden">Geburtstagsliste</h1>
      <div className="container mx-auto print:hidden">
        {liste?.Heute && (
          <div className="mt-8 mb-8">
            {liste.Heute.map((x) => (
              <Alert variant="destructive" key={x.ID + x.Name}>
                <AlertCircle className="h-4 w-4" />
                <AlertTitle>GEBURTSTAG</AlertTitle>
                <AlertDescription>
                  Heute hat {x.Name} Geburtstag
                </AlertDescription>
              </Alert>
            ))}
          </div>
        )}
        {alles && (
          <div className="print:hidden">
            <h2>Geburtstage</h2>
            <DataTable columns={columns} data={alles} />
          </div>
        )}
      </div>
    </>
  );
}
