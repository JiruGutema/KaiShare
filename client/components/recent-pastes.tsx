"use client";

import Link from "next/link";
import useSWR from "swr";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Lock, Flame, FileCode } from "lucide-react";
import { formatDistanceToNow } from "date-fns";
import { ApiBaseUrl } from "@/lib/utils";
import { apiFetch } from "@/lib/api";

const fetcher = (url: string) =>
  apiFetch(url)
    .then((res) => res.json())
    .then((data) => data.pastes);

interface Paste {
  id: string;
  title: string;
  language: string;
  createdAt: string;
  hasPassword: boolean;
  burnAfterRead: boolean;
}

export function RecentPastes() {
  const { data: pastes, isLoading } = useSWR(`${ApiBaseUrl()}/api/users/me`, fetcher, {
  revalidateOnFocus: false,
  revalidateOnReconnect: false,
  dedupingInterval: Infinity,
});
  if (isLoading) {
    return (
      <Card className="border-border bg-card">
        <CardHeader>
          <CardTitle className="text-lg font-medium text-foreground flex items-center gap-2">
            Your Recent Pastes
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-2 ">
            {[...Array(5)].map((_, i) => (
              <div
                key={i}
                className="h-12 animate-pulse rounded-none bg-secondary"
              />
            ))}
          </div>
        </CardContent>
      </Card>
    );
  }

  if (!pastes || pastes.length === 0) {
    return (
      <Card className="border-border bg-card">
        <CardHeader>
          <CardTitle className="text-lg font-medium text-foreground flex items-center gap-2">
            Your Recent Pastes
          </CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-sm text-muted-foreground text-center py-8">
            No your public pastes yet. Create one.
          </p>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card className="border-border bg-card">
      <CardHeader>
        <CardTitle className="text-lg font-medium text-foreground flex items-center gap-2">
          Recent Public Pastes
        </CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-2 rounded-none">
          {pastes.map((paste) => (
            <Link
              key={paste.id}
              href={`/p/${paste.id}`}
              className="flex items-center justify-between rounded-none border border-border bg-secondary p-3 transition-colors hover:bg-secondary/80 hover:border-primary/50"
            >
              <div className="flex items-center gap-3 min-w-0">
                <FileCode className="h-4 w-4 text-muted-foreground shrink-0" />
                <div className="min-w-0">
                  <p className="text-sm font-medium text-foreground truncate">
                    {paste.title || "Untitled"}
                  </p>
                  <p className="text-xs text-muted-foreground">
                    {formatDistanceToNow(new Date(paste.createdAt), {
                      addSuffix: true,
                    })}
                  </p>
                </div>
              </div>
              <div className="flex items-center gap-2 shrink-0">
                {paste.hasPassword && (
                  <Lock className="h-3.5 w-3.5 text-muted-foreground" />
                )}
                {paste.burnAfterRead && (
                  <Flame className="h-3.5 w-3.5 text-destructive" />
                )}
                <Badge variant="outline" className="text-xs">
                  {paste.language}
                </Badge>
              </div>
            </Link>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}
