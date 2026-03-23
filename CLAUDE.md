# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Optimus is a cross-platform desktop image optimization app (WebP/JPEG/PNG compression, conversion, and resizing) built with **Wails v2** — a Go backend bridged to a Vue.js frontend that compiles into a single native binary.

## Commands

### Development
```bash
wails dev          # Run in dev mode (starts both Go backend and Vue dev server)
```

### Building
```bash
wails build                          # Build for current OS
wails build -tags webkit2_41         # Linux build
```

### Frontend (run from `frontend/`)
```bash
npm install        # Install dependencies
npm run lint       # ESLint
npm run build      # Production build (also triggered by wails build)
```

### Backend
```bash
go vet ./backend/...   # Lint Go code
go test ./backend/...  # Run Go tests
```

## Architecture

### Backend → Frontend Bridge (Wails v2)
- `main.go` — creates the Wails app, binds Go structs as RPC, embeds `frontend/dist`
- Go methods exposed via Wails are auto-generated as TypeScript bindings in `frontend/src/wailsjs/go/`
- Events flow backend→frontend via `runtime.EventsEmit()` (e.g., `conversion:complete`, `notify`)

### Backend (`backend/`)
| Package | Role |
|---|---|
| `image/` | `FileManager` — batch orchestrator using goroutines; `File` — per-image decode/resize/encode |
| `config/` | App settings, loaded/saved to `conf.json` in the OS user config dir |
| `stat/` | Lifetime stats, persisted to `stats.json` |
| `jpeg/`, `png/`, `webp/` | Format-specific encode/decode implementations |
| `localstore/` | Cross-platform user config directory I/O |

Image processing pipeline: frontend base64-encodes file → Go `HandleFile()` decodes → resizes (fill/fit/smart crop via `imaging`/`smartcrop`) → re-encodes to target format → emits result event.

### Frontend (`frontend/src/`)
Vue 2 + Vuex. `App.vue` is the root with a sidebar + view router pattern (no vue-router — just conditional rendering). State lives in `store.js` (config, lifetime stats, session stats). Components communicate via both Vuex and a simple `event-bus.js`.

Key components: `Editor.vue` (drag-drop + file list), `Settings.vue` (quality/size/format config), `Stats.vue` (lifetime stats display).

## Key Dependencies
- **Go**: `wailsapp/wails/v2`, `chai2010/webp`, `disintegration/imaging`, `muesli/smartcrop`, `foobaz/lossypng`
- **JS**: Vue 2.6, Vuex 3, Tailwind CSS 1.9

## Wails v2 Notes
- Config in `wails.json` (schema v2)
- Frontend build is managed by Wails (`frontend:build`, `frontend:dev:watcher` hooks in `wails.json`)
- Linux requires `webkit2gtk` — CI uses `-tags webkit2_41`
- This repo is on branch `chore/migrate-to-wails-v2` — the migration from Wails v1 to v2 is the current focus
