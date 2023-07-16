---
layout: default
title: API Go Toolkit
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/api-go-kit
---

__[back]({% link docs/languages/golang/api.md %})__

Other resources:
- [lib homepage](https://gokit.io/)
- [gihutb](https://github.com/go-kit/kit)

-----

## Intro

Go Toolkit (GoKit) is a toolkit for building microservices in Go. It provides a set of libraries, patterns, and utilities to simplify the development of scalable, modular, and maintainable microservices.

## Sample


<details markdown="block">
  <summary>
    code snippet
  </summary>

```bash
# init
go mod init github.com/igorlima/go-kit-sample
go get github.com/go-kit/kit/v2
go get github.com/go-kit/kit/log
go mod tidy
```

```golang
package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"

  "github.com/go-kit/kit/log"
)

type Resource struct {
  ID   int    `json:"id"`
  Name string `json:"name"`
}

func main() {
  logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
  logger.Log("msg", "starting")

  // resources := make([]Resource, 1)
  var resources []Resource

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
      // retrieve all resources
      json.NewEncoder(w).Encode(resources)
      logger.Log("retrieved", resources)
    case "POST":
      // create a new resource
      var resource Resource
      err := json.NewDecoder(r.Body).Decode(&resource)
      if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
      }
      resources = append(resources, resource)
      w.WriteHeader(http.StatusCreated)
    default:
      w.WriteHeader(http.StatusMethodNotAllowed)
    }
  })

  fmt.Println("Listening on http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}
```

```bash
# CRUD operations
curl -X POST http://localhost:8080 -d '{"id":1,"name":"foo"}'
curl -X POST http://localhost:8080 -d '{"id":2,"name":"bar"}'
```

------

</details>
