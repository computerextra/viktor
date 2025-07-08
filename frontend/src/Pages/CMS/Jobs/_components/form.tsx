import {
  type JobReponse,
  CreateJob,
  GetJob,
  JobProps,
  UpdateJob,
} from "@/api/Jobs";
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
import { Switch } from "@/components/ui/switch";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import type z from "zod";

export default function JobForm({ id }: { id?: string }) {
  const navigate = useNavigate();
  const [Job, setJob] = useState<JobReponse | undefined>(undefined);
  const [loading, setLoading] = useState(false);
  const form = useForm<z.infer<typeof JobProps>>({
    resolver: zodResolver(JobProps),
    defaultValues: {
      name: Job?.name,
      online: Job?.online,
    },
  });

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const res = await GetJob(id);
      setJob(res);
      setLoading(false);
    })();
  }, [id]);

  useEffect(() => {
    if (Job == null) return;
    form.reset({
      name: Job.name,
      online: Job.online,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [Job]);

  const onSubmit = async (values: z.infer<typeof JobProps>) => {
    if (id != null && Job != null) {
      await UpdateJob(id, values);
    } else {
      await CreateJob(values);
    }
    await navigate("/CMS/Jobs");
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
          <FormField
            control={form.control}
            name="online"
            render={({ field }) => (
              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
                <div className="space-y-0.5">
                  <FormLabel>Auf Webseite anzeigen?</FormLabel>
                  <FormDescription>
                    Wenn aktiviert, wird der Job auf der Webseite angezeigt
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
          <Button type="submit">Speichern</Button>
        </form>
      </Form>
    </div>
  );
}
