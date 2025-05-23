import { Lieferant } from "@api/db";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { columns } from "./columns";
import { DataTable } from "./data-table";

export default function LieferantenTable() {
  const [Lieferanten, setLieferanten] = useState<
    Array<db.Lieferant> | undefined
  >(undefined);

  useEffect(() => {
    (async () => {
      const res = await Lieferant.GetAll();
      setLieferanten(res);
    })();
  }, []);

  return <DataTable columns={columns} data={Lieferanten ? Lieferanten : []} />;
}
