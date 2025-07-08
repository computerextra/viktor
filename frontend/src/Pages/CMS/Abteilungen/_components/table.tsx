import {
  DeleteAbteilung,
  GetAbteilungen,
  type GetAbteilungeRes,
} from "@/api/Abteilungen";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import type { ColumnDef } from "@tanstack/react-table";
import { MoreHorizontal } from "lucide-react";
import { useEffect, useState } from "react";
import { Link } from "react-router";
import CmsTable from "../../_components/cms-table";

export default function AbteilungenTable() {
  const [Abteilungen, setAbteilungen] = useState<
    GetAbteilungeRes[] | undefined
  >(undefined);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    (async () => {
      setLoading(true);
      const res = await GetAbteilungen();
      setAbteilungen(res);
      setLoading(false);
    })();
  }, []);

  if (loading) return <>Lädt...</>;

  const columns: ColumnDef<GetAbteilungeRes>[] = [
    {
      accessorKey: "name",
      header: "Name",
    },
    {
      id: "actions",
      cell: ({ row }) => {
        const x = row.original;
        return (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="noShadow" className="h-8 w-8 p-0">
                <span className="sr-only">Open menu</span>
                <MoreHorizontal className="h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>Actions</DropdownMenuLabel>
              <DropdownMenuItem asChild>
                <Link to={"/CMS/Abteilungen/" + x.id}>Bearbeiten</Link>
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={async () => {
                  await DeleteAbteilung(x.id);
                  location.reload();
                }}
              >
                Löschen
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        );
      },
    },
  ];

  return (
    <div className="mt-5">
      <Button asChild className="mb-5">
        <Link to="/CMS/Abteilungen/Neu">Neue Abteilung anlegen</Link>
      </Button>
      <CmsTable columns={columns} data={Abteilungen ?? []} />
    </div>
  );
}
