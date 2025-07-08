import {
  AbteilungProps,
  CreateAbteilung,
  GetAbteilung,
  UpdateAbteilung,
  type GetAbteilungeRes,
} from "@/api/Abteilungen";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import z from "zod";

export default function AbteilungForm({ id }: { id?: string }) {
  const navigate = useNavigate();

  const [Abteilung, setAbteilung] = useState<GetAbteilungeRes | undefined>(
    undefined
  );
  const [loading, setLoading] = useState(false);

  const form = useForm<z.infer<typeof AbteilungProps>>({
    resolver: zodResolver(AbteilungProps),
    defaultValues: {
      name: Abteilung?.name,
    },
  });

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const res = await GetAbteilung(id);
      setAbteilung(res);
      setLoading(false);
    })();
  }, [id]);

  useEffect(() => {
    if (Abteilung == null) return;

    form.reset({
      name: Abteilung.name,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [Abteilung]);

  const onSubmit = async (values: z.infer<typeof AbteilungProps>) => {
    if (id != null && Abteilung != null) {
      await UpdateAbteilung(id, values);
    } else {
      await CreateAbteilung(values);
    }
    await navigate("/CMS/Abteilungen");
  };

  if (loading) return <>Lädt...</>;

  return (
    <div className="mt-12 flex justify-center">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="mb-5 space-y-8">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit">Speichern</Button>
        </form>
      </Form>
    </div>
  );
}
