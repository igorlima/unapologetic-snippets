---
layout: default
title: JSON
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/json
---

other resources:
- [http server sample converting and responding in JSON]({% link docs/languages/golang/api.http.md %})
- [http client and http server interacting through JSON]({% link docs/languages/golang/api.http-client-server.md %})

<br/>

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

----

## Pretty Printing JSON Structs in Go

```golang
package main

import (
  "fmt"
  "encoding/json"
)

type Flight struct {
  Origin string `json:"origin"`
  Destination string `json:"destination"`
  Price int `json:"price"`
}

func main() {
  flight := Flight{
    Origin: "GLA",
    Destination: "JFK",
    Price: 2000,
  }

  b, err := json.MarshalIndent(flight, "", "  ")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Print(string(b))
}
```

## Nested map of map

```golang
package main

import (
  "encoding/json"
  "fmt"
)

func main() {
  mapA := map[string]string{
    "63": "Abalone",
    "48": "Açaí berries",
    "49": "Açaí juice",
    "42": "Acorn squash",
    "50": "Adzuki bean paste",
    "51": "Adzuki beans",
    "46": "Aged Japanese kurozu",
  }
  mapB := map[string]string{
    "63": "Bacon",
    "48": "Bananas",
    "49": "Beef",
    "42": "Black tea",
    "50": "Blackberries",
    "51": "Blueberries",
    "46": "Butter",
  }
  mapOfMap := map[string]map[string]string{}

  for id, foodA := range mapA {
    mapOfMap[id] = map[string]string{}
    mapOfMap[id]["A"] = foodA
    mapOfMap[id]["B"] = mapB[id]
  }

  /*
    for id, foodB := range mapB {
      mapOfMap[id]["B"] = foodB
    }
  */

  fmt.Println(mapOfMap)
  json, err := json.MarshalIndent(mapOfMap, "", "  ")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(json))

  fmt.Println(mapOfMap["46"]["A"])
  fmt.Println(mapOfMap["46"]["B"])
}
```
