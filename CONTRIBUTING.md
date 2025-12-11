# Contributing to goswift

Thanks for your interest in contributing! 🎉
`goswift` aims to be a tiny, expressive, and composable REST framework for Go.

This guide will help you get started.

---

## 🛠 Requirements
- Go 1.21+
- Git installed
- Basic understanding of Go modules

## ⚙️ Project Setup

```bash
git clone https://github.com/jawadalik91/goswift.git
cd goswift
go mod tidy
```

Run tests:

```bash
go test ./...
```

Run example server:

```bash
go run ./cmd/example
```

---

## 🔧 Development Rules

### 1. Keep things small and readable
Prefer simplicity over clever tricks.

### 2. Add tests for any new feature
New PRs must include a test unless it’s a docs-only change.

### 3. Follow Go idioms
Use `gofmt`, `go vet`, and clear naming.

### 4. Open an Issue Before Large Changes
Large PRs without discussion may be rejected.

---

## 🧪 Running Tests
```bash
go test -v ./...
```

To run tests with coverage:
```bash
go test -cover ./...
```

---

## 🚀 Submitting a Pull Request

1. Fork the repo
2. Create a feature branch:
   ```bash
   git checkout -b feature/my-change
   ```
3. Make changes with tests
4. Commit using clear messages
5. Push the branch
6. Open a Pull Request
7. Fill out the PR template

---

## ❤️ Thank You
We appreciate all contributions — from code to documentation to ideas!
