import z from "zod";
import { client, config } from "..";

const GetCmsCounterResponse = z.object({
  Abteilungen: z.number().int(),
  Angebote: z.number().int(),
  Jobs: z.number().int(),
  Mitarbeiter: z.number().int(),
  Partner: z.number().int(),
});

type GetCmsCounterResponse = z.infer<typeof GetCmsCounterResponse>;

const GetCmsCounter = async (): Promise<GetCmsCounterResponse | null> => {
  const response = await client.get<GetCmsCounterResponse>("/cms", config);
  return response.data ?? null;
};

export { GetCmsCounter, type GetCmsCounterResponse };
