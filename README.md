# MyDB — A Relational Database Engine from Scratch

> Building to understand how databases actually work under the hood.

## Why I built this
Databases are used everywhere but rarely understood at the implementation level. This project is my attempt to build one from first principles.

## What it supports
- SQL-like query parsing (SELECT, INSERT, UPDATE, DELETE)
- B+tree / page-based storage engine
- Basic indexing for faster lookups
- In-memory and file-backed storage modes
- Transaction support (in progress)

## Tech stack
`Go` `Standard Library only`

## Architecture
```
Client query → Parser → Planner → Execution Engine → Storage Layer
```

## How to build
```bash
git clone https://github.com/your-username/mydb
cd mydb
mkdir build && cd build
cmake .. && make
./mydb
```

## What I learned
- How B+trees enable O(log n) lookups
- Page management and disk I/O patterns
- How SQL parsers tokenize and plan queries
