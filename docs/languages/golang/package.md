---
layout: default
title: Package
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/package
---

Other examples:
- [go get]({% link docs/languages/golang/commands.md %}#go-get)
- [how to fetch private repository]({% link docs/languages/golang/commands.get.private-repo.md %})


# Package

- Golang Version: 1.20
  - In Go, one package == one directory. That is, all `.go` files for a package should be contained in the same directory, and a directory should contain the `.go` files for one package only. You shouldn't ever have `.go` files with different package names in the same directory.
  - For all non-main packages, the directory name that the code lives in should be the same as the package name. When choosing a name you should pick something that is short, descriptive, lower case and ideally one word.
  - The Go blog has a helpful post with additional guidance and some examples of good and bad names:
    - [https://go.dev/blog/package-names](https://go.dev/blog/package-names)

<details markdown="block">
  <summary>
    Go package sample
  </summary>


```
 |-package-main.go
 |-greeting
 | |-hello.go
 | |-hi.go
 |-version
 | |-get-version.go
 |-random
 | |-a.go
```

```sh
mkdir -p greeting version random

cat << "EOF" > package-main.go
paste the code here
EOF

cat << "EOF" > greeting/hello.go
paste the code here
EOF

cat << "EOF" > greeting/hi.go
paste the code here
EOF

cat << "EOF" > version/get-version.go
paste the code here
EOF

cat << "EOF" > random/a.go
paste the code here
EOF
```

```golang
// package-main.go
package main

/*
Golang Version: 1.20

The Go blog has a helpful post with additional guidance and
some examples of good and bad names: -
https://go.dev/blog/package-names -

In Go, one package == one directory. That is, all `.go`
files for a package should be contained in the same
directory, and a directory should contain the `.go` files
for one package only. You shouldn't ever have `.go` files
with different package names in the same directory.

For all non-main packages, the directory name that the code
lives in should be the same as the package name. When
choosing a name you should pick something that is short,
descriptive, lower case and ideally one word.
*/

import (
  foo "foo"
  greeting "package-example/greeting"
  _ "package-example/version"
)

func main() {
  foo.RandomA()
  greeting.Hi()
  greeting.Hello()
}

/*
OUTPUT

version/get-version.go ==> getVersion()
1.0.0
aaaaa
Hi!
Hello!
*/
```

```golang
// greeting/hello.go
package greeting

import "fmt"

func Hello() {
  fmt.Println("Hello!")
}
```

```golang
// greeting/hi.go
package greeting

import "fmt"

func Hi() {
  fmt.Println("Hi!")
}
```

```golang
// version/get-version.go
package version

import "fmt"

func init() {
  fmt.Println(getVersion())
}
func getVersion() string {
  fmt.Println("version/get-version.go ==> getVersion()")
  return "1.0.0"
}
```

```golang
// random/a.go
package random

import "fmt"

func RandomA() {
  fmt.Println("aaaaa")
}
```

```sh
# golang version: 1.20

go mod init package-example
go mod tidy

# create a sub package inside the current package
(cd random && \
go mod init package-random && \
go mod tidy)

# create an alias for the sub package above
# use the relative package path above
go mod edit -replace foo=./random
go get foo
go mod tidy
```

<br/>
</details>
<br/>

----

