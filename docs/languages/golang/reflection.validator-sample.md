---
layout: default
title: Reflection - Struct Validador Sample
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/reflection-validator-sample
---

__[back]({% link docs/languages/golang/reflection.md %})__

```golang
package main

import (
  "fmt"
  "net/url"
  "reflect"
)

type Message struct {
  Id      string `json:"id"`
  From    string `json:"from"`
  To      string `json:"to"`
  Subject string `json:"subject"`
  Content string `json:"content"`
}

func main() {
  fmt.Println(ValidateFields(Message{
    Id: "123",
  }))
  fmt.Println(ValidateFields(Message{
    Id: "123",
  }, []string{"Id", "Foo"}...))
  fmt.Println(ValidateFields(Message{
    Id: "123",
  }, []string{"Id"}...))
  fmt.Println("vim-go")

  /**
  OUTPUT

  map[From:[missing From]]
  map[Foo:[field does NOT exist Foo]]
  map[]
  */
}

// ValidateFields returns errors with the RequestBody
func ValidateFields(entity any, fields ...string) url.Values {
  errs := url.Values{}
  if len(fields) == 0 {
    fields = []string{"Id", "From"}
  }
  entityType, entityValue := reflect.TypeOf(entity), reflect.ValueOf(entity)
  for _, name := range fields {
    _, exists := entityType.FieldByName(name)
    if !exists {
      errs.Add(name, fmt.Sprintf("field does NOT exist: %v", name))
      continue
    }

    field := entityValue.FieldByName(name)
    valueStr := field.String()
    typeStr := field.Type().String()
    if typeStr == "string" && valueStr == "" {
      errs.Add(name, fmt.Sprintf("missing %v", name))
    }
  }
  return errs
}
```
