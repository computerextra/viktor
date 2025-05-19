import type { ColumnDef } from "@tanstack/react-table";
import type { db } from "@wails/go/models";
import { Link } from "react-router";

export const columns: ColumnDef<db.Mitarbeiter>[] = [
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
    accessorKey: "InternTelefon1",
    header: "Interne Durchwahl",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <div className="grid grid-cols-2">
          {x.Gruppenwahl && x.Gruppenwahl.length > 0 && (
            <>
              <span>Gruppe</span>
              <span>{x.Gruppenwahl}</span>
            </>
          )}
          {x.InternTelefon1 && x.InternTelefon1.length > 0 && (
            <>
              <span>DW 1</span>
              <span>{x.InternTelefon1}</span>
            </>
          )}
          {x.InternTelefon2 && x.InternTelefon2.length > 0 && (
            <>
              <span>DW 2</span>
              <span>{x.InternTelefon2}</span>
            </>
          )}
          {x.HomeOffice && x.HomeOffice.length > 0 && (
            <>
              <span>Homeoffice</span>
              <span>{x.HomeOffice}</span>
            </>
          )}
        </div>
      );
    },
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
              <a href={"tel:" + x.FestnetzPrivat}>{x.FestnetzPrivat}</a>
            </>
          )}
          {x.FestnetzBusiness && x.FestnetzBusiness.length > 0 && (
            <>
              <span>Business</span>
              <a href={"tel:" + x.FestnetzBusiness}>{x.FestnetzBusiness}</a>
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
              <a href={"tel:" + x.MobilPrivat}>{x.MobilPrivat}</a>
            </>
          )}
          {x.MobilBusiness && x.MobilBusiness.length > 0 && (
            <>
              <span>Business</span>
              <a href={"tel:" + x.MobilBusiness}>{x.MobilBusiness}</a>
            </>
          )}
        </div>
      );
    },
  },
];
