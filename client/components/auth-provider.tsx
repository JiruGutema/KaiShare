"use client";

import {
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";
import { IsLoggedIn } from "@/lib/utils";

const AuthContext = createContext<{
  loggedIn: boolean;
  authLoading: boolean;
  refresh: () => Promise<void>;
}>({
  loggedIn: false,
  authLoading: true,
  refresh: async () => {},
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [loggedIn, setLoggedIn] = useState(false);
  const [authLoading, setAuthLoading] = useState(true);

  async function refresh() {
    const ok = await IsLoggedIn();
    setLoggedIn(ok);
    setAuthLoading(false);
  }

  useEffect(() => {
    refresh();
  }, []);

  return (
    <AuthContext.Provider value={{ loggedIn, authLoading, refresh }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}
