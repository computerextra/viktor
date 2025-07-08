"use client";

import { baseURL } from "@/api";
import {
  ArchiveProps,
  SearchArchive,
  type ArchiveResponse,
} from "@/api/Archiv";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import type { ColumnDef } from "@tanstack/react-table";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { Link } from "react-router";
import { z } from "zod";
import ArchiveTable from "./table";

export default function ArchiveForm() {
  const [results, setResults] = useState<undefined | ArchiveResponse[]>(
    undefined
  );
  const [loading, setLoading] = useState(false);

  const form = useForm<z.infer<typeof ArchiveProps>>({
    resolver: zodResolver(ArchiveProps),
  });

  const onSubmit = async (values: z.infer<typeof ArchiveProps>) => {
    setLoading(true);
    const res = await SearchArchive(values);
    setResults(res);
    setLoading(false);
  };

  const columns: ColumnDef<ArchiveResponse>[] = [
    {
      accessorKey: "title",
      header: "Title",
      cell: ({ row }) => {
        const x = row.original;

        return (
          <Link
            className="cursor-pointer"
            target="_blank"
            to={baseURL + "/Archiv/" + x.id}
          >
            {x?.title}
          </Link>
        );
      },
    },
    {
      accessorKey: "id",
      header: "Download",
      cell: ({ row }) => {
        const x = row.original;

        return (
          <Button asChild>
            <Link
              className="cursor-pointer"
              target="_blank"
              to={baseURL + "/Archiv/" + x.id}
            >
              Download
            </Link>
          </Button>
        );
      },
    },
  ];

  return (
    <>
      <div className="mt-12 flex justify-center">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <FormField
              control={form.control}
              name="search"
              render={({ field }) => (
                <FormItem>
                  <div className="flex w-full max-w-sm items-center space-x-2">
                    <FormControl>
                      <Input
                        placeholder="Suche..."
                        disabled={loading}
                        {...field}
                      />
                    </FormControl>
                    <Button variant="noShadow" type="submit" disabled={loading}>
                      {loading ? "Sucht..." : "Suchen"}
                    </Button>
                  </div>
                  <FormMessage />
                </FormItem>
              )}
            />
          </form>
        </Form>
      </div>
      <div className="mt-5">
        <ArchiveTable columns={columns} data={results ?? []} />
      </div>
    </>
  );
}
