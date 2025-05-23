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
import { Switch } from "@/components/ui/switch";
import { Textarea } from "@/components/ui/textarea";
import { Mitarbeiter } from "@api/db";
import { zodResolver } from "@hookform/resolvers/zod";
import { UploadImage } from "@wails/go/main/App";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";

const formSchema = z.object({
  Dinge: z.string().optional(),
  Geld: z.string().optional(),
  Pfand: z.string().optional(),
  Abo: z.boolean(),
  Paypal: z.boolean(),
});

export default function EinkaufForm({
  mitarbeiter,
}: {
  mitarbeiter: db.Mitarbeiter;
}) {
  const [loading, setLoading] = useState(false);

  const [uploaded1, setUploaded1] = useState(false);
  const [uploaded2, setUploaded2] = useState(false);
  const [uploaded3, setUploaded3] = useState(false);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      Dinge: mitarbeiter.Dinge ?? "",
      Geld: mitarbeiter.Geld ?? "",
      Pfand: mitarbeiter.Pfand ?? "",
      Abo: mitarbeiter.Abonniert ?? false,
      Paypal: mitarbeiter.Paypal ?? false,
    },
  });
  const navigate = useNavigate();

  useEffect(() => {
    form.reset({
      Dinge: mitarbeiter.Dinge ?? "",
      Geld: mitarbeiter.Geld ?? "",
      Pfand: mitarbeiter.Pfand ?? "",
      Abo: mitarbeiter.Abonniert ?? false,
      Paypal: mitarbeiter.Paypal ?? false,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [mitarbeiter]);

  const handleSubmit = async (values: z.infer<typeof formSchema>) => {
    setLoading(true);
    await Mitarbeiter.UpdateEinkauf(mitarbeiter.ID, {
      Abonniert: values.Abo,
      Paypal: values.Paypal,
      Dinge: values.Dinge,
      Geld: values.Geld,
      Pfand: values.Pfand,
      Bild1: uploaded1,
      Bild2: uploaded2,
      Bild3: uploaded3,
    });
    setLoading(false);
    await navigate("/");
  };

  return (
    <div className="max-w-200 mx-auto mt-12">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(handleSubmit)} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <FormField
              control={form.control}
              name="Geld"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Geld</FormLabel>
                  <FormControl>
                    <Input disabled={loading} {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Pfand"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Pfand</FormLabel>
                  <FormControl>
                    <Input disabled={loading} {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className="grid grid-cols-2 gap-4">
            <FormField
              control={form.control}
              name="Paypal"
              render={({ field }) => (
                <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                  <div className="space-y-0.5">
                    <FormLabel>Paypal</FormLabel>
                  </div>
                  <FormControl>
                    <Switch
                      disabled={loading}
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Abo"
              render={({ field }) => (
                <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                  <div className="space-y-0.5">
                    <FormLabel>Abo</FormLabel>
                  </div>
                  <FormControl>
                    <Switch
                      disabled={loading}
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <FormField
            control={form.control}
            name="Dinge"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Dein Einkauf</FormLabel>
                <FormControl>
                  <Textarea disabled={loading} {...field} />
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />
          <div className="grid grid-cols-3 gap-4">
            {!uploaded1 ? (
              <Input
                type="file"
                disabled={loading || uploaded1}
                onClick={async (e) => {
                  e.preventDefault();
                  setLoading(true);
                  await UploadImage(mitarbeiter.ID, 1);
                  setUploaded1(true);
                  setLoading(false);
                }}
              />
            ) : (
              "Bild bereits gespeichert"
            )}
            {!uploaded2 ? (
              <Input
                type="file"
                disabled={loading || uploaded2}
                onClick={async (e) => {
                  e.preventDefault();
                  setLoading(true);
                  await UploadImage(mitarbeiter.ID, 2);
                  setUploaded2(true);
                  setLoading(false);
                }}
              />
            ) : (
              "Bild bereits gespeichert"
            )}
            {!uploaded3 ? (
              <Input
                type="file"
                disabled={loading || uploaded3}
                onClick={async (e) => {
                  e.preventDefault();
                  setLoading(true);
                  setUploaded3(true);
                  await UploadImage(mitarbeiter.ID, 3);
                  setLoading(false);
                }}
              />
            ) : (
              "Bild bereits gespeichert"
            )}
          </div>
          <Button type="submit" disabled={loading}>
            {loading ? "Speichert ..." : "Speichern"}
          </Button>
        </form>
      </Form>
      <div className="grid grid-cols-2 gap-4 mt-5">
        <Button
          disabled={loading}
          variant={"secondary"}
          onClick={async (e) => {
            e.preventDefault();
            setLoading(true);
            await Mitarbeiter.EinkaufSkip(mitarbeiter.ID);
            await navigate("/");
            setLoading(false);
          }}
        >
          Einkauf Überspringen
        </Button>
        <Button
          disabled={loading}
          variant={"destructive"}
          onClick={async (e) => {
            e.preventDefault();
            setLoading(true);
            await Mitarbeiter.EinkaufDelete(mitarbeiter.ID);
            await navigate("/");
            setLoading(false);
          }}
        >
          Einkauf Löschen
        </Button>
      </div>
    </div>
  );
}
