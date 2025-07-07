import z from "zod";
import { client, config } from "..";

const PartnerRes = z.object({
  id: z.string(),
  name: z.string(),
  link: z.string(),
  image: z.string(),
});
type PartnerRes = z.infer<typeof PartnerRes>;

export const PartnerProps = z.object({
  name: z.string(),
  image: z.string(),
  link: z.string(),
});

const GetPartners = async () => {
  const res = await client.get<PartnerRes[]>("/Partner", config);
  return res.data ?? null;
};

const GetPartner = async (id: string) => {
  const res = await client.get<PartnerRes>("/Partner/" + id, config);
  return res.data ?? null;
};

const CreatePartner = async (props: z.infer<typeof PartnerProps>) => {
  const data = new FormData();
  data.append("name", props.name);
  data.append("image", props.image);
  data.append("link", props.link);

  const res = await client.post("/Partner", data, config);
  return res.data ?? null;
};

const UpdatePartner = async (
  id: string,
  props: z.infer<typeof PartnerProps>
) => {
  const data = new FormData();
  data.append("name", props.name);
  data.append("image", props.image);
  data.append("link", props.link);

  const res = await client.post("/Partner/" + id, data, config);
  return res.data ?? null;
};

const DeletePartner = async (id: string) => {
  const res = await client.delete("/Partner/" + id, config);
  return res.data ?? null;
};

export { CreatePartner, DeletePartner, GetPartner, GetPartners, UpdatePartner };
