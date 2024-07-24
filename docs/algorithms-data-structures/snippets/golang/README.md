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

### Zero Values

Each type has a zero value. The zero value of a type can be viewed as the
default value of the type.
- The zero value of a boolean type is false.
- The zero value of a numeric type is zero, though zeros of different numeric
  types may have different sizes in memory.
- The zero value of a string type is an empty string.”

## String

### String value literals

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

### String Concatenation

```go
println("Go" + "lang") // Golang
var a = "Go"
a += "lang"
println(a) // Golang
```

### `fmt.Printf` Format Verbs

A list if format verbs:
- `%v`, which will be replaced with the general string representation of the
  corresponding argument.
- `%T`, which will be replaced with the type name or type literal of the
  corresponding argument.
- `%x`, which will be replaced with the hex string representation of the
  corresponding argument.  Note, the hex string representations for values of
  some kinds of types are not defined.  Generally, the corresponding arguments
  of `%x` should be strings, integers, integer arrays or integer slices (arrays
  and slices will be explained in a later article).
- `%s`, which will be replaced with the string representation of the
  corresponding argument. The corresponding argument should be a string or byte
  slice.
- Format verb `%%` represents a percent sign.


## Operators

### Boolean (a.k.a. Logical) Operators

| Operator   | Name                                        | Requirements for Operand(s)                                    |
| ---------- | ------                                      | ------------------------------                                 |
| `&&`       | boolean and (binary) a.k.a. conditional and | The two operands must be both values of the same boolean type. |
| `||`       | boolean or (binary) a.k.a. conditional or   | The two operands must be both values of the same boolean type. |
| `!`        | boolean not (unary)                         | The type of the only operand must be a boolean type.           |

### Comparison Operators

| Operator | Name                    | Requirements for the Two Operands                                                                  |
| -------- | ------                  | ------------------------------                                                                     |
| `==`     | equal to                | Generally, the types of its two operands must be the same.                                         |
| `!=`     | not equal to            |                                                                                                    |
| `<`      | less than               | The two operands must be both values of the same integer type, floating-point type or string type. |
| `<=`     | less than or equal to   |                                                                                                    |
| `>`      | larger than             |                                                                                                    |
| `>=`     | larger than or equal to |                                                                                                    |

## Functions

### Built-in Functions

There are some built-in functions in Go, for example, the `println` and `print`
functions. We can call these functions without importing any packages.

## Expressions, Statements

Simply speaking, an expression represents a value and a statement represents an
operation.

## Control Flows

There are three kinds of basic control flow code blocks in Go:
- `if-else` two-way conditional execution block.
- `for` loop block.
- `switch-case` multi-way conditional execution block.

### `if-else` Control Flow Blocks

The full form of a if-else code block is like:

```go
if InitSimpleStatement; Condition {
   // do something
} else {
   // do something
}
```

The `Condition` portion can be enclosed in a pair of `()` or not, but it can't
be enclosed together with the InitSimpleStatement portion.  If the
InitSimpleStatement is absent, then the semicolon following it is optional.


### `for` Loop Control Flow Blocks

The full form of a for loop block is:
```go
for InitSimpleStatement; Condition; PostSimpleStatement {
   // do something
}
```

### `for-range` Control Flow Blocks

The following code:
```go
for i := range anInteger {
   ...
}
```

is actually a short form of
```go
for i := 0; i < anInteger; i++ {
   ...
}
```

### `switch-case` Control Flow Blocks

The full form a switch-case block is like:
```go
switch InitSimpleStatement; CompareOperand0 {
case CompareOperandList1:
   // do something
case CompareOperandList2:
   // do something
...
case CompareOperandListN:
   // do something
default:
   // do something
}
```

Go provides a `fallthrough` keyword to do this task. For example, in the
following example, every branch code block will get executed, by their orders,
from top to down.

Note:
- a fallthrough statement must be the final statement in a branch.
- a fallthrough statement can't show up in the final branch in a switch-case
  control flow block.

Another obvious difference from many other languages is the order of the
`default` branch in a switch-case control flow block can be arbitrary.

### `goto` Statement

A `goto` keyword must be followed by a label to form a statement. A label is
declared with the form `LabelName:`, where `LabelName` must be an identifier. A
label which name is not the blank identifier must be used at least once.

__`break` and `continue` Statements With Labels__: a `break` or `continue`
statement can also contain a label, but the label is optional. Generally,
`break` containing labels are used in nested breakable control flow blocks and
`continue` statements containing labels are used in nested loop control flow
blocks.

