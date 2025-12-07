import { Terminal, Github } from "lucide-react";

export function Footer() {
  return (
    <footer className="border-t border-border bg-card mt-auto">
      <div className="mx-auto max-w-6xl px-4 py-6">
        <div className="flex flex-col sm:flex-row items-center justify-between gap-4">
          <div className="flex items-center gap-2 text-muted-foreground">
            <span className="text-sm">GoPastebin</span>
            <span className="text-xs">v1.0.0</span>
          </div>
          
          <div>
            <a
              href="/about"
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center gap-1 text-muted-foreground hover:text-foreground transition-colors text-sm"
            >

              <span className="">About</span>
            </a>
          </div>
          <div className="flex items-center gap-2">
            <a
              href="https://github.com/jirugutema/kaishare"
              target="_blank"
              rel="noopener noreferrer"
              className="text-muted-foreground hover:text-foreground transition-colors"
            >
              <Github className="h-4 w-4" />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
}
