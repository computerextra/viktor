import {
  DeleteEinkauf,
  EinkaufProps,
  GetEinkauf,
  SkipEinkauf,
  UpdateEinkauf,
  type EinkaufRes,
} from "@/api/Mitarbeiter";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
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
import { Separator } from "@/components/ui/separator";
import { Switch } from "@/components/ui/switch";
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { z } from "zod";

export default function EinkaufForm({ id }: { id?: string }) {
  const navigate = useNavigate();
  const [Einkauf, setEinkauf] = useState<EinkaufRes | undefined>(undefined);
  const [isLoading, setIsLoading] = useState(false);

  const form = useForm<z.infer<typeof EinkaufProps>>({
    resolver: zodResolver(EinkaufProps),
    defaultValues: {
      Abo: Einkauf?.Einkauf?.Abonniert,
      Dinge: Einkauf?.Einkauf?.Dinge,
      Geld: Einkauf?.Einkauf?.Geld,
      Pfand: Einkauf?.Einkauf?.Pfand,
      Paypal: Einkauf?.Einkauf?.Paypal,
    },
  });

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setIsLoading(true);
      const res = await GetEinkauf(id);
      setEinkauf(res);
      setIsLoading(false);
    })();
  }, [id]);

  useEffect(() => {
    if (Einkauf == null) return;
    form.reset({
      Abo: Einkauf.Einkauf.Abonniert,
      Dinge: Einkauf.Einkauf.Dinge,
      Geld: Einkauf.Einkauf.Geld,
      Pfand: Einkauf.Einkauf.Pfand,
      Paypal: Einkauf.Einkauf.Paypal,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [Einkauf]);

  if (isLoading) return <>Lädt ...</>;

  const onSubmit = async (values: z.infer<typeof EinkaufProps>) => {
    if (id == null) return;
    await UpdateEinkauf(id, values);
    await navigate("/Einkauf");
  };

  return (
    <>
      <div className="mt-12 flex justify-center">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <div className="grid grid-cols-2 gap-8">
              <FormField
                control={form.control}
                name="Geld"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Geld</FormLabel>
                    <FormControl>
                      <Input disabled={isLoading} {...field} />
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
                      <Input disabled={isLoading} {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
            <div>
              <div className="grid grid-cols-2 gap-8">
                <FormField
                  control={form.control}
                  name="Abo"
                  render={({ field }) => (
                    <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                      <div className="space-y-0.5">
                        <FormLabel>Abo?</FormLabel>
                        <FormDescription>
                          Diesen Einkauf Abonnieren? <br />
                          Der Einkauf wird dadurch <br />
                          jeden Tag Automatisch angezeigt.
                        </FormDescription>
                      </div>
                      <FormControl>
                        <Switch
                          disabled={isLoading}
                          checked={field.value}
                          onCheckedChange={field.onChange}
                        />
                      </FormControl>
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="Paypal"
                  render={({ field }) => (
                    <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                      <div className="space-y-0.5">
                        <FormLabel>Paypal</FormLabel>
                        <FormDescription>
                          Falls vorhanden, kann <br />
                          der Einkäufer eine <br />
                          Bezahlung per Paypal veranlassen
                        </FormDescription>
                      </div>
                      <FormControl>
                        <Switch
                          disabled={isLoading}
                          checked={field.value}
                          onCheckedChange={field.onChange}
                        />
                      </FormControl>
                    </FormItem>
                  )}
                />
              </div>
            </div>
            <FormField
              control={form.control}
              name="Dinge"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Dein Einkauf</FormLabel>
                  <FormControl>
                    <Textarea disabled={isLoading} {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div>
              <small className="text-center">
                Bilder Upload ist aktuell abgestellt.
              </small>
            </div>
            <Button disabled={isLoading} type="submit">
              Speichern
            </Button>
          </form>
        </Form>
      </div>

      {isLoading && (
        <p className="text-center text-6xl">Bitte warten, es lädt gerade</p>
      )}
      <div className="mx-auto mt-5 grid max-w-4xl grid-cols-3 gap-8"></div>
      <Separator className="my-8" />
      <div className="mx-auto grid max-w-[60%] grid-cols-2 gap-8">
        <AlertDialog>
          <AlertDialogTrigger asChild>
            <Button variant={"neutral"} disabled={isLoading}>
              Einkauf löschen
            </Button>
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Einkauf löschen</AlertDialogTitle>
              <AlertDialogDescription>
                Den Einkauf wirklich löschen? <br />
                Durch das löschen wird der Einkauf nicht mehr in der Liste
                angezeigt und muss neu eingegeben und gespeichert werden!
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Abbrechen</AlertDialogCancel>
              <AlertDialogAction
                onClick={async () => {
                  if (Einkauf == null) {
                    alert("Noch kein Einkauf für den Mitarbeiter gespeichert!");
                    return;
                  }
                  await DeleteEinkauf(Einkauf.id);
                  await navigate("/Einkauf");
                }}
              >
                Löschen
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
        <AlertDialog>
          <AlertDialogTrigger asChild>
            <Button variant={"neutral"} disabled={isLoading}>
              Einkauf auf morgen verschieben
            </Button>
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Einkauf Verschieben</AlertDialogTitle>
              <AlertDialogDescription>
                Hiermit wird der Einkauf auf den nächsten Tag verschoben. Der
                Einkauf wird morgen automatisch wieder angezeigt. Dies
                funktioniert auch, wenn der Einkauf Aboniert wurde.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Abbrechen</AlertDialogCancel>
              <AlertDialogAction
                onClick={async () => {
                  if (Einkauf == null) {
                    alert("Noch kein Einkauf für den Mitarbeiter gespeichert!");
                    return;
                  }
                  await SkipEinkauf(Einkauf.id);
                  await navigate("/Einkauf");
                }}
              >
                Überspringen
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
    </>
  );
}
