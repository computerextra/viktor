import { GetAbteilungen, type GetAbteilungeRes } from "@/api/Abteilungen";
import {
  CreateMitarbeiter,
  GetMitarbeiter,
  MitarbeiterProps,
  UpdateMitarbeiter,
  type MitarbeiterRes,
} from "@/api/Mitarbeiter";
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Separator } from "@/components/ui/separator";
import { Switch } from "@/components/ui/switch";
import { zodResolver } from "@hookform/resolvers/zod";
import { de } from "date-fns/locale";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";

export default function MitarbeiterForm({ id }: { id?: string }) {
  const navigate = useNavigate();

  const [Abteilungen, setAbteilungen] = useState<
    GetAbteilungeRes[] | undefined
  >(undefined);
  const [Mitarbeiter, setMitarbeiter] = useState<MitarbeiterRes | undefined>(
    undefined
  );
  const [loading, setLoading] = useState(false);

  const form = useForm<z.infer<typeof MitarbeiterProps>>({
    resolver: zodResolver(MitarbeiterProps),
    defaultValues: {
      abteilungId: Mitarbeiter?.abteilungId,
      Azubi: Mitarbeiter?.Azubi,
      focus: Mitarbeiter?.focus,
      Geburtstag: Mitarbeiter?.Geburtstag
        ? new Date(Mitarbeiter.Geburtstag)
        : undefined,
      Gruppenwahl: Mitarbeiter?.Gruppenwahl,
      HomeOffice: Mitarbeiter?.HomeOffice,
      image: Mitarbeiter?.image,
      sex: Mitarbeiter?.sex,
      mail: Mitarbeiter?.mail,
      Mobil_Business: Mitarbeiter?.Mobil_Business,
      Mobil_Privat: Mitarbeiter?.Mobil_Privat,
      name: Mitarbeiter?.name,
      short: Mitarbeiter?.short,
      Telefon_Business: Mitarbeiter?.Telefon_Business,
      Telefon_Intern_1: Mitarbeiter?.Telefon_Intern_1,
      Telefon_Intern_2: Mitarbeiter?.Telefon_Intern_2,
      Telefon_Privat: Mitarbeiter?.Telefon_Privat,
    },
  });

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const ma = await GetMitarbeiter(id);
      setMitarbeiter(ma);
      const ab = await GetAbteilungen();
      setAbteilungen(ab);
      setLoading(false);
    })();
  }, [id]);

  useEffect(() => {
    if (Mitarbeiter == null) return;
    form.reset({
      abteilungId: Mitarbeiter?.abteilungId,
      Azubi: Mitarbeiter?.Azubi,
      focus: Mitarbeiter?.focus,
      Geburtstag: Mitarbeiter?.Geburtstag
        ? new Date(Mitarbeiter.Geburtstag)
        : undefined,
      Gruppenwahl: Mitarbeiter?.Gruppenwahl,
      HomeOffice: Mitarbeiter?.HomeOffice,
      image: Mitarbeiter?.image,
      sex: Mitarbeiter?.sex,
      mail: Mitarbeiter?.mail,
      Mobil_Business: Mitarbeiter?.Mobil_Business,
      Mobil_Privat: Mitarbeiter?.Mobil_Privat,
      name: Mitarbeiter?.name,
      short: Mitarbeiter?.short,
      Telefon_Business: Mitarbeiter?.Telefon_Business,
      Telefon_Intern_1: Mitarbeiter?.Telefon_Intern_1,
      Telefon_Intern_2: Mitarbeiter?.Telefon_Intern_2,
      Telefon_Privat: Mitarbeiter?.Telefon_Privat,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [Mitarbeiter]);

  const onSubmit = async (values: z.infer<typeof MitarbeiterProps>) => {
    if (id != null && Mitarbeiter != null) {
      await UpdateMitarbeiter(id, values);
    } else {
      await CreateMitarbeiter(values);
    }
    navigate("/CMS/Mitarbeiter");
  };

  if (loading) return <>Lädt...</>;

  return (
    <div className="mt-12 flex justify-center">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="mb-5 space-y-8">
          <div className="grid grid-cols-2 gap-8">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name*</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="short"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Short</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="sex"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Geschlecht*</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="..." />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="w">Weiblich</SelectItem>
                      <SelectItem value="m">Männlich</SelectItem>
                    </SelectContent>
                  </Select>

                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="abteilungId"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Abteilung</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="..." />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="x">ohne</SelectItem>
                      {Abteilungen?.map((x) => (
                        <SelectItem key={x.id} value={x.id}>
                          {x.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <FormField
            control={form.control}
            name="image"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                <div className="space-y-0.5">
                  <FormLabel>Bild auf Webseite?</FormLabel>
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
          <FormField
            control={form.control}
            name="focus"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Focus (Komma getrennt)</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="mail"
            render={({ field }) => (
              <FormItem>
                <FormLabel>E-Mail</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Separator />
          <div className="grid grid-cols-2 gap-8">
            <FormField
              control={form.control}
              name="Gruppenwahl"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Gruppenwahl</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Telefon_Intern_1"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Telefon_Intern_1</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Telefon_Intern_2"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Telefon_Intern_2</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="HomeOffice"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>HomeOffice</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Telefon_Privat"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Telefon_Privat</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Telefon_Business"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Telefon_Business</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Mobil_Privat"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Mobil_Privat</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Mobil_Business"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Mobil_Business</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <FormField
            control={form.control}
            name="Azubi"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                <div className="space-y-0.5">
                  <FormLabel>Azubi?</FormLabel>
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

          <FormField
            control={form.control}
            name={"Geburtstag"}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Geburtstag</FormLabel>
                <FormControl>
                  <Calendar
                    mode="single"
                    selected={field.value}
                    onSelect={field.onChange}
                    className="rounded-md border shadow-sm"
                    captionLayout="dropdown"
                    locale={de}
                  />
                </FormControl>
                <FormDescription>
                  Gespeicherter Wert:{" "}
                  {Mitarbeiter?.Geburtstag &&
                    new Date(Mitarbeiter.Geburtstag).toLocaleDateString()}{" "}
                  <br />
                  Gewählter Wert:{" "}
                  {form.getValues("Geburtstag")?.toLocaleDateString()}
                </FormDescription>
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
