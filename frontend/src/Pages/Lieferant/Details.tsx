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
import { Ansprechpartner, Lieferant } from "@api/db";
import type { db } from "@wails/go/models";

import { useEffect, useState } from "react";
import { Link, useParams } from "react-router";
import { ApColumns } from "./_components/columns";
import { DataTable } from "./_components/data-table";

export const LieferantenDetails = () => {
  const { id } = useParams();

  const [lieferant, setLieferant] = useState<db.Lieferant | undefined>(
    undefined
  );

  useEffect(() => {
    (async () => {
      if (id == null) return;
      const i = parseInt(id);
      const res = await Lieferant.Get(i);
      setLieferant(res);
    })();
  }, [id]);

  return (
    <>
      <BackButton href="/Lieferant" />
      <div className="container mx-auto">
        <Card className="w-[800px] mx-auto">
          <CardHeader>
            <CardTitle>{lieferant?.Firma}</CardTitle>
          </CardHeader>
          <CardContent>
            <Table className="mb-5">
              <TableBody>
                <TableRow>
                  <TableCell className="font-medium">Kundennummer:</TableCell>
                  <TableCell>{lieferant?.Kundennummer}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell className="font-medium">Webseite:</TableCell>
                  <TableCell>
                    <a
                      href={lieferant?.Webseite}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="underline text-primary"
                    >
                      {lieferant?.Webseite}
                    </a>
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
            <h2>Ansprechpartner</h2>
            <DataTable
              columns={ApColumns}
              data={lieferant?.Ansprechpartner ? lieferant.Ansprechpartner : []}
            />
          </CardContent>
          <CardFooter>
            <Button asChild>
              <Link to={`/Lieferant/${id}/Bearbeiten`}>Bearbeiten</Link>
            </Button>
          </CardFooter>
        </Card>
      </div>
    </>
  );
};

export const AnsprechpartnerDetails = () => {
  const { id, ap } = useParams();

  const [ansprechpartner, setAnsprechpartner] = useState<
    db.Ansprechpartner | undefined
  >(undefined);

  useEffect(() => {
    (async () => {
      if (ap == null) return;
      const i = parseInt(ap);
      const res = await Ansprechpartner.Get(i);
      setAnsprechpartner(res);
    })();
  }, [ap]);

  return (
    <>
      <BackButton href={"/Lieferant/" + id} />
      <div className="container mx-auto">
        <Card className="w-[800px] mx-auto">
          <CardHeader>
            <CardTitle>{ansprechpartner?.Name}</CardTitle>
          </CardHeader>
          <CardContent>
            <Table className="mb-5">
              <TableBody>
                <TableRow>
                  <TableCell className="font-medium">Telefon:</TableCell>
                  <TableCell>
                    <a
                      className="underline text-primary"
                      href={"tel:" + ansprechpartner?.Telefon}
                    >
                      {ansprechpartner?.Telefon}
                    </a>
                  </TableCell>
                </TableRow>
                <TableRow>
                  <TableCell className="font-medium">Mobil:</TableCell>
                  <TableCell>
                    <a
                      className="underline text-primary"
                      href={"tel:" + ansprechpartner?.Mobil}
                    >
                      {ansprechpartner?.Mobil}
                    </a>
                  </TableCell>
                </TableRow>
                <TableRow>
                  <TableCell className="font-medium">Mail:</TableCell>
                  <TableCell>
                    <a
                      className="underline text-primary"
                      href={"mailto:" + ansprechpartner?.Mail}
                    >
                      {ansprechpartner?.Mail}
                    </a>
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </CardContent>
          <CardFooter>
            <Button asChild>
              <Link to={`/Lieferant/${id}/${ap}/Bearbeiten`}>Bearbeiten</Link>
            </Button>
          </CardFooter>
        </Card>
      </div>
    </>
  );
};
