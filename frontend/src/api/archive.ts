import type { ArchiveResult } from "@bindings/viktor/archive";
import { Get, SearchArchive } from "@bindings/viktor/backend/app";

export const Search = async (
  searchTerm: string
): Promise<Array<ArchiveResult>> => {
  const results = await SearchArchive(searchTerm);
  return results;
};

export const Download = async (id: number): Promise<boolean> => {
  const response = await Get(id);
  return response;
};
