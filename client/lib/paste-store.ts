// Simple in-memory store for pastes (demo purposes)
// In production, you'd use a database

export interface Paste {
  id: string
  title: string
  content: string
  language: string
  password?: string
  burnAfterRead: boolean
  expiresAt?: Date
  createdAt: Date
  views: number
  userId?: string
  isPublic: boolean
}

export interface User {
  id: string
  email: string
  username: string
  passwordHash: string
  createdAt: Date
}

// Global store (persists across requests in development)
const globalStore = globalThis as unknown as {
  pastes: Map<string, Paste>
  users: Map<string, User>
  sessions: Map<string, string>
}

if (!globalStore.pastes) {
  globalStore.pastes = new Map()
}
if (!globalStore.users) {
  globalStore.users = new Map()
}
if (!globalStore.sessions) {
  globalStore.sessions = new Map()
}

export const pastes = globalStore.pastes
export const users = globalStore.users
export const sessions = globalStore.sessions

export function generateId(): string {
  return Math.random().toString(36).substring(2, 10)
}

export function hashPassword(password: string): string {
  // Simple hash for demo - in production use bcrypt
  let hash = 0
  for (let i = 0; i < password.length; i++) {
    const char = password.charCodeAt(i)
    hash = (hash << 5) - hash + char
    hash = hash & hash
  }
  return hash.toString(36)
}
