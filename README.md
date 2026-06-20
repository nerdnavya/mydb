# MyDB — A Relational Database Engine from Scratch


 **Work in progress** — building a database from scratch in Go to understand how storage engines, indexing, and query execution actually work under the hood.

## Why I built this

Databases are used everywhere but rarely understood at the implementation level. This is my attempt to build one from first principles, in public, one layer at a time.

## Current status

| Component | Status |
|---|---|
| Disk-backed file storage | ✅ Working — can persist data to disk |
| Page management | 🚧 In progress |
| B+tree indexing | ⬜ Not started |
| SQL parser (SELECT/INSERT/UPDATE/DELETE) | ⬜ Not started |
| Query planner / execution engine | ⬜ Not started |
| Transactions | ⬜ Not started |

## Planned architecture
Client query → Parser → Planner → Execution Engine → Storage Layer
## Tech stack

Go (standard library only — no external dependencies, intentionally)

## Running it

```bash
git clone https://github.com/nerdnavya/mydb
cd mydb
go run .
```

Currently supports writing and persisting data to disk for any given file — this is the first working piece of the storage layer. Query parsing, indexing, and the execution engine are not built yet (see status table above).

## What I'm learning

* How page management and disk I/O actually work, not just in theory
