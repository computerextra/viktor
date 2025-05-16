import type { MitarbeiterModel } from "@api/db";
import type { ColumnDef } from "@tanstack/react-table";

export const columns: ColumnDef<MitarbeiterModel>[] = [
  {
    accessorKey: "Name",
    header: "Name",
  },
  {
    accessorKey: "Email",
    header: "Email",
  },
  {
    accessorKey: "Intern_telefon1",
    header: "Interne Durchwahl",
  },
  {
    accessorKey: "Home_office",
    header: "Home Office",
  },
  {
    accessorKey: "Festnetz_busines",
    header: "Festnetz",
  },
  {
    accessorKey: "Mobil_buiness",
    header: "Mobil",
  },
];
