import { Mitarbeiter } from "@api/db";
import { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { columns } from "./columns";
import { DataTable } from "./data-table";

export default function MitarbeiterTable() {
  const [Worker, setMitarbeiter] = useState<Array<db.Mitarbeiter> | undefined>(
    undefined
  );

  useEffect(() => {
    async function x() {
      const mitarbeiter = await Mitarbeiter.GetAll();
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
