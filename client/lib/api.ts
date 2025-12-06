import { ApiBaseUrl, Logger } from "./utils";
export interface ApiOptions {
  method?: string;
  headers?: Record<string, string>;
  body?: any;
}

export async function apiFetch(
  endpoint: string,
  options: ApiOptions = {},
): Promise<Response> {
  const baseUrl = ApiBaseUrl() || process.env.NEXT_PUBLIC_API_URL || "";
  const url = endpoint.startsWith("http") ? endpoint : `${baseUrl}${endpoint}`;

  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...options.headers,
  };

  const fetchOptions: RequestInit = {
    method: options.method || "GET",
    headers,
    credentials: "include",
  };
  if (options.body) {
    fetchOptions.body = JSON.stringify(options.body);
  }

  const res = await fetch(url, fetchOptions);

  return res;
}
