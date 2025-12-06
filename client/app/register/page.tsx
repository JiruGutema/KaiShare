import { Header } from "@/components/header"
import { Footer } from "@/components/footer"
import { AuthForm } from "@/components/auth-form"

export default function RegisterPage() {
  return (
    <div className="min-h-screen flex flex-col">
      <Header />

      <main className="flex-1 flex items-center justify-center py-8 px-4">
        <AuthForm mode="register" />
      </main>

      <Footer />
    </div>
  )
}
