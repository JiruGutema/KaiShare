"use client"

import { useEffect, useRef } from "react"
import hljs from "highlight.js"

interface SyntaxHighlighterProps {
  code: string
  language: string
}

export function SyntaxHighlighter({ code, language }: SyntaxHighlighterProps) {
  const codeRef = useRef<HTMLElement>(null)

  useEffect(() => {
    if (codeRef.current) {
      hljs.highlightElement(codeRef.current)
    }
  }, [code, language])

  return (
    <pre className="overflow-x-auto">
      <code ref={codeRef} className={`language-${language}`}>
        {code}
      </code>
    </pre>
  )
}
