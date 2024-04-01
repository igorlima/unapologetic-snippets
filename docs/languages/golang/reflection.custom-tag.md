---
layout: default
title: Reflection - Custom Tag
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/reflection-custom-tag
---

__[back]({% link docs/languages/golang/reflection.md %})__

```golang
package main

import (
  "crypto/sha1"
  "crypto/sha512"
  "encoding/hex"
  "fmt"
  "reflect"
)

// how to define a custom tag in Struct
func main() {
  fmt.Println("vim-go")

  // define a struct with a custom tag
  type Person struct {
    Name string `encode:"sha1"`
    Age  int    `encode:"sha512"`
  }

  // create an instance of the Person struct
  p := Person{
    Name: "John",
    Age:  25,
  }

  // get the type of the struct
  t := reflect.TypeOf(p)

  // iterate over ther fields of the struct
  for i := 0; i < t.NumField(); i++ {
    // get the field
    f := t.Field(i)

    // print the field name and tag
    tagName := f.Tag.Get("encode")
    if tagName == "sha1" {
      h := sha1.New()
      h.Write([]byte(p.Name))
      fmt.Printf("Field: %s, Encoded: %s\n", f.Name, hex.EncodeToString(h.Sum(nil)))
    } else if tagName == "sha512" {
      h := sha512.New()
      h.Write([]byte(p.Name))
      fmt.Printf("Field: %s, Encoded: %s\n", f.Name, hex.EncodeToString(h.Sum(nil)))
    }
  }
}
```
