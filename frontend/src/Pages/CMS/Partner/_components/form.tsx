import {
  CreatePartner,
  GetPartner,
  PartnerProps,
  type PartnerRes,
  UpdatePartner,
} from "@/api/Partner";
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
import { z } from "zod";

export default function PartnerForm({ id }: { id?: string }) {
  const navigate = useNavigate();
  const [Partner, setPartner] = useState<PartnerRes | undefined>(undefined);
  const [loading, setLoading] = useState(false);

  const form = useForm<z.infer<typeof PartnerProps>>({
    resolver: zodResolver(PartnerProps),
    defaultValues: {
      name: Partner?.name,
      link: Partner?.link,
      image: Partner?.image,
    },
  });

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const res = await GetPartner(id);
      setPartner(res);
      setLoading(false);
    })();
  }, [id]);

  useEffect(() => {
    if (Partner == null) return;
    form.reset({
      name: Partner?.name,
      link: Partner?.link,
      image: Partner?.image,
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [Partner]);

  const onSubmit = async (values: z.infer<typeof PartnerProps>) => {
    if (id != null && Partner != null) {
      await UpdatePartner(id, values);
    } else {
      await CreatePartner(values);
    }
    await navigate("/CMS/Partner");
  };

  if (loading) return <>Loading...</>;

  return (
    <div className="mt-12 flex justify-center">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
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
            name="link"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Link zum Partner</FormLabel>
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
                <FormLabel>Bildername wie auf Webseite</FormLabel>
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
