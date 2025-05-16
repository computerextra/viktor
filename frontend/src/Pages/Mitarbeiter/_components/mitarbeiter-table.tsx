import { Read, type MitarbeiterModel } from "@api/db";
import { useEffect, useState } from "react";
import { columns } from "./columns";
import { DataTable } from "./data-table";

export default function MitarbeiterTable() {
  const [Mitarbeiter, setMitarbeiter] = useState<
    Array<MitarbeiterModel> | undefined
  >(undefined);

  useEffect(() => {
    async function x() {
      const mitarbeiter = (await Read("Mitarbeiter")) as MitarbeiterModel[];
      setMitarbeiter(mitarbeiter);
    }

    x();
  }, []);

  return (
    <div className="container mx-auto py-10">
      {Mitarbeiter && <DataTable columns={columns} data={Mitarbeiter} />}
    </div>
  );
}
