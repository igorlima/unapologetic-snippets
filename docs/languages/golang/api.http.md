---
layout: default
title: API http
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/api-http
---

__[back]({% link docs/languages/golang/api.md %})__

Other resources:
- [Golang JSON]({% link docs/languages/golang/json.md %})
- [http client and http server]({% link docs/languages/golang/api.http-client-server.md %})

-----

http server - how to convert and respond in JSON

```golang
package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

// https://gobyexample.com/http-servers
func main() {
  // curl -X GET http://localhost:8090/hello
  http.HandleFunc("/hello", hello)
  // curl -X GET -H 'foo: bar' http://localhost:8090/headers
  http.HandleFunc("/headers", headers)
  // curl -X POST http://localhost:8090/body -d '{"foo":"bar"}'
  http.HandleFunc("/body", body)
  // curl -X POST -H 'foo: bar' http://localhost:8090/headerAndBody -d '{"foo":"bar"}'
  http.HandleFunc("/headerAndBody", headerAndBody)
  // curl -X POST http://localhost:8090/decodeJson -d '{"foo":"bar"}'
  http.HandleFunc("/decodeJson", decodeJson)
  // curl -X POST http://localhost:8090/structAsJson -d '{"foo":"bar","id":123,"name":"xyz"}'
  http.HandleFunc("/structAsJson", outputStructAsJson)
  // curl -X GET http://localhost:8090/stringAsJson
  http.HandleFunc("/stringAsJson", outputStringAsJson)
  // curl -X GET http://localhost:8090/mapAsJson
  http.HandleFunc("/mapAsJson", outputMapAsJson)
  // curl -X GET http://localhost:8090/setAsJson
  http.HandleFunc("/setAsJson", outputSetAsJson)

  http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
    for _, h := range headers {
      fmt.Fprintf(w, "%v: %v\n", name, h)
    }
  }
}

func body(w http.ResponseWriter, req *http.Request) {
  var data map[string]string
  jsonData, err := ioutil.ReadAll(req.Body)
  if err != nil {
    fmt.Fprintf(w, "error reading request body", err.Error())
  }
  err = json.Unmarshal([]byte(jsonData), &data)
  if err != nil {
    fmt.Fprintf(w, "error while decoding the data", err.Error())
  }
  fmt.Fprintf(w, "raw request body -> %v\n", string(jsonData))
  fmt.Fprintf(w, "request body map[string]string -> %v\n", data)
}

func headerAndBody(w http.ResponseWriter, req *http.Request) {
  headers(w, req)
  body(w, req)
}

func decodeJson(w http.ResponseWriter, req *http.Request) {
  var data map[string]string
  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&data)
  if err != nil {
    fmt.Fprintf(w, "error reading request body", err.Error())
  }
  fmt.Fprintf(w, "body decoded to map[string]string -> %v\n", data)
}

func outputStructAsJson(w http.ResponseWriter, req *http.Request) {
  type Message struct {
    Id   int64  `json:"id"`
    Name string `json:"name"`
  }

  jsonData, err := ioutil.ReadAll(req.Body)
  if err != nil {
    fmt.Fprintf(w, "error reading request body", err.Error())
  }

  var msg Message
  err = json.Unmarshal([]byte(jsonData), &msg)
  if err != nil {
    fmt.Fprintf(w, "error while unmarshaling", err.Error())
  } else {
    fmt.Fprintf(w, "%v", msg)
  }
}

func outputStringAsJson(w http.ResponseWriter, req *http.Request) {
  var person map[string]interface{}
  jsonData := `{
    "name":"Jake",
    "country":"US",
    "state":"Connecticut"
  }`
  err := json.Unmarshal([]byte(jsonData), &person)
  if err != nil {
    fmt.Fprintf(w, "error while decoding the data", err.Error())
  } else {
    fmt.Fprintf(w, "%v", person)
  }
}

func outputMapAsJson(w http.ResponseWriter, req *http.Request) {
  myMap := map[string]string{
    "name":    "Jake",
    "country": "US",
    "state":   "Connecticut",
  }
  jsonString, err := json.Marshal(myMap)
  if err != nil {
    fmt.Fprintf(w, "error while decoding the map", err.Error())
  } else {
    fmt.Fprintf(w, string(jsonString))
  }
}

func outputSetAsJson(w http.ResponseWriter, req *http.Request) {
  // https://pkg.go.dev/net/url#Values
  // values := url.Values{}
  values := make(url.Values)
  values.Set("name", "Jake")
  values.Set("country", "US")
  values.Set("state", "Connecticut")

  jsonString, err := json.Marshal(values)
  if err != nil {
    fmt.Fprintf(w, "error while decoding the values", err.Error())
  } else {
    fmt.Fprintf(w, string(jsonString))
  }
}
```
