import { Mitarbeiter } from "@api/db";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { columns } from "./_components/columns";
import { DataTable } from "./_components/data-table";

export default function Einkaufsliste() {
  const [mitarbeiter, setMitarbeiter] = useState<db.Mitarbeiter[] | undefined>(
    undefined
  );
  const [loading, setLoadin] = useState(false);

  useEffect(() => {
    async function x() {
      setLoadin(true);
      const res = await Mitarbeiter.Einkauf();
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
