import {
  AngeboteProps,
  type AngeboteRes,
  CreateAngebot,
  GetAngebot,
  UpdateAngebot,
} from "@/api/Angebote";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import z from "zod";

export default function AngebotForm({ id }: { id?: string }) {
  const navigate = useNavigate();
  const [Angebot, setAngebot] = useState<AngeboteRes | undefined>(undefined);
  const [loading, setLoading] = useState(false);
  const form = useForm<z.infer<typeof AngeboteProps>>({
    resolver: zodResolver(AngeboteProps),
    defaultValues: {
      anzeigen: Angebot?.anzeigen ?? false,
      date_start: Angebot?.date_start
        ? new Date(Angebot.date_start)
        : undefined,
      date_stop: Angebot?.date_stop ? new Date(Angebot.date_stop) : undefined,
      image: Angebot?.image,
      link: Angebot?.link,
      subtitle: Angebot?.subtitle,
      title: Angebot?.title,
    },
  });

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const res = await GetAngebot(id);
      setAngebot(res);
      setLoading(false);
    })();
  }, [id]);

  useEffect(() => {
    if (Angebot == null) return;

    form.reset({
      anzeigen: Angebot?.anzeigen ?? false,
      date_start: new Date(Angebot.date_start),
      date_stop: new Date(Angebot.date_stop),
      image: Angebot?.image,
      link: Angebot?.link,
      subtitle: Angebot?.subtitle,
      title: Angebot?.title,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [Angebot]);

  const onSubmit = async (values: z.infer<typeof AngeboteProps>) => {
    if (id != null && Angebot != null) {
      await UpdateAngebot(id, values);
    } else {
      await CreateAngebot(values);
    }
    await navigate("/CMS/Angebote");
  };

  if (loading) return <>Lädt...</>;

  return (
    <div className="mt-12 flex justify-center">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="mb-5 space-y-8">
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Title</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="subtitle"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Sub Title</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="link"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Link</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="image"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Bild</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="anzeigen"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                <div className="space-y-0.5">
                  <FormLabel>Auf Webseite anzeigen?</FormLabel>
                  <FormDescription>
                    Wenn aktiviert, wird das Angebot auf der Webseite angezeigt
                  </FormDescription>
                </div>
                <FormControl>
                  <Switch
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                </FormControl>
              </FormItem>
            )}
          />
          <div className="grid grid-cols-2 gap-8">
            <FormField
              control={form.control}
              name="date_start"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    Laufzeit von{" "}
                    {Angebot?.date_start
                      ? new Date(Angebot.date_start).toLocaleDateString("de-de")
                      : ""}
                  </FormLabel>
                  <FormControl>
                    <Calendar
                      mode="single"
                      selected={field.value}
                      onSelect={field.onChange}
                      captionLayout="dropdown"
                    />
                  </FormControl>
                  <FormDescription>
                    {field.value && (
                      <>Ausgewählt: {field.value.toLocaleDateString("de-de")}</>
                    )}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="date_stop"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    Laufzeit bis{" "}
                    {Angebot?.date_stop
                      ? new Date(Angebot.date_stop).toLocaleDateString("de-de")
                      : ""}
                  </FormLabel>
                  <FormControl>
                    <Calendar
                      mode="single"
                      selected={field.value}
                      onSelect={field.onChange}
                      captionLayout="dropdown"
                    />
                  </FormControl>
                  <FormDescription>
                    {field.value && (
                      <>Ausgewählt: {field.value.toLocaleDateString("de-de")}</>
                    )}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <Button type="submit">Speichern</Button>
        </form>
      </Form>
    </div>
  );
}
