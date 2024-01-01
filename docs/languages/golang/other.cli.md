---
layout: default
title: Other CLI
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other-cli
---

__[back]({% link docs/languages/golang/other.md %})__

# Command Line Interface (CLI)

## Creating CLI the Traditional Way

```golang
package main

import (
  "flag"
  "fmt"
  "time"
)

func main() {
  command := flag.String("command", "", "Command to execute")
  flag.Parse()

  if *command == "" {
    fmt.Println("Please provide a command.")
    flag.Usage()
    return
  }

  switch *command {
  case "greet":
    fmt.Println("Hello, World!")
  case "time":
    fmt.Println("Current time:", time.Now().Format("15:04:05"))
  default:
    fmt.Println("Unknown command:", command)
  }
}

/*
./main -command greet
*/
```

## Libraries

- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper)
