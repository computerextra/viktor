import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Search } from "@api/archive";
import { zodResolver } from "@hookform/resolvers/zod";
import type { archive } from "@wails/go/models";
import { useState } from "react";
import { useForm } from "react-hook-form";
import z from "zod";
import { columns } from "./_components/columns";
import { DataTable } from "./_components/data-table";

const formSchema = z.object({
  Suche: z.string(),
});

export default function Archiv() {
  const [loading, setLoading] = useState(false);
  const [results, setResults] = useState<
    Array<archive.ArchiveResult> | undefined
  >(undefined);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  const handleSubmit = async (values: z.infer<typeof formSchema>) => {
    setLoading(true);
    const res = await Search(values.Suche);
    console.log(res);
    setResults(res);
    setLoading(false);
  };

  return (
    <>
      <h1 className="text-center">CE Archiv</h1>
      <div className="container mx-auto">
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="mt-12 space-y-4 w-full"
          >
            <div className="flex flex-row w-full justify-center">
              <FormField
                control={form.control}
                name="Suche"
                render={({ field }) => (
                  <FormItem>
                    <FormControl>
                      <Input
                        placeholder="Suchebegriff"
                        className="w-2xl"
                        disabled={loading}
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <Button type="submit" disabled={loading}>
                {loading ? "Sucht ..." : "Suchen"}
              </Button>
            </div>
          </form>
        </Form>
      </div>
      <div className="container mx-auto mt-12">
        <DataTable columns={columns} data={results ? results : []} />
      </div>
    </>
  );
}
