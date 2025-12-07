"use client";

import type React from "react";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Switch } from "@/components/ui/switch";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { LANGUAGES, EXPIRATION_OPTIONS } from "@/lib/languages";
import { Lock, Flame, Clock, Globe, Eye, EyeOff, Send } from "lucide-react";
import { apiFetch } from "@/lib/api";
import { ApiBaseUrl, IsLoggedIn } from "@/lib/utils";
import { toast } from "sonner";

export function PasteForm() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState({
    title: "",
    content: "",
    language: "plaintext",
    password: "",
    burnAfterRead: false,
    expiration: "never",
    isPublic: false,
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData.content.trim()) return;

    setLoading(true);
    try {
      const res = await apiFetch(`${ApiBaseUrl()}/api/paste`, {
        method: "POST",
        body: formData,
      });

      if (res.ok) {
        const data = await res.json();
        toast.success("Paste created successfully!");
        router.push(`/p/${data.pasteId}`);
      } else {
        if (res.status == 500) {
          toast.error("Can't create paste. Internal server error");
        } else {
          toast.error("Failed to create paste. Please try again.");
        }
      }
    } catch (error) {
      console.error("Failed to create paste:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4 ">
      <Card className="border-border bg-card">
        <CardHeader className="pb-3">
          <CardTitle className="text-lg font-medium text-foreground">
            New Paste
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex flex-col sm:flex-row gap-4">
            <div className="flex-1">
              <Label htmlFor="title" className="text-muted-foreground text-sm">
                Title (optional)
              </Label>
              <Input
                id="title"
                placeholder="Untitled"
                value={formData.title}
                onChange={(e) =>
                  setFormData({ ...formData, title: e.target.value })
                }
                className="mt-1.5 bg-secondary border-border"
              />
            </div>
            <div className="w-full sm:w-48">
              <Label
                htmlFor="language"
                className="text-muted-foreground text-sm"
              >
                Syntax
              </Label>
              <Select
                value={formData.language}
                onValueChange={(value) =>
                  setFormData({ ...formData, language: value })
                }
              >
                <SelectTrigger
                  id="language"
                  className="mt-1.5 bg-secondary border-border"
                >
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  {LANGUAGES.map((lang) => (
                    <SelectItem key={lang.value} value={lang.value}>
                      {lang.label}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <div>
            <Label htmlFor="content" className="text-muted-foreground text-sm">
              Content
            </Label>
            <Textarea
              id="content"
              placeholder="Paste your code or text here..."
              value={formData.content}
              onChange={(e) =>
                setFormData({ ...formData, content: e.target.value })
              }
              className="mt-1.5 min-h-[300px] font-mono text-sm bg-secondary border-border resize-y"
              required
            />
          </div>
        </CardContent>
      </Card>

      <Card className="border-border bg-card">
        <CardHeader className="pb-3">
          <CardTitle className="text-lg font-medium text-foreground">
            Options
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid gap-6 sm:grid-cols-2">
            <div className="space-y-4">
              <div>
                <Label
                  htmlFor="password"
                  className="text-muted-foreground text-sm flex items-center gap-2"
                >
                  Password Protection
                </Label>
                <div className="relative mt-1.5">
                  <Input
                    id="password"
                    type={showPassword ? "text" : "password"}
                    placeholder="Leave empty for no password"
                    value={formData.password}
                    onChange={(e) =>
                      setFormData({ ...formData, password: e.target.value })
                    }
                    className="bg-secondary border-border pr-10"
                  />
                  <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
                  >
                    {showPassword ? (
                      <EyeOff className="h-4 w-4" />
                    ) : (
                      <Eye className="h-4 w-4" />
                    )}
                  </button>
                </div>
              </div>

              <div>
                <Label
                  htmlFor="expiration"
                  className="text-muted-foreground text-sm flex items-center gap-2"
                >
                  Expiration
                </Label>
                <Select
                  value={formData.expiration}
                  onValueChange={(value) =>
                    setFormData({ ...formData, expiration: value })
                  }
                >
                  <SelectTrigger
                    id="expiration"
                    className="mt-1.5 bg-secondary border-border"
                  >
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    {EXPIRATION_OPTIONS.map((opt) => (
                      <SelectItem key={opt.value} value={opt.value}>
                        {opt.label}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
            </div>

            <div className="space-y-4">
              <div className="flex items-center justify-between rounded-none border border-border bg-secondary p-3">
                <div className="flex items-center gap-2">
                  <div>
                    <p className="text-sm font-medium text-foreground">
                      Burn After Read
                    </p>
                    <p className="text-xs text-muted-foreground">
                      Delete after first view
                    </p>
                  </div>
                </div>
                <Switch
                  checked={formData.burnAfterRead}
                  onCheckedChange={(checked) =>
                    setFormData({ ...formData, burnAfterRead: checked })
                  }
                />
              </div>

              <div className="flex items-center justify-between rounded-none border border-border bg-secondary p-3">
                <div className="flex items-center gap-2">
                  <div>
                    <p className="text-sm font-medium text-foreground">
                      Public Paste
                    </p>
                    <p className="text-xs text-muted-foreground">
                      Visible in public search result
                    </p>
                  </div>
                </div>
                <Switch
                  checked={formData.isPublic}
                  onCheckedChange={(checked) =>
                    setFormData({ ...formData, isPublic: checked })
                  }
                />
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
      {IsLoggedIn() ? (
        <p className="text text-primary mb-4">
          <i className="text-orange-700">
            - If you don't have an account, don't forget to copy the link after
            posting, as there is no way to retrieve it later.
          </i>
        </p>
      ) : (
        ""
      )}
      <Button
        type="submit"
        size="lg"
        className="w-full"
        disabled={loading || !formData.content.trim()}
      >
        {loading ? "Creating Link..." : "Share"}
      </Button>
    </form>
  );
}
