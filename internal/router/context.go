package router

import (
    "encoding/json"
    "io"
    "net/http"
)

type Context struct {
    Writer  http.ResponseWriter
    Request *http.Request
    Params  map[string]string
}

func NewContext(w http.ResponseWriter, r *http.Request, p map[string]string) *Context {
    if p == nil {
        p = make(map[string]string)
    }
    return &Context{
        Writer:  w,
        Request: r,
        Params:  p,
    }
}

func (c *Context) Param(name string) string {
    return c.Params[name]
}

func (c *Context) Query(name string) string {
    return c.Request.URL.Query().Get(name)
}

func (c *Context) JSON(status int, v any) {
    c.Writer.Header().Set("Content-Type", "application/json")
    c.Writer.WriteHeader(status)
    _ = json.NewEncoder(c.Writer).Encode(v)
}

func (c *Context) String(status int, s string) {
    c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
    c.Writer.WriteHeader(status)
    _, _ = c.Writer.Write([]byte(s))
}

func (c *Context) BindJSON(dst any) error {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        return err
    }
    if len(body) == 0 {
        return nil
    }
    return json.Unmarshal(body, dst)
}
