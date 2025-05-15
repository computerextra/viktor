import type { archive } from "@wails/go/models";
import { SearchArchive, DownloadArchive } from "@wails/go/main/App";

const Search = async (
  searchTerm: string
): Promise<Array<archive.ArchiveResult>> => {
  const results = await SearchArchive(searchTerm);
  return results;
};

const Download = async (id: number): Promise<boolean> => {
  const response = await DownloadArchive(id);
  return response;
};

export default { Search, Download };
