---
layout: default
title: Other Env Var
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other-env-var
---

__[back]({% link docs/languages/golang/other.md %})__

# Get value from an environment variable

```golang
package main

import (
  "flag"
  "fmt"
  "os"
)

func main() {
  // In Go, it’s dead simple to get the value from an environment variable
  fmt.Println(os.Getenv("HOME"))
  fmt.Println(os.Getenv("FOO"))

  argsWithProg := os.Args
  argsWithoutProg := os.Args[1:]
  fmt.Println(argsWithProg)
  fmt.Println(argsWithoutProg)

  greetPtr := flag.String("greet", "default value", "a string for description")
  greetingPtr := flag.String("greeting", "default value", "a string for description")
  flag.Parse()
  fmt.Printf("greet: %+v: \n", *greetPtr)
  fmt.Printf("greeting: %+v: \n", *greetingPtr)

  fmt.Println("vim-go")
}

// OUTPUT
/*
  /Users/ilima
  bar
  [/var/folders/3x/2w12g8693mx0mgdvcwn50tr00000gq/T/go-build2921396100/b001/exe/main --greet=hello --greeting hi]
  [--greet=hello --greeting hi]
  greet: hello:
  greeting: hi:
  vim-go
*/
```


```bash
(FOO=bar go run main.go --greet=hello --greeting hi)
```
