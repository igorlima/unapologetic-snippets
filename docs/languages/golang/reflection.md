---
layout: default
title: Reflection
nav_order: 3
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/reflection
---

<br/>
other samples:
- [a struct validator sample]({% link docs/languages/golang/reflection.validator-sample.md %})


-----


Reflection in Go allows you to examine and manipulate variables and types while your program is running. This means that you can check the type of a variable, change its value, or even call its methods.[^1]

The __reflect__ package has two main types, `reflect.Type` and `reflect.Value`, which allow you to examine the type and value of a variable.

`reflect.Type` represents the type of a value, while `reflect.Value` represents an instance of a value. Use the `reflect.TypeOf()` and `reflect.ValueOf()` functions to obtain the type and value of any value, respectively.

```golang
package main

import (
  "fmt"
  "reflect"
)

func main() {
  num := 123
  numType := reflect.TypeOf(num)
  numValue := reflect.ValueOf(num)
  fmt.Println("Type:", numType)
  fmt.Println("Value:", numValue)

  // output:
  // Type: int
  // Value: 123
}
```

use the `Kind()` method to determine the type of value that x is (in this case, an int) or use the `Int()` method to get the integer value of x. Additionally, to change the value of x, use the `Set()` method. [^1]

__Important:__ note that a `reflect.Value` holds a value, not the variable itself. If the value is a pointer to a struct, the `reflect.Value` will hold the pointer, not the value that the pointer points to.

```golang
package main

import (
  "fmt"
  "reflect"
)

func main() {
  x := 10
  p := &x
  v := reflect.ValueOf(p)
  fmt.Println("Kind of x:", v.Kind())

  v = v.Elem()
  fmt.Println("Kind of x:", v.Kind())
  fmt.Println("Value of x:", v.Int())
  v.SetInt(20)
  fmt.Println("Value of x after change:", x)

  // output:
  // Kind of x: ptr
  // Kind of x: int
  // Value of x: 10
  // Value of x after change: 20
}
```

The `reflect.Type` in Go is like a comprehensive guide to types, providing all the necessary information about a type, such as its name and composition. It acts as a template, providing all the information you need to know about the type. [^1]

Use `reflect.Type` to find information about a type, like its name, or to determine the type of a variable, such as struct, slice or pointer.
- `Name()`: retrieves the name of the type.
- `Kind()`: retrieves the underlying type.
- `Size()`: retrieves the size of the type in bytes.
- when working with fields:
  - `NumField()`: retrieves the number of fields in a struct type.
  - `Field(i int) reflect.StructField`: retrieves the i-th field in a struct type.
  - `FieldByName(name string) (reflect.StructField, bool)`: retrieves the field with the specified name in a struct type and a flag indicating its existence.

- when working with methods:
  - `NumMethod()`: retrieves the number of methods in a type.
  - `Method(i int) reflect.Method`: retrieves the i-th method in a type.
  - `MethodByName(name string) (reflect.Method, bool)`: retrieves the method with the specified name in a type.

- when working with __non-primitive__ types:
  - `Elem()`: retrieves the type of the element of an array, slice, pointer, or map type, it will cause an error if the `reflect.Type` is not of these types.
  - `Len()`: retrieves an array type’s length, it will cause an error if it is not an Array type.

__Note__ that using some of these methods in the wrong way *(e.g. calling `Elem()` on a __non-pointer value__)* will cause panic. [^1]

```golang
package main

import (
  "fmt"
  "reflect"
)

func main() {
  var car Car
  t := reflect.TypeOf(car)

  fmt.Println("Name:", t.Name())
  fmt.Println("Kind:", t.Kind())
  fmt.Println("Number of fields:", t.NumField())

  for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Println("Field name:", field.Name)
    fmt.Println("Field type:", field.Type)
  }

  // output:
  // Name: Car
  // Kind: struct
  // Number of fields: 3
  // Field name: Model
  // Field type: string
  // Field name: Year
  // Field type: int
  // Field name: EngineSize
  // Field type: float64
}

type Car struct {
  Model      string
  Year       int
  EngineSize float64
}
```

`reflect.Value` - `value.Interface()` `.Int()` `.String()...`: returns the value of a `reflect.Value` as an `interface{}` / `int` / `string`.

Be cautious when using this method as it may cause a panic if used improperly, for example, by calling `.Int()` on a string value type. [^1]

`value.Elem()` - used to retrieve the value that a `reflect.Value` interface contains. If the current value is a pointer, this method will dereference the pointer and return the value that it points to.

When using `reflect.ValueOf(x)` and then call `val.Elem()`, it will cause a panic because the `Elem()` method is not applicable to __non-pointer types__. The error message would be `“reflect: call of reflect.Value.Elem on int Value”`.

An alternative is to use the `Indirect()` function from the reflect package, which I will discuss further below.
- `value.Kind()`: returns the underlying type, or “Kind,” of a reflect.Value.
- `IsNil()`: checks if the value is nil and will panic if the value’s Kind is not one of the following types: Chan, Func, Interface, Map, Ptr, or UnsafePointer
- `Set(x Value)`: sets the value of “v” to “x”. However, it will cause a panic if the “CanSet()” method returns false or if the type of “x” is not compatible with the type of “v”

`reflect.Indirect()`: can be used as an alternative to “Elem()” when working with non-pointer types.

__`Zero(typ reflect.Type)`: returns a Value representing the zero value for the specified type.__ [^1]

```golang
package main

import (
  "fmt"
  "reflect"
)

func main() {
  x1 := 10
  xPtr := &x1

  val1 := reflect.ValueOf(xPtr)
  val1 = val1.Elem()
  fmt.Println("Value:", val1.Int())
  // output:
  // Value: 10

  x2 := 10
  val2 := reflect.ValueOf(x2)
  val2 = reflect.Indirect(val2)
  fmt.Println("Value:", val2.Int())
  // output:
  // Value: 10
}
```

---

[^1]: [Reflection in Go: Everything you need to know to use it effectively](https://levelup.gitconnected.com/reflection-in-go-everything-you-need-to-know-to-use-it-effectively-52c78da1f4ff)
