import type { ColumnDef } from "@tanstack/react-table";
import type { db } from "@wails/go/models";

export const columns: ColumnDef<db.Mitarbeiter>[] = [
  {
    accessorKey: "Name",
    header: "Name",
  },
  {
    accessorKey: "Geburtstag",
    header: "Geburtstag",
    cell: ({ row }) => {
      const x = row.original;

      if (x.Geburtstag.Valid) {
        return new Date(x.Geburtstag.Time).toLocaleDateString("de-DE", {
          day: "2-digit",
          month: "long",
          year: "numeric",
        });
      }
    },
  },
  {
    id: "Geburtstag_Diff",
    header: "",
    cell: ({ row }) => {
      const x = row.original;
      if (x.Geburtstag.Valid) {
        const diff =
          (new Date().getTime() - new Date(x.Geburtstag.Time).getTime()) /
          (1000 * 60 * 60 * 24);
        if (diff < 0) {
          return `In ${Math.ceil(diff * -1)} Tagen`;
        } else if (diff > 1) {
          return `Vor ${Math.ceil(diff)} Tagen`;
        } else {
          return "Heute";
        }
      }
    },
  },
];
