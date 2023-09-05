---
layout: default
title: API Gin
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/api-gin
---

__[back]({% link docs/languages/golang/api.md %})__

# [Gin](https://github.com/gin-gonic/gin)

<details markdown="block">
  <summary>
    getting started
  </summary>

A sample from a Medium post [^1].

```bash
go mod init github.com/igorlima/gin-2023a09m04d-21h19
go get -u github.com/gin-gonic/gin
```

```golang
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.New()

  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello, world!")
  })

  r.Run(":3000")
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

/*
go mod init igorlima/my-gin-api
go mod tidy

require github.com/gin-gonic/gin v1.9.0

curl -X GET localhost:8080/ping
*/

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
  })
  r.Run() // listen and serve on 0.0.0.0:8080
}
```
----
<br/>
<!-- sample A -->
</details>

<details markdown="block">
  <summary>
    sample B
  </summary>
```js
// encode a string to Base64 in JavaScript
btoa('foo:bar')
// 'Zm9vOmJhcg=='
atob('Zm9vOmJhcg==')
// 'foo:bar'
```

```golang
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
  // Disable Console Color
  // gin.DisableConsoleColor()
  r := gin.Default()

  // Ping test
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  // Get user value
  r.GET("/user/:name", func(c *gin.Context) {
    user := c.Params.ByName("name")
    value, ok := db[user]
    if ok {
      c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
    } else {
      c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
    }
  })

  // Authorized group (uses gin.BasicAuth() middleware)
  // Same than:
  // authorized := r.Group("/")
  // authorized.Use(gin.BasicAuth(gin.Credentials{
  //    "foo":  "bar",
  //    "manu": "123",
  //}))
  authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
    "foo":  "bar", // user:foo password:bar
    "manu": "123", // user:manu password:123
  }))

  /* example curl for /admin with basicauth header
     Zm9vOmJhcg== is base64("foo:bar")

    curl -X POST \
      http://localhost:8080/admin \
      -H 'authorization: Basic Zm9vOmJhcg==' \
      -H 'content-type: application/json' \
      -d '{"value":"bar"}'
  */
  authorized.POST("admin", func(c *gin.Context) {
    user := c.MustGet(gin.AuthUserKey).(string)

    // Parse JSON
    var json struct {
      Value string `json:"value" binding:"required"`
    }

    if c.Bind(&json) == nil {
      db[user] = json.Value
      c.JSON(http.StatusOK, gin.H{"status": "ok"})
    }
  })

  return r
}

func main() {
  r := setupRouter()
  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}
```
----
<br/>
<!-- sample B -->
</details>

---

[^1]: [Go: Gin vs Fiber vs Echo Hello World performance](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-vs-echo-hello-world-performance-a69a76a64d34)
