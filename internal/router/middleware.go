package router

// Middleware is a function that wraps request handling.
// It receives the context and a next handler to call.
type Middleware func(*Context, HandlerFunc)

// HandlerFunc is the type for request handlers.
type HandlerFunc func(*Context)

// Example built-in logger middleware
func Logger() Middleware {
    return func(c *Context, next HandlerFunc) {
        // simple logger: you may expand
        // fmt.Printf("%s %s\n", c.Request.Method, c.Request.URL.Path)
        next(c)
    }
}

// Recovery middleware example
func Recovery() Middleware {
    return func(c *Context, next HandlerFunc) {
        defer func() {
            if r := recover(); r != nil {
                c.JSON(500, map[string]any{"error": "internal server error"})
            }
        }()
        next(c)
    }
}
