---
layout: default
title: API Fiber
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/api-fiber
---

__[back]({% link docs/languages/golang/api.md %})__

# [Fiber](https://github.com/gofiber/fiber)

An Express-inspired web framework written in Go.

Fiber is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.

Fiber supports multiple template engines, such as Handlebars and Pug, thanks to the [template middleware](https://github.com/gofiber/template/blob/d7d478414dd222733f234ca8c5f07b238607f3f0/README.md).

<details markdown="block">
  <summary>
    getting started
  </summary>

A sample from a Medium post [^1].

```bash
go mod init github.com/igorlima/fiber-2023a09m04d-21h19
go get -u github.com/gofiber/fiber/v2
```

```golang
package main

import (
  "log"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()
  port := ":3000"

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  log.Fatal(app.Listen(port))
}
```
----
<br/>
<!-- getting started -->
</details>

<details markdown="block">
  <summary>
    sample A
  </summary>
```golang
package main

import (
  "log"

  "github.com/gofiber/fiber/v2"
)

/**
$> go list -m -versions <module_name>
$> go list -m -versions github.com/gofiber/fiber/v2

$> go install github.com/gofiber/fiber/v2@v2.43.0
$> go install github.com/gofiber/fiber/v2@latest

$> go mod init hertz_demo
$> go mod tidy

$> curl http://127.0.0.1:3000/
*/

func main() {
  app := fiber.New()

  // Hello World
  app.Get("/hello", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  // Serve Static Files
  // curl localhost:3000/public/function.anonymous.2023a04m03d.go
  app.Static("/public", "./archive")

  // Routing
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("GET request")
  })
  app.Get("/:param", func(c *fiber.Ctx) error {
    return c.SendString("param: " + c.Params("param"))
  })
  app.Post("/", func(c *fiber.Ctx) error {
    return c.SendString("POST request")
  })

  // API-ready
  // curl localhost:3000/api/posts
  posts := [6]string{"post a", "post b", "post c"}
  app.Get("/api/posts", func(c *fiber.Ctx) error {
    if len(posts) == 0 {
      return c.Status(404).JSON(&fiber.Map{
        "success": false,
        "error":   "There are no posts!",
      })
    }
    return c.JSON(&fiber.Map{
      "success": true,
      "posts":   posts,
    })
  })

  // Start listening
  log.Fatal(app.Listen(":3000"))
}
```
----
<br/>
<!-- sample A -->
</details>


---

[^1]: [Go: Gin vs Fiber vs Echo Hello World performance](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-vs-echo-hello-world-performance-a69a76a64d34)
