---
layout: default
title: Test
nav_order: 4
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/test
---

- further info:
  - [the command `go test`]({% link docs/languages/golang/commands.md %}#go-test)
  - [how to debug a go test]({% link docs/languages/golang/commands.test.debugging.md %})

Writing basic tests in Go is:
- creating a file with the suffix `*_test.go`
- definig **a function** with the signature `TestXxx(t *testing.T)`
- running it using the [`go test` command]({% link docs/languages/golang/commands.md %}#go-test)
