import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Mitarbeiter } from "@api/db";
import useSession from "@hooks/useSession";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import EinkaufForm from "./_components/einkauf-form";

export default function Eingabe() {
  const session = useSession();
  const [mitarbeiter, setMitarbeiter] = useState<db.Mitarbeiter[] | undefined>(
    undefined
  );
  const [Aus, setAus] = useState<db.Mitarbeiter | undefined>(undefined);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    // TODO: scheint nicht korrekt zu laufen
    async function x() {
      setLoading(true);
      const mas: db.Mitarbeiter[] = [];
      const res = await Mitarbeiter.GetAll();
      res.forEach((ma) => {
        if (ma.Name == "Kaffeekasse") {
          mas.push(ma);
        }
        if (session && session.Name) {
          if (session.Mail == ma.Email) {
            mas.push(ma);
            setAus(ma);
          }
        }
      });
      setMitarbeiter(mas);
      setLoading(false);
    }

    x();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <>
      <h1 className="text-center">Einkauf Eingabe</h1>
      {!loading && (
        <div className="container mx-auto mt-5">
          {/* Auswahl */}

          <Select
            onValueChange={(e) => {
              const z = mitarbeiter?.find((y) => y.Name === e);
              setAus(z);
            }}
          >
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Auswahl..." />
            </SelectTrigger>
            <SelectContent>
              {mitarbeiter?.map((x) => (
                <SelectItem value={x.Name} key={x.ID}>
                  {x.Name}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          {Aus && <EinkaufForm mitarbeiter={Aus} />}
        </div>
      )}
    </>
  );
}
