import type { ColumnDef } from "@tanstack/react-table";
import type { sagedb } from "@wails/go/models";

export const columns: ColumnDef<sagedb.SearchResult>[] = [
  {
    accessorKey: "KundNr",
    header: "Kunden/Lieferant",
    cell: ({ row }) => {
      const x = row.original;
      if (x.KundNr && x.KundNr.length > 0) return "Kunde";
      if (x.LiefNr && x.LiefNr.length > 0) return "Lieferant";
    },
  },
  {
    accessorKey: "KundNr",
    header: "Nummer",
    cell: ({ row }) => {
      const x = row.original;
      if (x.KundNr && x.KundNr.length > 0) return x.KundNr;
      if (x.LiefNr && x.LiefNr.length > 0) return x.LiefNr;
    },
  },
  {
    accessorKey: "Suchbegriff",
    header: "Name",
  },
  {
    accessorKey: "Homepage",
    header: "Webseite",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <a
          href={x.Homepage}
          target="_blank"
          rel="noopener noreferrer"
          className="underline text-primary"
        >
          {x.Homepage}
        </a>
      );
    },
  },
  {
    accessorKey: "Telefon1",
    header: "Telefon",
    cell: ({ row }) => {
      const x = row.original;

      return (
        <div className="grid">
          {x.Telefon1 && x.Telefon1.length > 0 && (
            <a href={"tel:" + x.Telefon1} className="underline text-primary">
              {x.Telefon1}
            </a>
          )}
          {x.Telefon2 && x.Telefon2.length > 0 && (
            <a href={"tel:" + x.Telefon2} className="underline text-primary">
              {x.Telefon2}
            </a>
          )}
        </div>
      );
    },
  },
  {
    accessorKey: "Mobiltelefon1",
    header: "Mobil",
    cell: ({ row }) => {
      const x = row.original;

      return (
        <div className="grid">
          {x.Mobiltelefon1 && x.Mobiltelefon1.length > 0 && (
            <a
              href={"tel:" + x.Mobiltelefon1}
              className="underline text-primary"
            >
              {x.Mobiltelefon1}
            </a>
          )}
          {x.Mobiltelefon2 && x.Mobiltelefon2.length > 0 && (
            <a
              href={"tel:" + x.Mobiltelefon2}
              className="underline text-primary"
            >
              {x.Mobiltelefon2}
            </a>
          )}
        </div>
      );
    },
  },
  {
    accessorKey: "EMail1",
    header: "Email",
    cell: ({ row }) => {
      const x = row.original;

      return (
        <div className="grid">
          {x.EMail1 && x.EMail1.length > 0 && (
            <a href={"mailto:" + x.EMail1} className="underline text-primary">
              {x.EMail1}
            </a>
          )}
          {x.EMail2 && x.EMail2.length > 0 && (
            <a href={"mailto:" + x.EMail2} className="underline text-primary">
              {x.EMail2}
            </a>
          )}
        </div>
      );
    },
  },
  {
    accessorKey: "KundUmsatz",
    header: "Umsatz",
    cell: ({ row }) => {
      const x = row.original;
      const Umsatz = () => {
        if (x.KundUmsatz) return x.KundUmsatz.toFixed(2);
        if (x.LiefUmsatz) return x.LiefUmsatz.toFixed(2);

        return 0;
      };
      return <span>{Umsatz()}€</span>;
    },
  },
];
