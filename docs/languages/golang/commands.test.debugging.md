---
layout: default
title: Debug Go Test
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/commands-test-debugging
---

# Debugging go test

__[back]({% link docs/languages/golang/commands.md %}#go-test)__
<br/>

As it turns out, the go test command is completely configurable with all linker, loader, and runtime flags. For example, you can change your GOOS and GOARCH environment variables to compile different test files.

- compile a test binary:
   - `go test -c`
   - `go test -c -- switchers_test.go`
- run that binary through a debugger
   - with a compiled file:
      - `dlv exec switchers.test`
   - without a compiled file:
      - `dlv test -- switchers_test.go`

```sh
go mod init switchers
go mod tidy
```

```golang
// switchers.go
package switchers

func SwitchFunction(a, b int, c *int) string {
  switch *c {
  case a:
    return "a"
  case b:
    return "b"
  default:
    return "c"
  }
}
```

```golang
// switchers_test.go
package switchers

import (
  "testing"
)

func TestSwitch(t *testing.T) {

  var (
    a int
    b int
    c = &b
  )

  x := SwitchFunction(a, b, c)

  if x != "c" {
    t.Error("wtf?")
  }
}
```

```
OUTPUT
--- FAIL: TestSwitch (0.00s)
    switchers_test.go:19: wtf?
FAIL
exit status 1
FAIL    switchers       0.814s
```

