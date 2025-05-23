import { Button } from "@/components/ui/button";
import { Download } from "@api/archive";
import type { ColumnDef } from "@tanstack/react-table";
import type { archive } from "@wails/go/models";

export const columns: ColumnDef<archive.ArchiveResult>[] = [
  {
    accessorKey: "Title",
    header: "Titel",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <Button variant="link" onClick={async () => await Download(x.Id)}>
          {x.Title}
        </Button>
      );
    },
  },
  {
    accessorKey: "Body",
    header: "Inhalt",
    cell: ({ row }) => {
      const x = row.original;
      return (
        <div className="truncate w-200">
          <p className="text-ellipsis">{x.Body}</p>
        </div>
      );
    },
  },
];
