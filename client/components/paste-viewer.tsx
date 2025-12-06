"use client";

import type React from "react";

import { useState } from "react";
import { useRouter } from "next/navigation";
import useSWR from "swr";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { SyntaxHighlighter } from "@/components/syntax-highlighter";
import {
  Copy,
  Check,
  Lock,
  Flame,
  Eye,
  Calendar,
  Download,
  Trash2,
  AlertTriangle,
  Trash,
  Delete,
} from "lucide-react";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogOverlay,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { formatDistanceToNow } from "date-fns";
import { ApiBaseUrl, HandleDelete } from "@/lib/utils";
import { apiFetch } from "@/lib/api";
import { LanguageExtensions } from "@/lib/languages";
import { Toaster } from "./ui/sonner";
import { toast } from "sonner";

interface PasteViewerProps {
  id: string;
}

interface Paste {
  id: string;
  title: string;
  content: string;
  language: string;
  createdAt: string;
  views: number;
  burnAfterRead: boolean;
  expiresAt?: string;
  burned?: boolean;
  requiresPassword?: boolean;
}

const createFetcher = (password?: string) => async (url: string) => {
  const res = await apiFetch(url);

  const data = await res.json();
  if (!res.ok) {
    const error = new Error(data.error || "Failed to fetch");
    (error as Error & { info: typeof data }).info = data;
    throw error;
  }
  console.log("data", data.paste)
  return data.paste;
};

