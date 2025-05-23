import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import type { ColumnDef } from "@tanstack/react-table";
import type { db } from "@wails/go/models";
import { MoreHorizontal } from "lucide-react";
import { NavLink } from "react-router";
import { DataTable } from "./data-table";

export const columns: ColumnDef<db.Lieferant>[] = [
  {
    accessorKey: "Firma",
    header: "Firma",
    cell: ({ row }) => {
      const x = row.original;
      return <NavLink to={`/Lieferant/` + x.ID}>{x.Firma}</NavLink>;
    },
  },
  {
    accessorKey: "Kundennummer",
    header: "Kundennummer",
  },
  {
    accessorKey: "Webseite",
    header: "Webseite",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <a
          href={x.Webseite}
          target="_blank"
          rel="noopener noreferrer"
          className="underline text-primary"
        >
          {x.Webseite}
        </a>
      );
    },
  },
  {
    accessorKey: "Ansprechpartner",
    header: "Ansprechpartner",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <DataTable
          columns={ApColumns}
          data={x.Ansprechpartner.sort((a, b) => {
            const nameA = a.Name.toLowerCase();
            const nameB = b.Name.toLowerCase();
            if (nameA < nameB) return -1;
            if (nameA > nameB) return 1;
            return 0;
          })}
        />
      );
    },
  },
  {
    id: "actions",
    cell: ({ row }) => {
      const x = row.original;

      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>
            <DropdownMenuItem asChild>
              <NavLink to={`/Lieferant/${x.ID}/Neu`}>
                Neuer Ansprechpartner
              </NavLink>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];

export const ApColumns: ColumnDef<db.Ansprechpartner>[] = [
  {
    accessorKey: "Name",
    header: "Name",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <NavLink to={`/Lieferant/${x.LieferantenId}/${x.ID}`}>{x.Name}</NavLink>
      );
    },
  },
  {
    accessorKey: "Mail",
    header: "Mail",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <a href={"mailto:" + x.Mail} className="underline text-primary">
          {x.Mail}
        </a>
      );
    },
  },
  {
    accessorKey: "Mobil",
    header: "Mobil",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <a href={"tel:" + x.Mobil} className="underline text-primary">
          {x.Mobil}
        </a>
      );
    },
  },
  {
    accessorKey: "Telefon",
    header: "Telefon",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <a href={"tel:" + x.Telefon} className="underline text-primary">
          {x.Telefon}
        </a>
      );
    },
  },
];
