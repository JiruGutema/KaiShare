# KaiShare — Weaknesses & Roadmap

> Assessment of the KaiShare server (`server/`) and monorepo. Last reviewed: 2026-05-30.

## Critical Weaknesses (bugs that break things)

1. **`PUT /paste/:id` is broken at runtime** — the `UPDATE` SQL in
   `internal/repository/paste.repository.go:163` has a malformed `RETURNING`
   clause (trailing comma after `created_at,`). Update fails every time.
2. **Nil-pointer panic deleting anonymous pastes** — `DeletePasteService`
   (`internal/service/paste.services.go:114`) dereferences `*res.UserID` with no
   nil check. Anonymous pastes have `UserID == nil`, so deleting one panics the server.
3. **Notifications feature is dead code** — handler/service/repository all exist
   but no route registers them. Schema mismatch: code inserts a `relation_link`
   column the documented schema lacks.

## Security Weaknesses

1. **Secrets committed to git** — `server/.env` and `client/.env` are tracked.
   Remove them, rotate the secrets, and keep only `.env.example` in the repo.
2. **Hardcoded fallback `"dev-secret"`** session secret (`routes.go:28`) if the
   env var is unset — dangerous if it ever ships to prod.
3. **Debug `fmt.Println` leaks the submitted paste password to stdout**
   (`internal/handler/paste.handler.go:57`), plus stray prints elsewhere.
4. **Rate limiting is global, not per-IP/per-user** — one shared leaky bucket.
   A single abuser starves everyone and abuse isn't actually prevented.
5. **No real input validation** — paste title/content/language/expiry have no
   length limits or sanitization. `pkg/validation.go` is an empty stub.
6. **"Encryption" is fiction** — the README advertises encryption options, but
   content is stored plaintext; only optional passwords are bcrypt-hashed.
7. **Cookie/token expiry mismatch** — access-token cookie `MaxAge` is 1h but the
   JWT lives 24h.

## Quality / Maintenance Gaps

- No DB migrations (schema applied by hand from docs).
- Stale docs — `API.md` still says "Kaipaste" and documents the old `?password=`
  query param (now moved to the request body).
- Minimal tests; `TestGetAllPastesRoute_Exists` actually hits `/auth/register`
  (copy-paste bug).
- Empty `docker-compose.yaml`, unused `templates/`, unwired `pkg/logger.go`.

## What's Solid

bcrypt cost 14, parameterized SQL (no injection), JWT alg-confusion protection,
HttpOnly cookies, scoped CORS, ownership checks on delete/update.

## Roadmap

### Fix-first (before any new feature)
1. Repair the `UpdatePaste` SQL and add the nil-check in delete.
2. Remove committed `.env` files, rotate secrets, strip debug prints.
3. Wire up or delete the notification code.

### Then build

| Priority | Feature | Why |
|---|---|---|
| High | Real per-IP rate limiting | Current global limiter is ineffective against abuse |
| High | Input validation layer (length caps, allowed languages, expiry bounds) | Prevents DoS-by-huge-paste and bad data |
| High | Client-side / end-to-end encryption | Backs the README's main selling point |
| Medium | Paste editing history / versions | Natural extension of the update endpoint |
| Medium | Public paste discovery / search (`is_public` exists but is unused) | Schema is already there |
| Medium | Password reset & email verification | Auth exists but account recovery doesn't |
| Medium | Activate notifications (paste viewed, expiring soon) | Code is ~80% written already |
| Low | DB migrations (golang-migrate) + syntax-highlight metadata | Maintainability + UX |
| Low | Paste analytics (view counts already tracked) | Cheap win on existing data |
