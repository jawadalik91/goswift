package main

import (
    "fmt"
    "github.com/jawadalik91/goswift/internal/router"
    "github.com/jawadalik91/goswift"
)

func main() {
    app := goswift.New()

    // simple logger middleware
    app.Use(router.Logger())

    app.GET("/hello", func(c *router.Context) {
        c.JSON(200, map[string]any{"message": "Hello World"})
    })

    app.GET("/users/:id", func(c *router.Context) {
        id := c.Param("id")
        c.String(200, fmt.Sprintf("User ID = %s", id))
    })

    fmt.Println("Listening on :8080")
    if err := app.Listen(":8080"); err != nil {
        panic(err)
    }
}
