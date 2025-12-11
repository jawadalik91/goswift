# goswift

goswift is a tiny, expressive REST micro-framework for Go inspired by Express.js.
This repository contains a minimal but functional implementation to get you started.

## Quick start

```bash
go run ./cmd/example
```

Visit http://localhost:8080/hello

## Features
- Express-like routing: GET, POST
- Route params: /users/:id
- Middleware chaining
- Context helpers: JSON, String, Param, Query, BindJSON

## Project layout
```
goswift/
├── cmd/example
├── internal/router
├── goswift.go
└── README.md
```

## Contributing
See CONTRIBUTING.md for contribution guidelines.
