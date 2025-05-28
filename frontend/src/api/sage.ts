import { GetKundeWithKundennummer, SearchSage } from "@wails/go/main/App";
import type { sagedb } from "@wails/go/models";

export const Get = async (
  kundennummer: string
): Promise<sagedb.User | undefined> => {
  const user = await GetKundeWithKundennummer(kundennummer);
  return user;
};

export const Search = async (
  searchTerm: string
): Promise<Array<sagedb.SearchResult>> => {
  const results = await SearchSage(searchTerm);
  return results;
};
