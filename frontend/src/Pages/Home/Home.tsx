import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Mitarbeiter } from "@api/db";
import type { db } from "@wails/go/models";
import { AlertCircle } from "lucide-react";
import { useEffect, useState } from "react";
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
        res.Vergangen.forEach((e) => a.push(e));
      }
      if (res.Heute?.length > 0) {
        res.Heute.forEach((e) => a.push(e));
      }
      if (res.Zukunft?.length > 0) {
        res.Zukunft.forEach((e) => a.push(e));
      }
      setListe(res);
      setalles(a);
    }

    x();
  }, []);

  return (
    <>
      <h1 className="text-center">Einkauf</h1>
      <p>Somthing somehitng bla bla bla</p>
      <h1 className="text-center">Geburtstagsliste</h1>
      <div className="container mx-auto">
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
          <>
            <h2>Geburtstage</h2>
            <DataTable columns={columns} data={alles} />
          </>
        )}
      </div>
    </>
  );
}
