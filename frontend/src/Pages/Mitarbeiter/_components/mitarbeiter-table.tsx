import { Mitarbeiter as MitarbeiterAPI } from "@api/db";
import type { Mitarbeiter } from "@bindings/viktor/db/models";
import { useEffect, useState } from "react";
import { columns } from "./columns";
import { DataTable } from "./data-table";

export default function MitarbeiterTable() {
  const [Worker, setMitarbeiter] = useState<Array<Mitarbeiter> | null>(null);

  useEffect(() => {
    async function x() {
      const mitarbeiter = await MitarbeiterAPI.GetAll();
      setMitarbeiter(mitarbeiter);
    }

    x();
  }, []);

  return (
    <div className="container mx-auto py-10">
      {Worker && <DataTable columns={columns} data={Worker} />}
    </div>
  );
}
