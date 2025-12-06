"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Plus, User, LogOut, MoonStar } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import useSWR from "swr";
import { apiFetch } from "@/lib/api";
import { ApiBaseUrl } from "@/lib/utils";

const fetcher = (url: string) => apiFetch(url).then((res) => res.json());

export function Header() {
  const pathname = usePathname();
  const { data: session, mutate } = useSWR(
    `${ApiBaseUrl()}/api/users/me`,
    fetcher,
  );

const handleLogout = async () => {
  // Call logout API
  await fetch(`${ApiBaseUrl()}/api/auth/logout`, { method: "POST", credentials: "include" });

  // Clear SWR cache for user session
  mutate(`${ApiBaseUrl()}/api/users/me`, false);

  // Optional: redirect to homepage/login
  window.location.href = "/";
};

  return (
    <header className="border-b border-border bg-card">
      <div className="mx-auto max-w-6xl px-4 py-3">
        <div className="flex items-center justify-between">
          <Link href="/" className="flex items-center gap-2 group">
            <span className="text-xl font-bold text-foreground group-hover:text-primary transition-colors">
              <span className="text-orange-600">Kai</span>
              <span className="text-sky-700">Share</span>
            </span>
          </Link>

          <nav className="flex items-center gap-2">
            {pathname !== "/" && (
              <Button asChild variant="ghost" size="sm">
                <Link href="/">
                  <Plus className="mr-1 h-4 w-4" />
                  New Paste
                </Link>
              </Button>
            )}

            {session?.user ? (
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" size="sm">
                    <User className="mr-1 h-4 w-4" />
                    {session.user.username}
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuItem asChild>
                    <Link href="/dashboard">
                      User Account
                    </Link>
                  </DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem onClick={handleLogout} className="text-orange-600 font-bold">
                    Logout
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            ) : (
              <div className="flex items-center gap-2">
                <Button asChild variant="ghost" size="sm">
                  <Link href="/login">Login</Link>
                </Button>
                <Button asChild size="sm">
                  <Link href="/register">Sign Up</Link>
                </Button>
              </div>
            )}
            <Button
              variant="outline"
              size="sm"
              className="ml-2 border-none hover:bg-transparent focus:ring-0"
              onClick={() => {
                if (typeof window !== "undefined") {
                  const isDark =
                    document.documentElement.classList.toggle("dark");
                  localStorage.setItem("theme", isDark ? "dark" : "light");
                }
              }}
            >
              <MoonStar className="h-4 w-4" />
            </Button>
          </nav>
        </div>
      </div>
    </header>
  );
}
