---
layout: default
title: Test
nav_order: 4
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/test
---

Writing basic tests in Go is:
- creating a file with the suffix `*_test.go`
- definig **a function** with the signature `TestXxx(t *testing.T)`
- running it using the [`go test` command]({% link docs/languages/golang/commands.md %}#go-test)

```sh
go mod init add
go mod tidy
```

```golang
// add.go
package add

func Add(a ...int) int {
  total := 0
  for i := range a {
    total += a[i]
  }
  return total
}
```

```golang
// add_test.go
package add

import "testing"

func TestAdd(t *testing.T) {
  // case 1
  if res := Add(1, 2, 3); res != 6 {
    t.Errorf("Expected %d instead of %d", 6, res)
  }

  // case 2
  if res := Add(-1, -2); res != -3 {
    t.Errorf("Expected %d instead of %d", -3, res)
  }
}
```

Further info:
  - [the command `go test`]({% link docs/languages/golang/commands.md %}#go-test)
  - [how to debug a go test]({% link docs/languages/golang/commands.test.debugging.md %})
