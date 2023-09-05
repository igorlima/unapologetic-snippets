---
layout: default
title: API echo
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/api-echo
---

__[back]({% link docs/languages/golang/api.md %})__

# [Echo](https://github.com/labstack/echo)

High performance, minimalist Go web framework

<details markdown="block">
  <summary>
    getting started
  </summary>

A sample from a Medium post [^1].

```bash
go mod init github.com/igorlima/echo-2023a09m04d-21h19
go get github.com/labstack/echo/v4
```

```golang
package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })
  e.Logger.Fatal(e.Start(":3000"))
}
```
----
<br/>
<!-- getting started -->
</details>

---

[^1]: [Go: Gin vs Fiber vs Echo Hello World performance](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-vs-echo-hello-world-performance-a69a76a64d34)
