import z from "zod";
import { client, config } from "..";

const JobReponse = z.object({
  id: z.string(),
  name: z.string(),
  online: z.boolean(),
});

export type JobReponse = z.infer<typeof JobReponse>;

export const JobProps = z.object({
  name: z.string(),
  online: z.boolean().default(false).optional(),
});

const GetJobs = async () => {
  const res = await client.get<JobReponse[]>("/Job", config);
  return res.data ?? null;
};

const GetJob = async (id: string) => {
  const res = await client.get<JobReponse>("/Job" + id, config);
  return res.data ?? null;
};

const ToggleJob = async (id: string) => {
  const res = await client.post(`/Job/${id}/toggle`, {}, config);
  return res.data ?? null;
};

const CreateJob = async (props: z.infer<typeof JobProps>) => {
  const data = new FormData();
  data.append("name", props.name);
  const res = await client.post("/Job", data, config);
  return res.data ?? null;
};

const UpdateJob = async (id: string, props: z.infer<typeof JobProps>) => {
  const data = new FormData();
  data.append("name", props.name);
  const res = await client.post("/Job/" + id, data, config);
  return res.data ?? null;
};

const DeleteJob = async (id: string) => {
  const res = await client.delete("/Job/" + id, config);
  return res.data ?? null;
};

export { CreateJob, DeleteJob, GetJob, GetJobs, ToggleJob, UpdateJob };
