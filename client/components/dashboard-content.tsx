"use client";

import { useRouter } from "next/navigation";
import Link from "next/link";
import useSWR from "swr";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { Badge } from "@/components/ui/badge";
import {
  Plus,
  FileCode,
  Lock,
  Flame,
  Globe,
  Eye,
  Trash2,
  Calendar,
} from "lucide-react";
import { formatDistanceToNow } from "date-fns";
import { ApiBaseUrl, HandleDelete } from "@/lib/utils";
import { apiFetch } from "@/lib/api";
import { toast } from "sonner";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogOverlay,
  DialogTrigger,
} from "@radix-ui/react-dialog";
import { DialogFooter, DialogHeader } from "./ui/dialog";

const sessionFetcher = (url: string) => apiFetch(url).then((res) => res.json());

const pastesFetcher = (url: string) =>
  apiFetch(url)
    .then((res) => res.json())
    .then((data) => data.pastes || []);

interface Paste {
  id: string;
  title: string;
  language: string;
  createdAt: string;
  hasPassword: boolean;
  burnAfterRead: boolean;
  isPublic: boolean;
  views: number;
}

export function DashboardContent() {
  const router = useRouter();
  const [deleting, setDeleting] = useState(false);
  const [toBeDeleted, setToBeDeleted] = useState<string | null>(null);
  const [showLogoutDialog, setShowLogoutDialog] = useState(false);
  const { data: session, isLoading: sessionLoading } = useSWR(
    `${ApiBaseUrl()}/api/users/me`,
    sessionFetcher,
  );

  const {
    data: pastes,
    isLoading: pastesLoading,
    mutate,
  } = useSWR<Paste[]>(
    session?.user ? `${ApiBaseUrl()}/api/paste/mine` : null,
    pastesFetcher,
  );
  const handleDelete = async () => {
    setDeleting(true);
    try {
      const res = await HandleDelete(toBeDeleted!);
      if (res == 1) {
        mutate(
          pastes?.filter((paste) => paste.id !== toBeDeleted),
          false, // don't revalidate
        );

        setShowLogoutDialog(false);
      } else {
        setDeleting(false);
        toast.error("You are not authorized to delete this paste.");
      }
    } catch {
      setDeleting(false);
    }
  };

  if (sessionLoading) {
    return (
      <div className="flex items-center justify-center py-12">
        <div className="h-8 w-8 animate-spin rounded-full border-2 border-primary border-t-transparent" />
      </div>
    );
  }

  if (!session?.user) {
    router.push("/login");
    return null;
  }

  return (
    <div className="space-y-6 rounded-sm">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-foreground">My Pastes</h1>
        </div>
      </div>

      {pastesLoading ? (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {[...Array(6)].map((_, i) => (
            <div
              key={i}
              className="h-32 animate-pulse rounded-lg bg-card border border-border"
            />
          ))}
        </div>
      ) : !pastes || pastes.length === 0 ? (
        <Card className="border-border bg-card">
          <CardContent className="py-12 text-center">
            <FileCode className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
            <h2 className="text-xl font-medium text-foreground mb-2">
              No pastes yet
            </h2>
            <p className="text-muted-foreground mb-4">
              Create your first paste to get started
            </p>
            <Button asChild>
              <Link href="/">
                <Plus className="mr-2 h-4 w-4" />
                Create Paste
              </Link>
            </Button>
          </CardContent>
        </Card>
      ) : (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {pastes.map((paste) => (
            <Card
              key={paste.id}
              className="border-border bg-card group hover:border-primary/50 transition-colors"
            >
              <CardHeader className="pb-2">
                <div className="flex items-start justify-between">
                  <CardTitle className="text-base font-medium text-foreground truncate pr-2">
                    {paste.title || "Untitled"}
                  </CardTitle>
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-8 w-8 opacity-0 group-hover:opacity-100 transition-opacity text-destructive hover:text-destructive"
                    onClick={() => {
                      setToBeDeleted(paste.id);
                      setShowLogoutDialog(true);
                    }}
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </div>
              </CardHeader>
              <CardContent>
                <Link href={`/p/${paste.id}`} className="block">
                  <div className="flex items-center gap-2 mb-3">
                    <Badge variant="outline" className="text-xs">
                      {paste.language}
                    </Badge>
                    {paste.hasPassword && (
                      <Lock className="h-3.5 w-3.5 text-muted-foreground" />
                    )}
                    {paste.burnAfterRead && (
                      <Flame className="h-3.5 w-3.5 text-destructive" />
                    )}
                    {paste.isPublic ? (
                      <Globe className="h-3.5 w-3.5 text-primary" />
                    ) : (
                      <Lock className="h-3.5 w-3.5 text-muted-foreground" />
                    )}
                  </div>
                  <div className="flex items-center justify-between text-xs text-muted-foreground">
                    <span className="flex items-center gap-1">
                      <Calendar className="h-3 w-3" />
                      {formatDistanceToNow(new Date(paste.createdAt), {
                        addSuffix: true,
                      })}
                    </span>
                    <span className="flex items-center gap-1">
                      <Eye className="h-3 w-3" />
                      {paste.views}
                    </span>
                  </div>
                </Link>
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      <Dialog open={showLogoutDialog} onOpenChange={setShowLogoutDialog}>
        <DialogTrigger asChild></DialogTrigger>

        <DialogOverlay className="fixed z-20 inset-0 bg-black/50" />

        <DialogContent className="fixed top-1/2 left-1/2 z-50 w-[90%] max-w-md -translate-x-1/2 -translate-y-1/2 rounded-none bg-background p-6 shadow-lg text-foreground">
          <DialogHeader>
            <DialogHeader>Are you sure?</DialogHeader>
            <DialogDescription>
              Are you sure you want to Delete?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <DialogClose asChild>
              <Button
                variant="outline"
                className="text-foreground"
                disabled={deleting}
              >
                Cancel
              </Button>
            </DialogClose>
            <Button
              variant="destructive"
              onClick={handleDelete}
              disabled={deleting}
            >
              {deleting ? " Deleting..." : "Yes, Delete "}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
