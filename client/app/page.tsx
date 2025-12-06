import { Header } from "@/components/header"
import { Footer } from "@/components/footer"
import { PasteForm } from "@/components/paste-form"
import { RecentPastes } from "@/components/recent-pastes"
import { FileCode, Lock, Flame, Users } from "lucide-react"

export default function HomePage() {
  return (
    <div className="min-h-screen flex flex-col">
      <Header />

      <main className="flex-1 py-8">
        <div className="mx-auto max-w-6xl px-4">
          <div className="mb-8 text-center">
            <h1 className="text-xl font-bold text-foreground mb-2">
              Share Code & Text <span className="text-primary">Securely</span>
            </h1>
            <p className="text-muted-foreground">Syntax highlighting, password protection, and burn after read</p>
          </div>

          <div className="grid gap-8 lg:grid-cols-3">
            <div className="lg:col-span-2">
              <PasteForm />
            </div>
            <div className="space-y-6">
              <RecentPastes />

              <div className="grid grid-cols-2 gap-3">
                <div className="rounded-none border border-border bg-card p-4 text-center">
                  <FileCode className="h-6 w-6 text-primary mx-auto mb-2" />
                  <p className="text-xs text-muted-foreground">25+ Languages</p>
                </div>
                <div className="rounded-none border border-border bg-card p-4 text-center">
                  <Lock className="h-6 w-6 text-primary mx-auto mb-2" />
                  <p className="text-xs text-muted-foreground">Password Lock</p>
                </div>
                <div className="rounded-none border border-border bg-card p-4 text-center">
                  <Flame className="h-6 w-6 text-destructive mx-auto mb-2" />
                  <p className="text-xs text-muted-foreground">Burn After Read</p>
                </div>
                <div className="rounded-none border border-border bg-card p-4 text-center">
                  <Users className="h-6 w-6 text-primary mx-auto mb-2" />
                  <p className="text-xs text-muted-foreground">User Accounts</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>

      <Footer />
    </div>
  )
}
