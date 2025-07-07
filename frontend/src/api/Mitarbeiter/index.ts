import z from "zod";
import { client, config } from "..";

const MitarbeiterRes = z.object({
  id: z.string(),
  name: z.string(),
  short: z.string().optional(),
  image: z.boolean(),
  sex: z.string(),
  focus: z.string().optional(),
  mail: z.string().email().optional(),
  abteilungId: z.string().optional(),
  Azubi: z.boolean(),
  Geburtstag: z.date().optional(),
  Gruppenwahl: z.string().optional(),
  Mobil_Business: z.string().optional(),
  Mobil_Privat: z.string().optional(),
  Telefon_Business: z.string().optional(),
});

type MitarbeiterRes = z.infer<typeof MitarbeiterRes>;
type MitarbeiterMitAbteilungRes = MitarbeiterRes & {
  Abteilung: {
    id: string;
    name: string;
  } | null;
};

export const MitarbeiterProps = z.object({
  name: z.string(),
  short: z.string().optional(),
  sex: z.string(),
  abteilungId: z.string().optional(),
  image: z.boolean().default(false).optional(),
  Azubi: z.boolean().default(false).optional(),
  focus: z.string().optional(),
  mail: z.string().optional(),
  Gruppenwahl: z.string().optional(),
  Geburtstag: z.date().optional(),
  HomeOffice: z.string().optional(),
  Mobil_Business: z.string().optional(),
  Mobil_Privat: z.string().optional(),
  Telefon_Business: z.string().optional(),
  Telefon_Intern_1: z.string().optional(),
  Telefon_Intern_2: z.string().optional(),
  Telefon_Privat: z.string().optional(),
});

const GetMitarbeiters = async () => {
  const res = await client.get<MitarbeiterRes[]>("/Mitarbeiter", config);
  return res.data ?? null;
};

const GetMitarbeiter = async (id: string) => {
  const res = await client.get<MitarbeiterRes>("/Mitarbeiter/" + id, config);
  return res.data ?? null;
};

const GetMitarbeitersMitAbteilung = async () => {
  const res = await client.get<MitarbeiterMitAbteilungRes[]>(
    "Mitarbeiter/Abteilung",
    config
  );
  return res.data ?? null;
};

const CreateMitarbeiter = async (props: z.infer<typeof MitarbeiterProps>) => {
  const data = new FormData();
  data.append("name", props.name);
  if (props.short) data.append("short", props.short);
  data.append("sex", props.sex);
  if (props.abteilungId) data.append("abteilungId", props.abteilungId);
  data.append("image", props.image ? "true" : "false");
  data.append("Azubi", props.Azubi ? "true" : "false");
  if (props.focus) data.append("focus", props.focus);
  if (props.mail) data.append("mail", props.mail);
  if (props.Gruppenwahl) data.append("Gruppenwahl", props.Gruppenwahl);
  if (props.Geburtstag)
    data.append(
      "Geburtstag",
      props.Geburtstag.toLocaleDateString("de-de", {
        day: "2-digit",
        month: "2-digit",
        year: "numeric",
      })
    );
  if (props.HomeOffice) data.append("HomeOffice", props.HomeOffice);
  if (props.Mobil_Business) data.append("Mobil_Business", props.Mobil_Business);
  if (props.Mobil_Privat) data.append("Mobil_Privat", props.Mobil_Privat);
  if (props.Telefon_Business)
    data.append("Telefon_Business", props.Telefon_Business);
  if (props.Telefon_Intern_1)
    data.append("Telefon_Intern_1", props.Telefon_Intern_1);
  if (props.Telefon_Intern_2)
    data.append("Telefon_Intern_2", props.Telefon_Intern_2);
  if (props.Telefon_Privat) data.append("Telefon_Privat", props.Telefon_Privat);

  const res = await client.post("/Mitarbeiter", data, config);
  return res.data ?? null;
};

const UpdateMitarbeiter = async (
  id: string,
  props: z.infer<typeof MitarbeiterProps>
) => {
  const data = new FormData();
  data.append("name", props.name);
  if (props.short) data.append("short", props.short);
  data.append("sex", props.sex);
  if (props.abteilungId) data.append("abteilungId", props.abteilungId);
  data.append("image", props.image ? "true" : "false");
  data.append("Azubi", props.Azubi ? "true" : "false");
  if (props.focus) data.append("focus", props.focus);
  if (props.mail) data.append("mail", props.mail);
  if (props.Gruppenwahl) data.append("Gruppenwahl", props.Gruppenwahl);
  if (props.Geburtstag)
    data.append(
      "Geburtstag",
      props.Geburtstag.toLocaleDateString("de-de", {
        day: "2-digit",
        month: "2-digit",
        year: "numeric",
      })
    );
  if (props.HomeOffice) data.append("HomeOffice", props.HomeOffice);
  if (props.Mobil_Business) data.append("Mobil_Business", props.Mobil_Business);
  if (props.Mobil_Privat) data.append("Mobil_Privat", props.Mobil_Privat);
  if (props.Telefon_Business)
    data.append("Telefon_Business", props.Telefon_Business);
  if (props.Telefon_Intern_1)
    data.append("Telefon_Intern_1", props.Telefon_Intern_1);
  if (props.Telefon_Intern_2)
    data.append("Telefon_Intern_2", props.Telefon_Intern_2);
  if (props.Telefon_Privat) data.append("Telefon_Privat", props.Telefon_Privat);

  const res = await client.post("/Mitarbeiter/" + id, data, config);
  return res.data ?? null;
};

const DeleteMitarbeiter = async (id: string) => {
  const res = await client.delete("/Mitarbeiter/" + id, config);
  return res.data ?? null;
};

export {
  CreateMitarbeiter,
  DeleteMitarbeiter,
  GetMitarbeiter,
  GetMitarbeiters,
  GetMitarbeitersMitAbteilung,
  UpdateMitarbeiter,
};
