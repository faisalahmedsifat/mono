# monorun

**A lightweight, language-agnostic development process manager for monorepos.**

Run your entire Turborepo (or any multi-service project) with **one command** — Go services, Hono APIs, Next.js apps, workers, and more — all with smart file watching and automatic restarts.

---

## What is monorun?

`monorun` is a simple yet powerful dev tool written in Go that:

- Starts multiple services concurrently (different languages, frameworks, ports)
- Watches files per service and intelligently restarts only what's needed
- Provides clean, prefixed, colored logs
- Handles graceful shutdowns
- Supports environment variables and dependencies between services

It's designed specifically for **mixed-language monorepos** like Turborepo, where you might have:
- Go APIs / daemons / CLIs
- TypeScript backends (Hono, Express, etc.)
- Next.js / Vite / React frontends
- Python scripts, Rust tools, etc.

---

## Why does it matter?

Modern monorepos are powerful but messy to develop in:

- Running `turbo dev`, `concurrently`, or multiple terminals gets painful fast.
- Different tools have different watch behaviors (`air` for Go, `vite` for frontend, `nodemon` for TS, etc.).
- You lose unified logging and control.

**monorun** solves this by giving you **one unified dev experience** while staying extremely flexible.

### Benefits:
- One command to rule them all (`monorun`)
- Instant feedback when editing any part of the stack
- Better DX than mixing `air` + `concurrently` + `turbo`
- Full control because you built (or can easily extend) it
- Lightweight and fast (native Go)

---

## Who is it for?

- **Developers working in Turborepos** with mixed languages (Go + TypeScript especially)
- Full-stack engineers who want a clean dev environment
- Teams maintaining multiple services in one repo
- Anyone tired of managing 5+ terminals during development
- People who enjoy owning their tooling (and learning Go along the way)

---

## How to Use It

### 1. Installation

```bash
# After building
go install monorun@latest
```

Or run directly during development:
```bash
go run cmd/monorun/main.go
```

### 2. Configuration (`monorun.yaml`)

Create a file called `monorun.yaml` in your repository root:

```yaml
services:
  api-go:
    dir: apps/api-go
    command: go run main.go
    watch:
      - "**/*.go"
      - "go.mod"
      - "go.sum"
    env:
      PORT: "8081"
      ENV: "development"

  hono-api:
    dir: apps/hono-api
    command: pnpm dev
    watch:
      - "**/*.ts"
      - "**/*.tsx"

  web-next:
    dir: apps/web
    command: pnpm dev
    watch:
      - "**/*.{ts,tsx,js,jsx}"
    port: 3000   # optional - used for health checks later

  worker:
    dir: services/worker
    command: go run .
    watch:
      - "**/*.go"
```

### 3. Running

```bash
# Run everything
monorun

# Run only specific services
monorun --only api-go,web-next

# Use custom config
monorun --config monorun.staging.yaml
```

**Example output:**
```
[api-go]   2026/06/15 06:45:12 INFO: Starting server on :8081
[web-next] Ready in 234ms
[hono-api] Hono dev server running...
```

Press `Ctrl+C` → everything shuts down gracefully.

---

## Roadmap (Future Features)

- Service dependencies (`depends_on`)
- Health checks + automatic URL opening
- Better log filtering and levels
- `monorun init` command to generate config
- Support for custom build steps before running

---

## Why build this yourself?

- Deeply understand process management and file watching in Go
- Learn practical concurrency patterns (`errgroup`, channels, contexts)
- Create a tool perfectly tailored to **your** stack
- Level up your tooling skills (very high ROI for a serious developer)

---