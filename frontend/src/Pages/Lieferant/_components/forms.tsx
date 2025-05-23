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
import {
  Ansprechpartner,
  AnsprechpartnerParams,
  Lieferant,
  LiefertantenParams,
} from "@api/db";
import { zodResolver } from "@hookform/resolvers/zod";
import type { db } from "@wails/go/models";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";

const LieferatenFormSchema = LiefertantenParams;
const AnsprechpartnerFornSchema = AnsprechpartnerParams;

export const LieferantenForm = ({
  lieferant,
}: {
  lieferant?: db.Lieferant;
}) => {
  const navigate = useNavigate();

  const form = useForm<z.infer<typeof LieferatenFormSchema>>({
    resolver: zodResolver(LieferatenFormSchema),
    defaultValues: {
      Firma: lieferant?.Firma ? lieferant.Firma : "",
      Kundennummer: lieferant?.Kundennummer ? lieferant.Kundennummer : "",
      Webseite: lieferant?.Webseite ? lieferant.Webseite : "",
    },
  });

  const onSubmit = async (values: z.infer<typeof LieferatenFormSchema>) => {
    if (lieferant) {
      await Lieferant.Update(lieferant.ID, values);
      await navigate("/Lieferant");
    } else {
      await Lieferant.Create(values);
      await navigate("/Lieferant");
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="mt-12 space-y-4">
        <FormField
          control={form.control}
          name="Firma"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Firma</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="Kundennummer"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Kundennummer</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="Webseite"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Webseite</FormLabel>
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
  );
};

export const AnsprechpartnerForm = ({
  ansprechpartner,
  LieferantenId,
}: {
  ansprechpartner?: db.Ansprechpartner;
  LieferantenId: number;
}) => {
  const navigate = useNavigate();

  const form = useForm<z.infer<typeof AnsprechpartnerFornSchema>>({
    resolver: zodResolver(AnsprechpartnerFornSchema),
    defaultValues: {
      Mail: ansprechpartner?.Mail ? ansprechpartner.Mail : "",
      Mobil: ansprechpartner?.Mobil ? ansprechpartner.Mobil : "",
      Name: ansprechpartner?.Name ? ansprechpartner.Name : "",
      Telefon: ansprechpartner?.Telefon ? ansprechpartner.Telefon : "",
      LieferantenId: LieferantenId,
    },
  });

  const onSubmit = async (
    values: z.infer<typeof AnsprechpartnerFornSchema>
  ) => {
    if (ansprechpartner) {
      await Ansprechpartner.Update(ansprechpartner.ID, values);
      await navigate("/Lieferant");
    } else {
      await Ansprechpartner.Create(values);
      await navigate("/Lieferant");
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="mt-12 space-y-4">
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
        <FormField
          control={form.control}
          name="Mail"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Mail</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="Telefon"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Telefon</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="Mobil"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Mobil</FormLabel>
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
  );
};
