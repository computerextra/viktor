import BackButton from "@/components/BackButton";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table";
import { Mitarbeiter } from "@api/db";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { Link, useParams } from "react-router";

export default function MitarbeiterDetails() {
  const { id } = useParams();
  const [mitarbeiter, setMitarbeiter] = useState<db.Mitarbeiter | undefined>(
    undefined
  );

  useEffect(() => {
    async function x() {
      if (id == null) return;
      const i = parseInt(id);
      const res = await Mitarbeiter.Get(i);
      setMitarbeiter(res);
    }

    x();
  }, [id]);

  const data: { title: string; description?: React.ReactNode }[] = [
    {
      title: "E-Mail",
      description:
        mitarbeiter?.Email && mitarbeiter.Email.length > 0 ? (
          <a
            href={"mailto:" + mitarbeiter.Email}
            className="underline text-primary"
          >
            {mitarbeiter.Email}
          </a>
        ) : (
          "-"
        ),
    },
    {
      title: "Gruppenwahl",
      description:
        mitarbeiter?.Gruppenwahl && mitarbeiter.Gruppenwahl.length > 0
          ? mitarbeiter.Gruppenwahl
          : "-",
    },
    {
      title: "Interne Durchwahl 1",
      description:
        mitarbeiter?.InternTelefon1 && mitarbeiter.InternTelefon1.length > 0
          ? mitarbeiter.InternTelefon1
          : "-",
    },
    {
      title: "Interne Durchwahl 2",
      description:
        mitarbeiter?.InternTelefon2 && mitarbeiter.InternTelefon2.length > 0
          ? mitarbeiter.InternTelefon2
          : "-",
    },
    {
      title: "Home Office",
      description:
        mitarbeiter?.HomeOffice && mitarbeiter.HomeOffice.length > 0
          ? mitarbeiter.HomeOffice
          : "-",
    },
    {
      title: "Festnetz Privat",
      description:
        mitarbeiter?.FestnetzPrivat && mitarbeiter.FestnetzPrivat.length > 0 ? (
          <a
            href={"tel:" + mitarbeiter.FestnetzPrivat}
            className="underline text-primary"
          >
            {mitarbeiter.FestnetzPrivat}
          </a>
        ) : (
          "-"
        ),
    },
    {
      title: "Festnetz Business",
      description:
        mitarbeiter?.FestnetzBusiness &&
        mitarbeiter.FestnetzBusiness.length > 0 ? (
          <a
            href={"tel:" + mitarbeiter.FestnetzBusiness}
            className="underline text-primary"
          >
            {mitarbeiter.FestnetzBusiness}
          </a>
        ) : (
          "-"
        ),
    },
    {
      title: "Mobil Privat",
      description:
        mitarbeiter?.MobilPrivat && mitarbeiter.MobilPrivat.length > 0 ? (
          <a
            href={"tel:" + mitarbeiter.MobilPrivat}
            className="underline text-primary"
          >
            {mitarbeiter.MobilPrivat}
          </a>
        ) : (
          "-"
        ),
    },
    {
      title: "Mobil Business",
      description:
        mitarbeiter?.MobilBusiness && mitarbeiter.MobilBusiness.length > 0 ? (
          <a
            href={"tel:" + mitarbeiter.MobilBusiness}
            className="underline text-primary"
          >
            {mitarbeiter.MobilBusiness}
          </a>
        ) : (
          "-"
        ),
    },
    {
      title: "Short Code",
      description: mitarbeiter?.Short ?? "-",
    },
  ];

  return (
    <>
      <BackButton href="/Mitarbeiter" />
      <div className="container mx-auto">
        <Card className="w-[500px] mx-auto">
          <CardHeader>
            <CardTitle>{mitarbeiter?.Name}</CardTitle>
          </CardHeader>
          <CardContent>
            <Table>
              <TableBody>
                {data.map((notification, index) => (
                  <TableRow key={index}>
                    <TableCell className="font-medium">
                      {notification.title}
                    </TableCell>
                    <TableCell>{notification.description}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </CardContent>
          <CardFooter>
            <Button asChild>
              <Link to={`/Mitarbeiter/${id}/Bearbeiten`}>Bearbeiten</Link>
            </Button>
          </CardFooter>
        </Card>
      </div>
    </>
  );
}
