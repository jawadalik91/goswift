package router

import (
    "net/http"
    "strings"
    "time"
)

type route struct {
    pattern string
    handler HandlerFunc
    parts   []string
}

type Router struct {
    routes     map[string][]*route
    middleware []Middleware
}

func NewRouter() *Router {
    return &Router{
        routes: make(map[string][]*route),
    }
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    method := req.Method
    path := req.URL.Path

    if list, ok := r.routes[method]; ok {
        for _, rt := range list {
            if params, ok := matchPath(rt.parts, path); ok {
                ctx := NewContext(w, req, params)
                r.runMiddleware(ctx, rt.handler)
                return
            }
        }
    }

    http.NotFound(w, req)
}

func (r *Router) Add(method, pattern string, handler HandlerFunc) {
    parts := splitPath(pattern)
    rt := &route{
        pattern: pattern,
        handler: handler,
        parts:   parts,
    }
    r.routes[method] = append(r.routes[method], rt)
}

func (r *Router) Use(m Middleware) {
    r.middleware = append(r.middleware, m)
}

func (r *Router) runMiddleware(c *Context, final HandlerFunc) {
    // Build chain: middleware -> final
    idx := 0
    var exec func(*Context)
    exec = func(ctx *Context) {
        if idx >= len(r.middleware) {
            final(ctx)
            return
        }
        m := r.middleware[idx]
        idx++
        // middleware calls next handler by invoking provided next function
        m(ctx, func(nc *Context) {
            exec(nc)
        })
    }
    start := time.Now()
    exec(c)
    _ = start // placeholder: you could log duration
}

// helpers

func splitPath(p string) []string {
    p = strings.TrimSpace(p)
    if p == "" || p == "/" {
        return []string{}
    }
    p = strings.TrimPrefix(p, "/")
    return strings.Split(p, "/")
}

func matchPath(parts []string, path string) (map[string]string, bool) {
    target := splitPath(path)
    if len(parts) != len(target) {
        return nil, false
    }
    params := make(map[string]string)
    for i := 0; i < len(parts); i++ {
        if strings.HasPrefix(parts[i], ":") {
            name := strings.TrimPrefix(parts[i], ":")
            params[name] = target[i]
        } else if parts[i] != target[i] {
            return nil, false
        }
    }
    return params, true
}
