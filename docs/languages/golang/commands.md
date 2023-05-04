---
layout: default
title: Commands
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/commands
---

<br/>
<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

## go run

The `go run` command is used to compile and run a Go program in a single step. [^1][^2][^3]

```golang
package main
import "fmt"

func main() {
 fmt.Println("Hello, World!")
}
```

To run the program: `go run myprogram.go`.

## go build

The go build command is used to compile a Go program into an executable binary.

To compile the program: `go build myprogram.go`.

## go install

The `go install` command is used to compile a Go package and install it in the `$GOPATH` directory.

```golang
package main
import "fmt"
func main() {
 fmt.Println("Hello, World!")
}
```

To install the program: `go install`.

To run the installed program: `myprogram`.

## go test

The `go test` command is used to run tests for a Go package.

```golang
package main
import "testing"
func TestMyProgram(t *testing.T) {
 result := MyProgram()
 if result != "Hello, World!" {
 t.Errorf("Expected 'Hello, World!', but got '%s'", result)
 }
}
```

To run the tests: `go test`.

- further info:
   - [the basic]({% link docs/languages/golang/test.md %})
   - [how to debug a go test]({% link docs/languages/golang/commands.test.debugging.md %})


## go get

The `go get` command is used to download and install packages from remote repositories.

```sh
go get github.com/gin-gonic/gin
```

After the package is installed, you can import it into your Go code by adding the following line to your source code:

```golang
import "github.com/gin-gonic/gin"
```


- _further info:_
   - [how to fetch modules from private repository]({% link docs/languages/golang/commands.get.private-repo.md %})


## go fmt

The `go fmt` command is used to format Go code according to the standard Go formatting rules.

```sh
go fmt myprogram.go
```

## go vet

The `go vet` command is used to check Go code for common mistakes and errors.

## go mod

The `go mod` command is used to manage dependencies in Go modules, which are a way of organizing and sharing code dependencies in Go.

Use the `go mod init` command to initialize a new Go module:

```sh
go mod init myproject
```

Next, you can use the `go mod tidy` command to download and manage the dependencies of your module:

```sh
go mod tidy
```

## go doc

The `go doc` command is used to generate documentation for Go packages.

You can also use the `go doc` command to view documentation for specific functions or other elements within a package.

```sh
go doc mypackage.HelloWorld
```


- _further info:_
   - [writing documentation in golang]({% link docs/languages/golang/commands.doc.md %})


## go generate

The `go generate` command is used to run code generators in Go.

## golint

The `golint` command is a tool for checking Go code for style and correctness issues.

----

[^1]: [Useful Go Commands](https://njkhanh.com/useful-go-commands-p5f3234343234)
[^2]: [Learn Go: Useful Commands](https://medium.com/@karan99/learn-go-useful-commands-7087dfc05e44)
[^3]: [Useful Commands](https://www.karanpratapsingh.com/courses/go/useful-commands)
