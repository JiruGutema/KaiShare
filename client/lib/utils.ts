import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { apiFetch } from "./api";

const isDevelopment = process.env.NEXT_PUBLIC_NODE_ENV === "development";
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export async function HandleDelete(id: string) {
  try {
   const res =  await apiFetch(`${ApiBaseUrl()}/api/paste/${id}`, { method: "DELETE" });
    if(res.ok){
      return 1
    }
    else {
      return 0
    }
  } catch {
    return 0
  }
}

export const Logger = {
  log: (...inputs: any[]) => {
    if (isDevelopment) {
      console.log(...inputs);
    }
  },
  error: (...inputs: any[]) => {
    if (isDevelopment) {
      console.error(...inputs);
    }
  },
  warn: (...inputs: any[]) => {
    if (isDevelopment) {
      console.warn(...inputs);
    }
  },
};

export  const ApiBaseUrl = () => {
  return  process.env.NEXT_PUBLIC_SERVER_BASE_URL || "";
};
