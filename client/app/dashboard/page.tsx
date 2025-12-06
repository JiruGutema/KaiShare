import { Header } from "@/components/header"
import { Footer } from "@/components/footer"
import { DashboardContent } from "@/components/dashboard-content"

export default function DashboardPage() {
  return (
    <div className="min-h-screen flex flex-col">
      <Header />

      <main className="flex-1 py-8">
        <div className="mx-auto max-w-6xl px-4">
          <DashboardContent />
        </div>
      </main>

      <Footer />
    </div>
  )
}
