import { Lieferant as LieferantAPI } from "@api/db";
import type { Lieferant } from "@bindings/viktor/db/models";
import { useEffect, useState } from "react";
import { columns } from "./columns";
import { DataTable } from "./data-table";

export default function LieferantenTable() {
  const [Lieferanten, setLieferanten] = useState<Array<Lieferant> | undefined>(
    undefined
  );

  useEffect(() => {
    (async () => {
      const res = await LieferantAPI.GetAll();
      setLieferanten(res);
    })();
  }, []);

  return <DataTable columns={columns} data={Lieferanten ? Lieferanten : []} />;
}
