import z from "zod";
import { client, config } from "..";

const ArchiveResponse = z.object({
  id: z.number().int(),
  title: z.string(),
});

export type ArchiveResponse = z.infer<typeof ArchiveResponse>;

export const ArchiveProps = z.object({
  search: z.string(),
});

const SearchArchive = async (props: z.infer<typeof ArchiveProps>) => {
  const data = new FormData();
  data.append("search", props.search);
  const res = await client.post<ArchiveResponse[]>("/Archiv", data, config);
  return res.data ?? null;
};

export { SearchArchive };
