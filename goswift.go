package goswift

import (
    "net/http"
    "github.com/jawadalik91/goswift/internal/router"
)

// App is the main application object
type App struct {
    router *router.Router
}

// New creates a new App
func New() *App {
    return &App{
        router: router.NewRouter(),
    }
}

// Use registers a middleware
func (a *App) Use(m router.Middleware) {
    a.router.Use(m)
}

// GET registers a GET route
func (a *App) GET(path string, handler router.HandlerFunc) {
    a.router.Add("GET", path, handler)
}

// POST registers a POST route
func (a *App) POST(path string, handler router.HandlerFunc) {
    a.router.Add("POST", path, handler)
}

// Listen starts HTTP server on addr
func (a *App) Listen(addr string) error {
    return http.ListenAndServe(addr, a.router)
}
