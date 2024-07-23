# Golang

Makefile for Golang:
- https://igorlima.github.io/unapologetic-snippets/docs/languages/golang/other-makefile
  - Download Makefile:
    - `https://makefile-go-gist.ilima.xyz/`
    - `curl -L https://makefile-go-download.ilima.xyz -o Makefile`

## Keywords

Keywords are predefined, reserved identifiers that have special meanings to the
compiler. Go - up to now version 1.22 - has 25 keywords. They can be
categorized as four groups:

- `const`, `func`, `import`, `package`, `type`, and `var` are used to declare
  all kinds of identifiers.
- `chan`, `interface`, `map`, and `struct` are used to declare composite types.
- `break`, `case`, `continue`, `default`, `else`, `fallthrough`, `for`, `goto`,
  `if`, `range`, `return`, `select`, and `switch` are used to control the flow
  of a program.
- `defer`, and `go` to are also used to control the flow of a program, but in
  other specific ways. They modify function calls.

## Built-in Basic Types

Go - up to now version 1.22 - supports following built-in basic types:
- one boolean built-in boolean type: `bool`.
- 11 built-in integer numeric types (basic integer types): `int8`, `uint8`,
  `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `int`, `uint`, and
  `uintptr`.
- two built-in floating-point numeric types: `float32` and `float64`.
- two built-in complex numeric types: `complex64` and `complex128`.
- one built-in string type: `string`.

## Zero Values

Each type has a zero value. The zero value of a type can be viewed as the
default value of the type.
- The zero value of a boolean type is false.
- The zero value of a numeric type is zero, though zeros of different numeric
  types may have different sizes in memory.
- The zero value of a string type is an empty string.”

## String value literals

There are two forms of string value literals, interpreted string literal
(double quotes form) and raw string literal (back quotes form). For example,
the following two string literals are equivalent:

```
// the interpreted form.
"Hello\nworld!\n\"你好世界\""

// the raw form.
`Hello
world!
"你好世界"`
```

In a raw string literal, no character sequences will be escaped. The back quote
character is not allowed to appear in a raw string literal.

