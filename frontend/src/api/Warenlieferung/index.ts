import { client, config } from "..";

export const GenerateWarenlieferung = async () => {
  const res = await client.post("Warenlieferung/Generate", {}, config);
  return res.data ?? null;
};

export const SendWarenlieferung = async () => {
  const res = await client.post("Warenlieferung/Send", {}, config);
  return res.data ?? null;
};
