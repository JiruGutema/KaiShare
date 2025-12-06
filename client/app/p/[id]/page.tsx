import { Header } from "@/components/header"
import { Footer } from "@/components/footer"
import { PasteViewer } from "@/components/paste-viewer"

interface PastePageProps {
  params: Promise<{ id: string }>
}

export default async function PastePage({ params }: PastePageProps) {
  const { id } = await params

  return (
    <div className="min-h-screen flex flex-col">
      <Header />

      <main className="flex-1 py-8">
        <div className="mx-auto max-w-4xl px-4">
          <PasteViewer id={id} />
        </div>
      </main>

      <Footer />
    </div>
  )
}
