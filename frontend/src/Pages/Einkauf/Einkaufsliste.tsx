import { Mitarbeiter as MitarbeiterAPI } from "@api/db";
import type { Mitarbeiter } from "bindings/viktor/db/models";
import { useEffect, useState } from "react";
import { columns } from "./_components/columns";
import { DataTable } from "./_components/data-table";

export default function Einkaufsliste() {
  const [mitarbeiter, setMitarbeiter] = useState<Mitarbeiter[] | undefined>(
    undefined
  );
  const [loading, setLoadin] = useState(false);

  useEffect(() => {
    async function x() {
      setLoadin(true);
      const res = await MitarbeiterAPI.Einkauf();
      setMitarbeiter(res);
      setLoadin(false);
    }
    x();
  }, []);

  return (
    <div className="container mx-auto">
      {!loading && mitarbeiter && mitarbeiter.length > 0 && (
        <DataTable columns={columns} data={mitarbeiter} />
      )}
    </div>
  );
}
