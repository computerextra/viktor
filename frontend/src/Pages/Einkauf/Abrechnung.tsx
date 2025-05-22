import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Mitarbeiter } from "@api/db";
import { Paypal } from "@wails/go/main/App";
import { db } from "@wails/go/models";
import { useEffect, useState } from "react";

export default function Abrechnung() {
  const [mitarbeiter, setMitarbeiter] = useState<db.Mitarbeiter[] | undefined>(
    undefined
  );
  const [username, setUsername] = useState<string | undefined>(undefined);
  const [Betrag, setBetrag] = useState<string | undefined>(undefined);
  const [loading, setLoading] = useState(false);
  const [selected, setSelected] = useState<db.Mitarbeiter | undefined>(
    undefined
  );

  useEffect(() => {
    (async () => {
      setLoading(true);
      const mas: db.Mitarbeiter[] = [];
      const res = await Mitarbeiter.Einkauf();
      res.forEach((x) => {
        if (x.Email != null) {
          mas.push(x);
        }
      });
      setMitarbeiter(res);
      setLoading(false);
    })();
  }, []);

  const handleSubmit = async () => {
    if (username == null) return;
    if (Betrag == null) return;
    if (loading) return;
    if (selected == null) return;

    // TODO: Resettet nicht das Formular... Mach ma neu!
    setLoading(true);
    const res = await Paypal(username, Betrag, selected.ID);
    if (res) {
      setBetrag(undefined);
      setSelected(undefined);
    } else {
      alert("Fehler beim Senden der E-Mail");
    }
    setLoading(false);
  };

  return (
    <div className="flex flex-col items-center">
      <h1 className="text-center">PayPal Abrechnung</h1>
      <form onSubmit={(e) => e.preventDefault()} className="space-y-4 mt-12">
        <Select
          onValueChange={(e) => {
            const m = mitarbeiter?.find((x) => x.Email == e);
            setSelected(m);
          }}
          disabled={loading}
        >
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Mitarbeiter" />
          </SelectTrigger>
          <SelectContent>
            {mitarbeiter?.map((x) => {
              if (x.Email && x.Email.length > 0)
                return (
                  <SelectItem key={x.ID} value={x.Email}>
                    {x.Name}
                  </SelectItem>
                );
            })}
          </SelectContent>
        </Select>
        <div className="grid w-full max-w-sm items-center gap-1.5">
          <Label htmlFor="username">Paypal Benutzername</Label>
          <Input
            disabled={loading}
            type="text"
            id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className="grid w-full max-w-sm items-center gap-1.5">
          <Label htmlFor="betrag">Betrag in €</Label>
          <Input
            disabled={loading}
            type="text"
            id="betrag"
            value={Betrag}
            onChange={(e) => setBetrag(e.target.value)}
          />
        </div>

        <Button disabled={loading} onClick={handleSubmit}>
          {loading ? "Sendet..." : "Mail Senden"}
        </Button>
      </form>
    </div>
  );
}
