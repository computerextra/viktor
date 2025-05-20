import type { ColumnDef } from "@tanstack/react-table";
import type { db } from "@wails/go/models";
import { Check, Cross } from "lucide-react";

export const columns: ColumnDef<db.Mitarbeiter>[] = [
  {
    accessorKey: "Name",
    header: "Name",
  },
  {
    accessorKey: "Geld",
    header: "Geld / Pfand",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <div className="grid grid-cols-2">
          <span>Pfand</span>
          <span>{x.Pfand}</span>
          <span>Geld</span>
          <span>{x.Geld}</span>
        </div>
      );
    },
  },
  {
    accessorKey: "Paypal",
    header: "Paypal / Abo",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <div className="grid grid-cols-2">
          <span>Abo</span>
          {x.Abonniert ? (
            <Check className="text-green-500 h-4 w-4" />
          ) : (
            <Cross className="text-red-500 h-4 w-4 rotate-45" />
          )}
          <span>Paypal</span>
          {x.Paypal ? (
            <Check className="text-green-500 h-4 w-4" />
          ) : (
            <Cross className="text-red-500 h-4 w-4 rotate-45" />
          )}
        </div>
      );
    },
  },
  {
    accessorKey: "Dinge",
    header: "Dinge",
    cell: ({ row }) => {
      const x = row.original;
      return <pre className="font-sans">{x.Dinge}</pre>;
    },
  },
  {
    accessorKey: "Bild1",
    header: "Bilder",
    cell: ({ row }) => {
      const x = row.original;
      let b1_ok = false;
      let b2_ok = false;
      let b3_ok = false;

      if (x.Bild1Date.Valid) {
        const diff = Math.ceil(
          (new Date().getTime() - new Date(x.Bild1Date.Time).getTime()) /
            (1000 * 60 * 60 * 24)
        );

        if (diff == 1) {
          b1_ok = true;
        }
      }

      if (x.Bild2Date.Valid) {
        const diff = Math.ceil(
          (new Date().getTime() - new Date(x.Bild2Date.Time).getTime()) /
            (1000 * 60 * 60 * 24)
        );
        if (diff == 1) {
          b2_ok = true;
        }
      }

      if (x.Bild3Date.Valid) {
        const diff = Math.ceil(
          (new Date().getTime() - new Date(x.Bild3Date.Time).getTime()) /
            (1000 * 60 * 60 * 24)
        );
        if (diff == 1) {
          b3_ok = true;
        }
      }

      return (
        <div className="grid grid-cols-3">
          {b1_ok && (
            <img src={x.Bild1} className="h-15 w-auto object-contain" />
          )}
          {b2_ok && (
            <img src={x.Bild2} className="h-15 w-auto object-contain" />
          )}
          {b3_ok && (
            <img src={x.Bild3} className="h-15 w-auto object-contain" />
          )}
        </div>
      );
    },
  },
];