export function PasteViewer({ id }: PasteViewerProps) {
  const router = useRouter();
  const [password, setPassword] = useState("");
  const [submittedPassword, setSubmittedPassword] = useState<string>();
  const [copied, setCopied] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [showLogoutDialog, setShowLogoutDialog] = useState(false);

  const {
    data: paste,
    error,
    isLoading,
    mutate,
  } = useSWR<Paste>(`${ApiBaseUrl()}/api/paste/${id}`, createFetcher(submittedPassword), {
    revalidateOnFocus: false,
  });

  const handlePasswordSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setSubmittedPassword(password);
    mutate();
  };

  const handleCopy = async () => {
    if (paste?.content) {
      await navigator.clipboard.writeText(paste.content);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

const handleDownload = () => {
  if (paste) {
    const ext = LanguageExtensions[paste.language] || "txt";
    const blob = new Blob([paste.content], { type: "text/plain" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `${paste.title || "paste"}.${ext}`;
    a.click();
    URL.revokeObjectURL(url);
  }
};

  const handleDelete = async () => {
    setDeleting(true);
    try {
      const res = await HandleDelete(id);
      if(res == 1){
        router.push("/");
      }
      else {
        setDeleting(false);
        toast.error("You are not authorized to delete this paste.");
      }
    } catch {
      setDeleting(false);
    }
  };

  if (isLoading) {
    return (
      <Card className="border-border bg-card">
        <CardContent className="py-12">
          <div className="flex items-center justify-center">
            <div className="h-8 w-8 animate-spin rounded-none border-2 border-primary border-t-transparent" />
          </div>
        </CardContent>
      </Card>
    );
  }

  const errorInfo = error
    ? (
        error as Error & {
          info?: {
            requiresPassword?: boolean;
            title?: string;
            language?: string;
          };
        }
      ).info
    : null;

  if (errorInfo?.requiresPassword) {
    return (
      <Card className="border-border bg-card">
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-foreground">
            <Lock className="h-5 w-5 text-primary" />
            Password Protected
          </CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground mb-4">
            This paste is protected. Enter the password to view it.
          </p>
          <form onSubmit={handlePasswordSubmit} className="flex gap-2">
            <Input
              type="password"
              placeholder="Enter password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="bg-secondary border-border"
            />
            <Button type="submit">Unlock</Button>
          </form>
          {submittedPassword && error && (
            <p className="text-destructive text-sm mt-2">Invalid password</p>
          )}
        </CardContent>
      </Card>
    );
  }

  if (error) {
    return (
      <Card className="border-border bg-card">
        <CardContent className="py-12 text-center">
          <h2 className="text-xl font-bold text-foreground mb-2">
            Paste Not Found
          </h2>
          <p className="text-muted-foreground mb-4">
            This paste may have been deleted, expired, or burned after reading.
          </p>
          <Button onClick={() => router.push("/")}>Create New Paste</Button>
        </CardContent>
      </Card>
    );
  }

  if (!paste) return null;

  const lines = paste.content.split("\n");

  return (
    <div className="space-y-4">
      {paste.burned && (
        <Card className="border-destructive bg-destructive/10">
          <CardContent className="py-4">
            <div className="flex items-center gap-2 text-destructive">
              <Delete className="h-5 w-5" />
              <span className="font-medium">
                This paste has been burned and will no longer be accessible.
              </span>
            </div>
          </CardContent>
        </Card>
      )}

      <Card className="border-border bg-card">
        <CardHeader className="border-b border-border pb-4">
          <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
            <div>
              <CardTitle className="text-xl font-bold text-foreground">
                {paste.title || "Untitled"}
              </CardTitle>
              <div className="flex flex-wrap items-center gap-2 mt-2 text-sm text-muted-foreground">
                <Badge variant="outline">{paste.language}</Badge>
                <span className="flex items-center gap-1">
                  <Calendar className="h-3.5 w-3.5" />
                  {formatDistanceToNow(new Date(paste.createdAt), {
                    addSuffix: true,
                  })}
                </span>
                <span className="flex items-center gap-1">
                  <Eye className="h-3.5 w-3.5" />
                  {paste.views} view{paste.views !== 1 ? "s" : ""}
                </span>
                {paste.burnAfterRead && (
                  <span className="flex items-center gap-1 text-destructive">
                    <Flame className="h-3.5 w-3.5" />
                    Burns after read
                  </span>
                )}
              </div>
            </div>
            <div className="flex items-center gap-2">
              <Button variant="outline" size="sm" onClick={handleCopy}>
                {copied ? (
                  <Check className="h-4 w-4" />
                ) : (
                  <Copy className="h-4 w-4" />
                )}
                <span className="ml-1 hidden sm:inline">
                  {copied ? "Copied!" : "Copy"}
                </span>
              </Button>
              <Button variant="outline" size="sm" onClick={handleDownload}>
                <Download className="h-4 w-4" />
                <span className="ml-1 hidden sm:inline">Download</span>
              </Button>
              <Button
                variant="outline"
                size="sm"
                onClick={() => {
                  setShowLogoutDialog(true);
                }}
                disabled={deleting}
                className="text-destructive hover:text-destructive bg-transparent"
              >
                <Trash2 className="h-4 w-4" />
              </Button>
            </div>
          </div>
        </CardHeader>
        <CardContent className="p-0">
          <div className="flex">
            <div className="shrink-0 select-none border-r border-border bg-secondary/50 px-3 py-4 text-right font-mono text-xs text-muted-foreground">
              {lines.map((_, i) => (
                <div key={i} className="leading-6">
                  {i + 1}
                </div>
              ))}
            </div>
            <div className="flex-1 overflow-x-auto p-4 font-mono text-sm">
              <SyntaxHighlighter
                code={paste.content}
                language={paste.language}
              />
            </div>
          </div>
        </CardContent>
      </Card>
      <Dialog open={showLogoutDialog} onOpenChange={setShowLogoutDialog}>
        <DialogTrigger asChild></DialogTrigger>

        <DialogOverlay className="fixed z-20 inset-0 bg-black/50" />

        <DialogContent className="fixed top-1/2 left-1/2 z-50 w-[90%] max-w-md -translate-x-1/2 -translate-y-1/2 rounded-none bg-background p-6 shadow-lg text-foreground">
          <DialogHeader>
            <DialogTitle>Are you sure?</DialogTitle>
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
