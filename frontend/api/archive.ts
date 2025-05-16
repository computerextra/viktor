import { DownloadArchive, SearchArchive } from "@wails/go/main/App";
import type { archive } from "@wails/go/models";

export const Search = async (
  searchTerm: string
): Promise<Array<archive.ArchiveResult>> => {
  const results = await SearchArchive(searchTerm);
  return results;
};

export const Download = async (id: number): Promise<boolean> => {
  const response = await DownloadArchive(id);
  return response;
};
