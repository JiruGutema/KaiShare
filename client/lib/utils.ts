import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { apiFetch } from "./api";
import { Session, User } from "./types";

const isDevelopment = process.env.NEXT_PUBLIC_NODE_ENV === "development";
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export async function HandleDelete(id: string) {
  try {
    const res = await apiFetch(`${ApiBaseUrl()}/api/paste/${id}`, {
      method: "DELETE",
    });
    if (res.ok) {
      return 1;
    } else {
      return 0;
    }
  } catch {
    return 0;
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

export function GetLocalUser(): Session | null {
  if (typeof window === "undefined") return null;
  const stored = localStorage.getItem("session");
  if (!stored) return null;

  try {
    const user: User = JSON.parse(stored);
    const session = {
      user: user,
    };
    return session;
  } catch (err) {
    console.error("Failed to parse user from localStorage", err);
    return null;
  }
}

export function setLocalUser(user: Session): void {
  if (typeof window === "undefined") return;

  try {
    localStorage.setItem("session", JSON.stringify(user));
  } catch (err) {
    console.error("Failed to save user to localStorage", err);
  }
}

export function removeLocalUser(): void {
  if (typeof window === "undefined") return;
  localStorage.removeItem("session");
}
export const ApiBaseUrl = () => {
  return process.env.NEXT_PUBLIC_SERVER_BASE_URL || "";
};

export async function IsLoggedIn() {
  try {
    const res = await apiFetch(`${ApiBaseUrl()}/api/auth/check`);

    if (!res.ok) {
      removeLocalUser();
      return false;
    }
    return true;
  } catch (error) {
    return false;
  }
}

export async function HandleLogout() {
  removeLocalUser();
  let loggedOut = false;
  if (await IsLoggedIn()) {
    const res = await apiFetch(`${ApiBaseUrl()}/api/auth/logout`, {
      method: "POST",
    });

    if (res.ok) {
      loggedOut = true;
    } else {
      loggedOut = false;
    }
  }

  if (loggedOut) {
    return true;
  } else {
    return false;
  }
}
