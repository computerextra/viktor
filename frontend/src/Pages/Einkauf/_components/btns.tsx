import { GetMitarbeiters, type MitarbeiterRes } from "@/api/Mitarbeiter";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useEffect, useState } from "react";
import { Link } from "react-router";

export default function EinkaufKnöppe() {
  const [Mitarbeiter, setMitarbeiter] = useState<MitarbeiterRes[] | undefined>(
    undefined
  );
  const [loading, setLoading] = useState(false);
  const [auswahl, setAuswahl] = useState<string | undefined>(undefined);

  useEffect(() => {
    (async () => {
      setLoading(true);
      const res = await GetMitarbeiters();
      setMitarbeiter(res);
      setLoading(false);
    })();
  }, []);

  if (loading) return <>Lädt...</>;

  return (
    <div className="mb-5 grid grid-cols-3 gap-8">
      <AlertDialog>
        <AlertDialogTrigger asChild>
          <Button>Eingabe</Button>
        </AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Auswahl</AlertDialogTitle>
            <AlertDialogDescription>
              <Select onValueChange={(e) => setAuswahl(e)}>
                <SelectTrigger>
                  <SelectValue placeholder="Mitarbeiter" />
                </SelectTrigger>
                <SelectContent>
                  {Mitarbeiter?.map((x) => (
                    <SelectItem key={x.id} value={x.id}>
                      {x.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Abbrechen</AlertDialogCancel>
            <AlertDialogAction asChild>
              <Link to={"/Einkauf/" + auswahl}>
                {auswahl == null || auswahl.length < 1
                  ? "Niemand ausgewählt"
                  : "Eingeben"}
              </Link>
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
      <Button asChild>
        <a
          href="https://www.edeka.de/markt-id/10001842/prospekt/"
          target="_blank"
          rel="noopener noreferrer"
        >
          Edeka Blättchen
        </a>
      </Button>
      <Button onClick={() => window.print()}>Drucken</Button>
    </div>
  );
}
