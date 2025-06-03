import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Mitarbeiter as MitarbeiterAPI } from "@api/db";
import useSession from "@hooks/useSession";
import type { Mitarbeiter } from "bindings/viktor/db/models";
import { useEffect, useState } from "react";
import EinkaufForm from "./_components/einkauf-form";

export default function Eingabe() {
  const session = useSession();
  const [mitarbeiter, setMitarbeiter] = useState<Mitarbeiter[] | undefined>(
    undefined
  );
  const [Aus, setAus] = useState<Mitarbeiter | undefined>(undefined);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    async function x() {
      if (session == null) return;
      setLoading(true);
      const mas: Mitarbeiter[] = [];
      const res = await MitarbeiterAPI.GetAllEinkauf();
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
  }, [session]);

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
            defaultValue={Aus?.Name}
          >
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Auswahl..." />
            </SelectTrigger>
            <SelectContent>
              {mitarbeiter?.map((x) => (
                <SelectItem value={x.Name} key={x.Id}>
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
