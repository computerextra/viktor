import axios, {
  type AxiosRequestConfig,
  type RawAxiosRequestHeaders,
} from "axios";

export const baseURL = "http://localhost:3000/api";

export const client = axios.create({
  baseURL: baseURL,
});

export const config: AxiosRequestConfig = {
  headers: {
    Accept: "application/json",
    "Content-Type": "multipart/form-data",
  } as RawAxiosRequestHeaders,
};
