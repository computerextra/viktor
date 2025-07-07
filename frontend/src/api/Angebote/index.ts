import z from "zod";
import { client, config } from "..";

const AngeboteRes = z.object({
  id: z.string(),
  title: z.string(),
  subtitle: z.string(),
  date_start: z.date(),
  date_stop: z.date(),
  link: z.string(),
  image: z.string(),
  anzeigen: z.boolean(),
});

export const AngeboteProps = z.object({
  title: z.string(),
  subtitle: z.string().optional(),
  image: z.string(),
  link: z.string(),
  date_start: z.date(),
  date_stop: z.date(),
  anzeigen: z.boolean(),
});

export type AngeboteRes = z.infer<typeof AngeboteRes>;

const GetAngebote = async () => {
  const res = await client.get<AngeboteRes[]>("/Angebot", config);
  return res.data ?? null;
};

const GetAngebot = async (id: string) => {
  const res = await client.get<AngeboteRes>("/Angebot/" + id, config);
  return res.data ?? null;
};
// 02.01.2006
const CreateAngebot = async (props: z.infer<typeof AngeboteProps>) => {
  const data = new FormData();
  data.append("title", props.title);
  if (props.subtitle) data.append("subtitle", props.subtitle);
  data.append("image", props.image);
  data.append("link", props.link);
  data.append(
    "date_start",
    props.date_start.toLocaleString("de-de", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
    })
  );
  data.append(
    "date_stop",
    props.date_stop.toLocaleString("de-de", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
    })
  );
  data.append("anzeigen", props.anzeigen ? "true" : "false");

  const res = await client.post("/Angebot", data, config);
  return res.data ?? null;
};

const UpdateAngebot = async (
  id: string,
  props: z.infer<typeof AngeboteProps>
) => {
  const data = new FormData();
  data.append("title", props.title);
  if (props.subtitle) data.append("subtitle", props.subtitle);
  data.append("image", props.image);
  data.append("link", props.link);
  data.append(
    "date_start",
    props.date_start.toLocaleString("de-de", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
    })
  );
  data.append(
    "date_stop",
    props.date_stop.toLocaleString("de-de", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
    })
  );
  data.append("anzeigen", props.anzeigen ? "true" : "false");

  const res = await client.post("/Angebot/" + id, data, config);
  return res.data ?? null;
};

const ToggleAngebot = async (id: string) => {
  const res = await client.post(`/Angebot/${id}/toggle`, {}, config);
  return res.data ?? null;
};

const DeleteAngebote = async (id: string) => {
  const res = await client.delete("/Angebot/" + id, config);
  return res.data ?? null;
};

export {
  CreateAngebot,
  DeleteAngebote,
  GetAngebot,
  GetAngebote,
  ToggleAngebot,
  UpdateAngebot,
};
