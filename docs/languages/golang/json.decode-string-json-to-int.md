---
layout: default
title: Decode string JSON to int
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/json-decode-string-json-to-int
---

__[back]({% link docs/languages/golang/json.md %})__

-----

Sometimes a client-side application could send a valid integer as a string
while the application expects an integer. Unfortunately, that leads to an error
when converting the JSON into a struct.
- `json: cannot unmarshal string into Go struct field Item.item_id of type int`

For example consider the snippet below [^1]:

```golang
package main

import (
  "encoding/json"
  "fmt"
  "log"
)

func main() {
  type Item struct {
    Name   string `json:"name"`
    ItemId int    `json:"item_id"`
  }
  jsonData := []byte(`{"name":"item 1","item_id":"30"}`)
  var item Item
  err := json.Unmarshal(jsonData, &item)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("vim-go")
}
```

The expectation is the application converting the string "30" to an integer 30
since it's a valid integer.

However, what happens is the following error instead:
- `json: cannot unmarshal string into Go struct field Item.item_id of type int`

How to solve this issue, though?

The good news is that Golang allows writing a custom JSON _marshaller_ and
_unmarshalers_. It allows creating a new type alias for the int type and adding
a custom _marshaller_ and _unmarshaler_ for the JSON; this way, it can check
the data and convert the string to int before _unmarshal_. [^2]

```golang
// StringInt create a type alias for type int
type StringInt int

// UnmarshalJSON create a custom unmarshal for the StringInt
// this helps us check the type of our value before unmarshalling it

func (st *StringInt) UnmarshalJSON(b []byte) error {
  // convert the bytes into an interface
  // this will help us check the type of our value
  // if it is a string that can be converted into a int we convert it
  // otherwise we return an error
  var item interface{}
  if err := json.Unmarshal(b, &item); err != nil {
    return err
  }
  switch v := item.(type) {
  case int:
    *st = StringInt(v)
  case float64:
    *st = StringInt(int(v))
  case string:
    // here convert the string into
    // an integer
    i, err := strconv.Atoi(v)
    if err != nil {
      // the string might not be of integer type
      // so return an error
      return err
    }
    *st = StringInt(i)
  }
  return nil
}

type Item struct {
  Name   string    `json:"name"`
  ItemId StringInt `json:"item_id"`
}
```

-----

<details markdown="block">
  <summary>
    a full code snippet
  </summary>

A full code snippet below [^2]:

```golang
package main

import (
  "encoding/json"
  "fmt"
  "log"
  "strconv"
)

// StringInt create a type alias for type int
type StringInt int

// UnmarshalJSON create a custom unmarshal for the StringInt
// this helps us check the type of our value before unmarshalling it

func (st *StringInt) UnmarshalJSON(b []byte) error {
  // convert the bytes into an interface
  // this will help us check the type of our value
  // if it is a string that can be converted into a int we convert it
  // otherwise we return an error
  var item interface{}
  if err := json.Unmarshal(b, &item); err != nil {
    return err
  }
  switch v := item.(type) {
  case int:
    *st = StringInt(v)
  case float64:
    *st = StringInt(int(v))
  case string:
    // here convert the string into
    // an integer
    i, err := strconv.Atoi(v)
    if err != nil {
      // the string might not be of integer type
      // so return an error
      return err
    }
    *st = StringInt(i)
  }
  return nil
}

type Item struct {
  Name   string    `json:"name"`
  ItemId StringInt `json:"item_id"`
}

func main() {
  fmt.Println("-----------------")
  printStruct([]byte(`{"name":"item 1","item_id":"30"}`))
  fmt.Println("-----------------")
  printStruct([]byte(`{"name":"item 2","item_id":40}`))
  fmt.Println("-----------------")

  var item Item
  jsonData := []byte(`{"name":"item 3","item_id":"50"}`)
  err := json.Unmarshal(jsonData, &item)
  if err != nil {
    log.Fatal(err)
  }
  printIndentedJSON(jsonData)
  fmt.Printf("%+v\n", item)
  itemId := int(item.ItemId)
  fmt.Printf("%+v\n", itemId)

  fmt.Println("-----------------")
  jsonData = []byte(`{"name":"item 4","item_id":60}`)
  err = json.Unmarshal(jsonData, &item)
  if err != nil {
    log.Fatal(err)
  }
  printIndentedJSON(jsonData)
  fmt.Printf("%+v\n", item)
  itemId = int(item.ItemId)
  fmt.Printf("%+v\n", itemId)

  fmt.Println("-----------------")

  /** OUTPUT
  -----------------
  {
    "item_id": "30",
    "name": "item 1"
  }
  {Name:item 1 ItemId:30}
  30
  -----------------
  {
    "item_id": 40,
    "name": "item 2"
  }
  {Name:item 2 ItemId:40}
  40
  -----------------
  {
    "item_id": "50",
    "name": "item 3"
  }
  {Name:item 3 ItemId:50}
  50
  -----------------
  {
    "item_id": 60,
    "name": "item 4"
  }
  {Name:item 4 ItemId:60}
  60
  -----------------
  */
}

func printStruct(jsonData []byte) Item {
  var item Item
  err := json.Unmarshal(jsonData, &item)
  if err != nil {
    log.Fatal(err)
  }
  printIndentedJSON(jsonData)
  fmt.Printf("%+v\n", item)
  itemId := int(item.ItemId)
  fmt.Printf("%+v\n", itemId)
  return item
}

func printIndentedJSON(jsonData []byte) {
  var obj map[string]interface{}
  json.Unmarshal(jsonData, &obj)
  str, _ := json.MarshalIndent(obj, "", "  ")
  fmt.Printf("%+v\n", string(str))
}
```

<br/>
</details>

[^1]: [Code sahara - Golang Cannot unmarshal string into Go value of type int](https://codesahara.com/blog/golang-cannot-unmarshal-string-into-go-value-of-type-int/)
[^2]: [Cannot unmarshal string into Go value of type int](https://github.com/Keithwachira/golang-decode-string-json-to-int)
