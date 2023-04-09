---
layout: default
title: Golang
nav_order: 6
parent: Programming Languages
has_children: true
permalink: /docs/languages/golang
---

Other examples:
- [from unknown to expert: golang statements]({% link docs/languages/golang/index.2023a04m04d.md %}) [^1]

# Golang

The `$GOPATH` environment variable lists places for Go to look for Go Workspaces. By default, Go assumes the `GOPATH` location is at `$HOME/go`.

`GOPATH` is not `$GOROOT` - the `$GOROOT` is where Go's code, compiler, and tolling lives - this is not our source code. The `GOROOT` is usually something like `/usr/local/go`. The `$GOPATH` is usually something like `$HOME/go`. Inside of a Go Workspace, or `GOPATH`, there are three directories: `bin`, `pkg`, and `src`.

The `$GOPATH/pkg` directory is where Go stores pre-compiled object files to speed up subsequent compiling of programs.

If the code lives at `$GOPATH/src/blue/red` then its package name should de `red`. The import statemnet for the red package would be: `import "blue/red"`.

Strings exist within either back quotes `` ` `` or double quotes `"` in Go and have different characteristics depending on which quotes is used. If the usage is back quotes, then a raw string literal is created. If the usage is double quotes, then an interpreted string literal is created.

__Zero Values__. All built-in types have a zero value. Any allocated variable is usable even if it never has a value assigned. In Go, because all values have a zero value, there's no way to have `undefined` value like some other languages.

```golang
package main

import "fmt"

func main() {
  var a int
  var b string
  var c float64
  var d bool

  // the %T tells the function to print the data type for the variable
  fmt.Printf("var a %T = %+v\n", a, a)
  fmt.Printf("var b %T = %+v\n", b, b)
  fmt.Printf("var c %T = %+v\n", c, c)
  fmt.Printf("var d %T = %+v\n", d, d)
}
```

If a variable starts with an uppercase letter, then that variable is accessible outside the package it was declared in (or exported). If a variable starts with a lowercase letter, then it is only available within the package it is declared in.o

Global variables exist outside of functions. Local variables exist within functions.o

The standard library provides two built-in functions to create errors: `errors.New` and `fmt.Errorf`. Both of these functions allow to specify a custom error message that can later present to the users.

```golang
errA := errors.New("barnacles")
errB := fmt.Errorf("error occurred at: %v", time.Now())
```

Errors that a program encounters fall into two broad categories: those the programmer has anticipated and those the programmer has not. Panics fall into the second category of errors, which are unanticipated by the programmer. Also, use `defer` statements along with the `recover` function to capture panics before they have a chance to unexpectedly terminate the Go program running.

```golang
package main

import "fmt"

func main() {
  defer func() {
    fmt.Println("hello from the deferred function!")
  }()

  panic("oh no!")
}
```

```golang
package main

import "fmt"

func main() {
  defer func() {
    fmt.Println("hello from the deferred function!")
    if err := recover(); err != nil {
      fmt.Println("panic occurred:!", err)
    }
    fmt.Println("recovered")
  }()

  panic("oh no!")
}
```

To modify the names of packages and their functions within Go is by putting alias name in front of the imported package. The construction of this statement looks like this: `import another_name "package"`

Go modules are configuration files placed in the package directory that tell the compiler where to import packages from.

```
# go.mod
module github.com/gopherguides/cmd

replace github.com/gopherguides/logging => ../logging
```

A variadic function is a function that accepts zero, one, or more values as a single argument. A function with a parameter that is preceded with a set of ellipses `(...)` is considered a variadic function. It's only possible to have one variadic parameter in a function, and it must be the last parameter defined in the function.

In go, the predefined `init()` function sets off a piece of code to run before any other part of the package. This code will execute as soon as the package is imported, and can be used when the application needs to initialize in a specific state, such as when it's needed to have a specific configuration or set of resources with which the application needs to start. Per the Go language specification for Package Initialization, when multiple files are encountered in a package, they are processed alphabetically.

In Go, a build tag, or a build constraint, is an identifier added to a piece of code that determines when the file should be included in a package during the build process. This allows to build different version of Go application from the same source code and to toggle between them in a fast and organized manner. Many developers use build tags to improve the workflow of building cross-platform compatible applications, such as programs that require code changes to account for variances between different operating systems. Build tags are also used for integration testing, allowing to quickly switch between the integrated code and the code with a mock service or stub, and for differing levels of feature sets within an application.

The `+build` declaration tells the `go build` command that this isn't a comment, but instead is a build tag.

```golang
// +build pro enterprise
package main
```

When putting multiple build tags on the same line a `.go` file, `go build` interprets them as using `OR` logic. Instead of putting both tags on the same line, put them on separate lines, then `go build` will interpret those tags using `AND` logic.

```golang
// +build pro
// +build enterprise
package main
```

```bash
go build -tags "enterprise pro"
```

When using a pointer to a variable, there are a couple of different syntax elements that needs to understand. The first one is the use of the ampersand `&`. An ampersand in front of a variable name states that the desire in to get the `address`, or a pointer to that variable. The second syntax element is the use of the asterisk `*` or dereferencing operator. When declaring a point variable, follow the variable name with the type of the variable that the pointer points to, prefixed with an `*`: `var myPointer *int32 = &someint`. The memory address is a hexadecimal number, and not meant to be human-readable.

Nil Pointers. All variables in Go have a zero value. This is true even for a pointer. When declaring a pointer to a type, but assign no value, the zero value will be `nil`. `nil` is a way to say that "nothing has been initialized" for the variable.

In addition to defining a new type to represent a struct, it's also possible to define an inline struct.

```golang
package main

import "fmt"

func main() {
  c := struct {
    Name string
    Type string
  }{
    Name: "Sammy",
    Type: "Shark",
  }
  fmt.Println(c.Name, "the", c.Type)
}
```

Pointer Receivers. The syntax for defininf methods on the pointer receiver is nearly identical to defininf methods on the value receiver. The difference is prefixing the name of the type in the receiver declaration with an asterisk `(    *)`. The following example defines a method on the pointer receiver to a type:

```golang
package main

import "fmt"

func main() {
  b := &Boat{
    Name: "Hey Ho",
  }
  b.Add("Sammy the Shark").Add("Larry the Lobster")
  b.Print()
}

type Boat struct {
  Name      string
  occupants []string
}

func (b *Boat) Add(name string) *Boat {
  b.occupants = append(b.occupants, name)
  return b
}

func (b Boat) Print() {
  fmt.Println("The", b.Name, "has the following occupants:")
  for _, n := range b.occupants {
    fmt.Println("\t", n)
  }
}
```

The Go tooling has command that can print a list of the possible platforms that Go can build on. This list can change with each new Go release, so the combinations might not be the same on different version of Go. To find this list of possible platforms, run the following: `go tool dist list`

When running a command like `go build`, Go uses the current platform's `GOOS` and `GOARCH` to determine how to build the binary. To find out what combination the platform is, use the `go env` command and pass `GOOS` and `GOARCH` as arguments: `go env GOOS GOARCH`.

To build a Go program for a different operating systems and architecture, run something like: `GOOS=linux GOARCH=ppc64 go build` 


----

[^1]: [From Unknown to Expert: Golang Statements You Need to Know](https://blog.canopas.com/from-unknown-to-expert-golang-statements-you-need-to-know-a1ac97213c04)
