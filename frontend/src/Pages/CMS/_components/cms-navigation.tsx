import { GetCmsCounter, type GetCmsCounterResponse } from "@/api/cms";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useEffect, useState } from "react";
import { Link } from "react-router";

export default function CmsNavigation() {
  const [Counts, setCounts] = useState<GetCmsCounterResponse | undefined>(
    undefined
  );
  const [loading, setLoading] = useState(false);
  const [err, setErr] = useState<string | undefined>(undefined);

  useEffect(() => {
    (async () => {
      setLoading(true);
      const res = await GetCmsCounter();
      if (res == null) {
        setErr("Fehler beim abrufen von Daten");
        setLoading(false);
        return;
      }
      setCounts(res);
      setLoading(false);
    })();
  }, []);

  if (loading) return <>Lädt...</>;
  if (err) return <>{err}</>;

  const data = [
    {
      title: "Abteilungen",
      count: Counts?.Abteilungen,
    },
    {
      title: "Angebote",
      count: Counts?.Angebote,
    },
    {
      title: "Jobs",
      count: Counts?.Jobs,
    },
    {
      title: "Mitarbeiter",
      count: Counts?.Mitarbeiter,
    },
    {
      title: "Partner",
      count: Counts?.Partner,
    },
  ];

  return (
    <Table className="mt-5">
      <TableHeader>
        <TableRow>
          <TableHead>Titel</TableHead>
          <TableHead>Mege der Datensätze</TableHead>
          <TableHead className="text-right">Go there</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {data.map((x) => (
          <TableRow key={x.title}>
            <TableCell className="font-base">{x.title}</TableCell>
            <TableCell>{x.count}</TableCell>
            <TableCell className="text-right">
              <Button asChild variant={"reverse"}>
                <Link to={"/CMS/" + x.title}>Link</Link>
              </Button>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
