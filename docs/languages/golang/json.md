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
- [decode string JSON to int]({% link docs/languages/golang/json.decode-string-json-to-int.md %})
  - `json: cannot unmarshal string into Go struct field Item.item_id of type int`

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

## Marshalling and Unmarshalling

What exactly is marshalling and unmarshalling? In simple terms:
- Marshalling is the process of converting a Go data structure into a format that can be stored or transmitted, like JSON or XML.
- Unmarshalling is the reverse - taking data in a format like JSON and converting it back into a Go data structure.

It's kinda like packing luggage for a trip (*__marshalling__*) and then unpacking when you get home (*__unmarshalling__*). Isn't?

<details markdown="block">
  <summary>
    👇 a simple sample...
  </summary>

This is where marshalling comes in handy!
```golang
package main

import (
  "encoding/json"
  "fmt"
)

type Person struct {
  Name string
  Age  int
}

func main() {
  p := Person{"rahul", 12}
  jsondata, _ := json.Marshal(p)

  fmt.Println(string(jsondata)) // Json Data:=> {"Name":"rahul","Age":12}
  fmt.Println(jsondata)         // Marshal Data:=> [123 34 78 97 109 101 34 58 34 114 97 104 117 108 34 44 34 65 103 101 34 58 49 50 125]

}
```

Unmarshalling is equally easy
```golang
var unmarshalledP Person
json.Unmarshal(jsonData, &unmarshalledP)
fmt.Println(unmarshalledP)
```

*Marshalling and unmarshalling use the concept of __encoders__ and __decoders__ in Go. An encoder writes structured Go data into a standard format like JSON. A decoder reads data in a standard format and converts it back into Go data structures.*

------

</details>

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

