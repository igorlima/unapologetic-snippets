---
layout: default
title: API HttpClient and HttpServer
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/api-http-client-server
---

__[back]({% link docs/languages/golang/api.md %})__

Other resources:
- [Golang JSON]({% link docs/languages/golang/json.md %})
- [a http server sample]({% link docs/languages/golang/api.http.md %})

<br/>

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

-----

Make an HTTP POST request in Go

## HTTP Client

`http-client.go`

```golang
package main

import (
  "fmt"
  "io"
  "net/http"
  "net/url"
  "strings"
)

// make file=http-client.go target=app2 run
// https://gobyexample.com/http-clients
// https://www.kirandev.com/http-post-golang
func main() {
  posturl := "http://localhost:8091/form"
  values := make(url.Values)
  values.Set("name", "Jake")
  values.Set("country", "US")
  values.Set("state", "Connecticut")

  req, err := http.NewRequest(http.MethodPost, posturl, strings.NewReader(values.Encode()))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  if err != nil {
    fmt.Printf("error %v", err)
    return
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("issue during the request: %v", err)
    return
  }

  defer resp.Body.Close()
  bodyBytes, err := io.ReadAll(resp.Body)
  bodyString := string(bodyBytes)
  if err != nil {
    fmt.Println("issue reading response body")
  }

  fmt.Println(bodyString)
  fmt.Println("vim-go")
}
```

## HTTP Server

`http-server.go`

```golang
package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strings"
)

// make file=http-server.go target=app1 run
// https://gobyexample.com/http-servers
// https://golangbyexample.com/http-mutipart-form-body-golang/
func main() {
  http.HandleFunc("/form", handlerForm)
  http.HandleFunc("/multipart", handlerMultipart)
  http.ListenAndServe(":8091", nil)
  /*
   curl -X POST \
     http://localhost:8091/form \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "param1=value1&param2=value2" \
     | jq

   curl -X POST \
     http://localhost:8091/multipart \
     -F 'username=foo;' \
     -F 'password=bar;' \
     -F 'param1=value1;' \
     -F 'param2=value2;' \
     | jq
  */
}

func handlerForm(w http.ResponseWriter, req *http.Request) {
  data := map[string]string{}
  err := req.ParseForm()
  if err != nil {
    fmt.Fprintf(w, "error parsing the form", err.Error())
    return
  }
  for k, v := range req.Form {
    data[k] = strings.Join(v, " ")
  }
  jsonString, err := json.Marshal(data)
  if err != nil {
    fmt.Fprintf(w, "error converting to json", err.Error())
    return
  }
  fmt.Fprintf(w, string(jsonString))
}

func handlerMultipart(w http.ResponseWriter, req *http.Request) {
  data := map[string]string{}
  err := req.ParseMultipartForm(32 << 20)
  if err != nil {
    fmt.Fprintf(w, "error parsing the multipart form", err.Error())
    return
  }
  for k, v := range req.MultipartForm.Value {
    data[k] = strings.Join(v, " ")
  }
  jsonString, err := json.Marshal(data)
  if err != nil {
    fmt.Fprintf(w, "error converting to json", err.Error())
    return
  }
  fmt.Fprintf(w, string(jsonString))
}
```
