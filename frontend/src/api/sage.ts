import {
  GetKundeWithKundennummer,
  SearchSage,
} from "@bindings/viktor/backend/app";
import type { SearchResult, User } from "@bindings/viktor/sagedb/models";

export const Get = async (kundennummer: string): Promise<User | null> => {
  const user = await GetKundeWithKundennummer(kundennummer);
  return user;
};

export const Search = async (
  searchTerm: string
): Promise<Array<SearchResult>> => {
  const results = await SearchSage(searchTerm);
  return results;
};
