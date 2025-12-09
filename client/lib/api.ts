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

  let res = await fetch(url, fetchOptions);

  if (res.status === 401) {
    const data = await res
      .clone()
      .json()
      .catch(() => null);

    if (data?.error === "token_invalid" || data?.error === "token_expired") {
      const refresh = await fetch(`${ApiBaseUrl()}/api/auth/refresh`, {
        method: "POST",
        credentials: "include",
      });

      if (refresh.ok) {
        res = await fetch(url, fetchOptions);
      }
    }
  }

  return res;
}
