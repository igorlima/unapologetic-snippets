---
layout: default
title: API Hertz
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/api-hertz
---

__[back]({% link docs/languages/golang/api.md %})__

# [Hertz](https://github.com/cloudwego/hertz)

Go HTTP framework with high- performance and strong- extensibility for building micro-services.

```golang
package main

import (
  "context"
  "fmt"

  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "github.com/cloudwego/hertz/pkg/protocol/consts"
)

/**
$> go list -m -versions <module_name>
$> go list -m -versions github.com/cloudwego/hertz/cmd/hz

$> go install github.com/cloudwego/hertz/cmd/hz@v0.6.1
$> go install github.com/cloudwego/hertz/cmd/hz@latest

$> go mod init hertz_demo
$> go mod tidy

$> curl http://127.0.0.1:8888/ping
*/

func main() {
  h := server.Default()

  h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
    ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
  })

  h.Spin()
  fmt.Println("vim-go")
}
```

## more examples

The Hertz-Examples repository provides code out of the box.
- [more](https://www.cloudwego.io/docs/hertz/tutorials/example/)
- [more examples](https://github.com/cloudwego/hertz-examples/blob/100eb4b5bcaa3161a5194a5b31a3900aaf943e8d/README.md)
