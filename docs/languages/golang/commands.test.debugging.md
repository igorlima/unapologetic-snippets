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
- to run a single unit test:
   - `dlv test -- -test.run ^TestSwitchC$`
      - once the debugger builds the test binary and launches, then set breakpoints relative to the start of the function.
      - witin the [go-vim](https://github.com/fatih/vim-go) plugin allows the same functionality using the `:GoDebugTestFunc` command with the cursor on the required function.


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

func TestSwitchC(t *testing.T) {
  var (
    a int
    b int
    c int = 8
  )

  x := SwitchFunction(a, b, &c)
  if x != "c" {
    t.Error("wtf?")
  }
}
```

```
OUTPUT
--- FAIL: TestSwitch (0.00s)
    switchers_test.go:17: wtf?
FAIL
exit status 1
FAIL    switchers       1.358s
```

## Setting breakpoints

To debug a specific test, it's needed to set a breakpoint on that test function.

To print list of functions.
```
funcs [<regex>]
```

Discover the test functions with the following command:

```
(dlv) funcs switchers.*
```

With the provided list of test on hand, then set a breakpoint on a specific test:

```
(dlv) break switchers.TestSwitch
(dlv) break switchers.SwitchFunction
```

To print all local variables:
```
(dlv) locals
```

To print a specific local, use the print (p) command:
```
p a
p b
p c
p x
```

