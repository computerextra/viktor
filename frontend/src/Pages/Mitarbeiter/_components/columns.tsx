import type { ColumnDef } from "@tanstack/react-table";
import type { db } from "@wails/go/models";
import { Check, Cross } from "lucide-react";
import { Link } from "react-router";

export const columns: ColumnDef<db.Mitarbeiter>[] = [
  {
    accessorKey: "Azubi",
    header: "Azubi",
    cell: ({ row }) => {
      const x = row.original;
      if (x.Azubi) return <Check className="text-green-500 h-4 w-4" />;
      if (!x.Azubi) return <Cross className="text-red-500 h-4 w-4 rotate-45" />;
    },
  },
  {
    accessorKey: "Name",
    header: "Name",
    cell: ({ row }) => {
      const x = row.original;
      return <Link to={`/Mitarbeiter/${x.ID}`}>{x.Name}</Link>;
    },
  },
  {
    accessorKey: "Email",
    header: "Email",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <a href={"mailto:" + x.Email} className="underline text-primary">
          {x.Email}
        </a>
      );
    },
  },
  {
    accessorKey: "Gruppenwahl",
    header: "Interne Durchwahl",
  },
  {
    accessorKey: "FestnetzBusiness",
    header: "Festnetz",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <div className="grid grid-cols-2">
          {x.FestnetzPrivat && x.FestnetzPrivat.length > 0 && (
            <>
              <span>Privat</span>
              <a
                className="underline text-primary"
                href={"tel:" + x.FestnetzPrivat}
              >
                {x.FestnetzPrivat}
              </a>
            </>
          )}
          {x.FestnetzBusiness && x.FestnetzBusiness.length > 0 && (
            <>
              <span>Business</span>
              <a
                className="underline text-primary"
                href={"tel:" + x.FestnetzBusiness}
              >
                {x.FestnetzBusiness}
              </a>
            </>
          )}
        </div>
      );
    },
  },
  {
    accessorKey: "MobilBusinesss",
    header: "Mobil",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <div className="grid grid-cols-2">
          {x.MobilPrivat && x.MobilPrivat.length > 0 && (
            <>
              <span>Privat</span>
              <a
                className="underline text-primary"
                href={"tel:" + x.MobilPrivat}
              >
                {x.MobilPrivat}
              </a>
            </>
          )}
          {x.MobilBusiness && x.MobilBusiness.length > 0 && (
            <>
              <span>Business</span>
              <a
                className="underline text-primary"
                href={"tel:" + x.MobilBusiness}
              >
                {x.MobilBusiness}
              </a>
            </>
          )}
        </div>
      );
    },
  },
];
