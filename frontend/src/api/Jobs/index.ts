import z from "zod";
import { client, config } from "..";

const JobReponse = z.object({
  id: z.string(),
  name: z.string(),
  online: z.boolean(),
});

type JobReponse = z.infer<typeof JobReponse>;

const GetJobs = async () => {
  const res = await client.get<JobReponse[]>("/Job", config);
  return res.data ?? null;
};

const GetJob = async (id: string) => {
  const res = await client.get<JobReponse>("/Job" + id, config);
  return res.data ?? null;
};

const CreateJob = async (name: string) => {
  const data = new FormData();
  data.append("name", name);
  const res = await client.post("/Job", data, config);
  return res.data ?? null;
};

const UpdateJob = async (id: string, name: string) => {
  const data = new FormData();
  data.append("name", name);
  const res = await client.post("/Job/" + id, data, config);
  return res.data ?? null;
};

const DeleteJob = async (id: string) => {
  const res = await client.delete("/Job/" + id, config);
  return res.data ?? null;
};

export { CreateJob, DeleteJob, GetJob, GetJobs, UpdateJob };
