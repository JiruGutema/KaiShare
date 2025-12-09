"use server";

import { Header } from "./header";
import { Card } from "./ui/card";

export async function AboutPage() {
  return (
    <>
      <Header />
      <main className="mx-auto max-w-3xl px-4 py-12">
        <section className="bg-card/90 flex flex-col rounded-none shadow gap-10 p-8 border border-none">
          {/* Title and Intro */}
          <Card className="p-8">
            <p className="text text-muted-foreground mb-8">
              <span className="text-orange-600 font-bold">Kai</span>
              <span className="text-sky-700 font-bold">Share</span> is a secure
              platform for sharing code snippets and text with others. Built
              with a focus on privacy and simplicity, KaiShare helps developers
              and users exchange information safely—ideal for teams,
              communities, and individuals working on collaborative coding and
              sharing stuff.
            </p>
          </Card>

          <Card className="p-8">
            <section className="mb-8">
              <p className="text-muted-foreground dark:text-gray-300">
                We believe sharing code and text should be simple, secure, and
                sustainable. KaiShare empowers users to control their visibility
                and privacy while making collaboration effortless—whether
                sharing sensitive snippets or texts with your team without{" "}
                <strong className="text-red-500"> bloating</strong> your social
                media . Instead, share them only a link.
              </p>
            </section>
          </Card>

          <Card className="p-8">
            <section className="mb-8">
              <h2 className="text font-semibold text-primary mb-2">
                Key Features
              </h2>
              <ul className="list-none space-y-2 pl-6 text-gray-700 dark:text-gray-300">
                <li>
                  <strong>Secure Paste Sharing:</strong> Share code or text with
                  robust options—password-protection and{" "}
                  <em>burn-after-read</em> ensure your data stays private when
                  needed.
                </li>
                <li>
                  <strong>Rich Syntax Highlighting:</strong> Over 25 programming
                  languages supported, with a true monochrome custom highlight
                  theme that always ensures contrast and readability.
                </li>
                <li>
                  <strong>Account System:</strong> User accounts let you
                  revisit, manage, or remove your pastes at any time.
                </li>
                <li>
                  <strong>Privacy Controls:</strong> Choose public, unlisted, or
                  private for every share; enable or disable features per-paste.
                </li>
                <li>
                  <strong>Mobile-Ready Experience:</strong> Fully responsive,
                  fast UI—share and read pastes anywhere, securely.
                </li>
                <li>
                  <strong>Completely Free & Open Source:</strong> KaiShare is
                  open and community-driven. Fork, self-host, or contribute.
                </li>
              </ul>
            </section>
          </Card>

          <Card className="p-8">
            <p className="text text-primary mb-4">
              <span className="text-orange-600 font-bold"> PS</span>- If you
              don't have an account, don't forget to copy the link (like
              <span className="font-mono bg-muted px-1 py-0.5 rounded-none mx-1">
                https://kaishare.vercel.app/p/9913f9-d344-11f0-80bb-8c8d28d905b2){" "}
              </span>
              after posting, as there is no way to retrieve it later.
            </p>
          </Card>
          <section className="text-center mt-8">
            <p className="text text-muted-foreground">
              Explore KaiShare. Share code and snippets securely.
              <br />
              The future of paste sharing is safe, elegant, and open-source.
            </p>
            <a href="/" aria-label="Get Started with KaiShare">
              <div className="mt-6 inline-block bg-primary text-primary-foreground rounded-none px-6 py-2 font-semibold shadow-sm hover:opacity-90 transition">
                Get Started
              </div>
            </a>
          </section>
        </section>
      </main>
    </>
  );
}
