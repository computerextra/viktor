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

  const handleSubmit = async (values: z.infer<typeof formSchema>) => {
    await Mitarbeiter.UpdateEinkauf(mitarbeiter.ID, {
      Abonniert: values.Abo,
      Paypal: values.Paypal,
      Dinge: values.Dinge,
      Geld: values.Geld,
      Pfand: values.Pfand,
    });
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
                    <Input {...field} />
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
                    <Input {...field} />
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
                  <Textarea {...field} />
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />
          <div className="grid grid-cols-3 gap-4">
            <Input
              type="file"
              onClick={async (e) => {
                e.preventDefault();
                await UploadImage(mitarbeiter.ID, 1);
              }}
            />
            <Input
              type="file"
              onClick={async (e) => {
                e.preventDefault();
                await UploadImage(mitarbeiter.ID, 2);
              }}
            />
            <Input
              type="file"
              onClick={async (e) => {
                e.preventDefault();
                await UploadImage(mitarbeiter.ID, 3);
              }}
            />
          </div>
          <Button type="submit">Speichern</Button>
        </form>
      </Form>
    </div>
  );
}
