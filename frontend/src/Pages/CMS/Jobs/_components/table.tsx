import { DeleteJob, GetJobs, type JobReponse } from "@/api/Jobs";
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
import { Check, Cross, MoreHorizontal } from "lucide-react";
import { useEffect, useState } from "react";
import { Link } from "react-router";
import CmsTable from "../../_components/cms-table";

const No = () => <Cross className="h-4 w-4 rotate-45 text-red-500" />;
const Yes = () => <Check className="h-4 w-4 text-green-500" />;

export default function JobTable() {
  const [Jobs, setJobs] = useState<JobReponse[] | undefined>(undefined);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    (async () => {
      setLoading(true);
      const res = await GetJobs();
      setJobs(res);
      setLoading(false);
    })();
  }, []);

  if (loading) return <>Lädt...</>;

  const columns: ColumnDef<JobReponse>[] = [
    {
      accessorKey: "name",
      header: "Name",
    },
    {
      accessorKey: "online",
      header: "Online",
      cell: ({ row }) => {
        const x = row.original;
        return <>{x.online ? <Yes /> : <No />}</>;
      },
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
                <Link to={"/CMS/Jobs/" + x.id}>Bearbeiten</Link>
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={async () => {
                  await DeleteJob(x.id);
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
        <Link to="/CMS/Jobs/Neu">Neuen Job anlegen</Link>
      </Button>
      <CmsTable columns={columns} data={Jobs ?? []} />
    </div>
  );
}
