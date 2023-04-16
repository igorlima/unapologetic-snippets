---
layout: default
title: Writing Documentation for Golang
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/commands-doc
---

# Writing Documentation for Golang
__[back]({% link docs/languages/golang/commands.md %}#go-doc)__
<br/>

To use `go doc`, it's needed to include comments in the code that are formatted in a specific way. These comments should be placed immediately above the code that they describe.

```sh
go mod init add
```

```golang
// Package math provides basic math operations.
// It has only one file:
//   - add.go
package add

// Add adds integers and returns the result.
//
// Example:
//
//  sum := add.Add(1, 2)
//  fmt.Println(sum) // Output: 3
//  sum := add.Add(1, 2, 3)
//  fmt.Println(sum) // Output: 6
func Add(a ...int) int {
  total := 0
  for i := range a {
    total += a[i]
  }
  return total
}
```

```sh
go doc add
```
```
package add // import "add"

Package math provides basic math operations. It has only one file:
  - add.go

func Add(a ...int) int
```

```sh
go doc add.Add
```
```
package add // import "add"

func Add(a ...int) int
    Add adds integers and returns the result.

    Example:

        sum := add.Add(1, 2)
        fmt.Println(sum) // Output: 3
        sum := add.Add(1, 2, 3)
        fmt.Println(sum) // Output: 6
```


```sh
# as of 1.18, it should be installed via:
# `go get` in 1.18 and beyond will no longer install executables.
go get golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/godoc

godoc -http=:6060
open http://localhost:6060
```



