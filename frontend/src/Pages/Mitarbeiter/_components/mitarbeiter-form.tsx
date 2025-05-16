import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { Create, Update } from "@api/db";
import { zodResolver } from "@hookform/resolvers/zod";
import type { db } from "@wails/go/models";
import { format } from "date-fns";
import { de } from "date-fns/locale";
import { CalendarIcon } from "lucide-react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";

const formSchema = z.object({
  Name: z.string(),
  Short: z.string().optional(),
  Gruppenwahl: z.string().optional(),
  InternTelefon1: z.string().optional(),
  InternTelefon2: z.string().optional(),
  FestnetzPrivat: z.string().optional(),
  FestnetzBusiness: z.string().optional(),
  HomeOffice: z.string().optional(),
  MobilBusiness: z.string().optional(),
  MobilPrivat: z.string().optional(),
  Email: z.string().optional(),
  Geburtstag: z.date().optional(),
  Azubi: z.boolean(),
});

export default function MitarbeiterForm({
  mitarbeiter,
}: {
  mitarbeiter?: db.MitarbeiterModel;
}) {
  const navigate = useNavigate();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      Short: mitarbeiter?.Short,
      Name: mitarbeiter?.Name,
      Email: mitarbeiter?.Email,
      FestnetzBusiness: mitarbeiter?.Festnetz_busines,
      FestnetzPrivat: mitarbeiter?.Festnetz_privat,
      Gruppenwahl: mitarbeiter?.Gruppenwahl,
      HomeOffice: mitarbeiter?.Home_office,
      InternTelefon1: mitarbeiter?.Intern_telefon1,
      InternTelefon2: mitarbeiter?.Intern_telefon2,
      MobilBusiness: mitarbeiter?.Mobil_buiness,
      MobilPrivat: mitarbeiter?.Mobil_privat,

      Geburtstag: mitarbeiter?.Geburtstag
        ? new Date(mitarbeiter.Geburtstag)
        : undefined,
      Azubi: mitarbeiter?.Azubi ?? false,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (mitarbeiter == null) {
      const res = await Create("Mitarbeiter", values);
      if (res) {
        navigate("/Mitarbeiter");
      } else {
        alert("Fehler beim Speichern des Mitarbeiters");
        return;
      }
    }

    if (mitarbeiter?.id == null) return;
    const res = await Update("Mitarbeiter", values, mitarbeiter.id);
    if (res) {
      navigate("/Mitarbeiter");
    } else {
      alert("Fehler beim Speichern des Mitarbeiters");
      return;
    }
  };

  return (
    <div className="container max-w-md mx-auto mt-10">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <div className="grid grid-cols-4 gap-4">
            <div className="col-span-3">
              <FormField
                control={form.control}
                name="Name"
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
            </div>
            <div className="col-span-1">
              <FormField
                control={form.control}
                name="Short"
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
            </div>
          </div>
          <FormField
            control={form.control}
            name="Email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input type="email" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="grid grid-cols-4 gap-4">
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
              name="InternTelefon1"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Intern1</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="InternTelefon2"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Intern2</FormLabel>
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
                  <FormLabel>Home Office</FormLabel>
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
              name="FestnetzPrivat"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Festnetz Privat</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="FestnetzBusiness"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Festnetz Business</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="MobilPrivat"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Mobil Privat</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="MobilBusiness"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Mobil Business</FormLabel>
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
              name="Azubi"
              render={({ field }) => (
                <FormItem className="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4 shadow">
                  <FormControl>
                    <Checkbox
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                  <div className="space-y-1 leading-none">
                    <FormLabel>Azubi</FormLabel>
                  </div>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="Geburtstag"
              render={({ field }) => (
                <FormItem className="flex flex-col">
                  <FormLabel>Geburtstag</FormLabel>
                  <Popover>
                    <PopoverTrigger asChild>
                      <FormControl>
                        <Button
                          variant={"outline"}
                          className={cn(
                            "w-[240px] pl-3 text-left font-normal",
                            !field.value && "text-muted-foreground"
                          )}
                        >
                          {field.value ? (
                            format(field.value, "PPP", {
                              locale: de,
                            })
                          ) : (
                            <span>Pick a date</span>
                          )}
                          <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                        </Button>
                      </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-auto p-0" align="start">
                      <Calendar
                        mode="single"
                        locale={de}
                        fixedWeeks
                        showWeekNumber
                        selected={field.value}
                        onSelect={field.onChange}
                        initialFocus
                      />
                    </PopoverContent>
                  </Popover>

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
