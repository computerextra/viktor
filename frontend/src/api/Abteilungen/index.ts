import z from "zod";
import { client, config } from "..";

const GetAbteilungeRes = z.object({
  id: z.string(),
  name: z.string(),
});

export type GetAbteilungeRes = z.infer<typeof GetAbteilungeRes>;

const GetAbteilungen = async () => {
  const res = await client.get<GetAbteilungeRes[]>("/Abteilung", config);
  return res.data ?? null;
};

const GetAbteilung = async (id: string) => {
  const res = await client.get<GetAbteilungeRes>("/Abteilung/" + id, config);
  return res.data ?? null;
};

const CreateAbteilung = async (name: string) => {
  const data = new FormData();
  data.append("name", name);
  const res = await client.post<GetAbteilungeRes>("/Abteilung", data, config);
  return res.data ?? null;
};

const UpdateAbteilung = async (id: string, name: string) => {
  const data = new FormData();
  data.append("name", name);
  const res = await client.post<GetAbteilungeRes>(
    "/Abteilung/" + id,
    data,
    config
  );
  return res.data ?? null;
};

const DeleteAbteilung = async (id: string) => {
  const res = await client.delete("/Abteilung" + id, config);
  return res.data ?? null;
};

export {
  CreateAbteilung,
  DeleteAbteilung,
  GetAbteilung,
  GetAbteilungen,
  UpdateAbteilung,
};
