---
layout: default
title: WebAssembly with Go
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other-webassembly
---

# WebAssembly with Go
__[back]({% link docs/languages/golang/other.md %})__
<br/>

WebAssembly (often abbreviated as __Wasm__) is a binary instruction format for a stack-based virtual machine. [^1]

WebAssembly code is compiled into a binary format executed directly by the browser’s virtual machine rather than interpreted like traditional JavaScript code.

It runs in a sandboxed environment, preventing it from accessing sensitive user data or interfering with other web page parts.

WebAssembly is also designed to be portable, meaning that it can be used across a wide range of platforms and devices. This makes it an ideal tool for building web applications that need to work seamlessly across different devices and operating systems.

## Write WebAssembly With Go

WebAssembly is a binary format that allows for code deployment on the web. It is designed to be fast, efficient, and secure and can be used with various programming languages, including Go.

To use WebAssembly with Go, you must compile your Go code to WebAssembly. After compiling the code, it can be used just like any other JavaScript code in the web application.

```sh
go mod init wasmcounter
```

```golang
// main.go
package main

import "fmt"

func main() {
  fmt.Println("Hello from Go WebAssembly!")
}
```

Compile our code into wasm format. Use two flags `GOOS=js` and `GOARCH=wasm` for the Go compiler.

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go

# copy the `wasm_exec.js` with this command:
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

```html
<!-- index.html -->
<html>
<head>
  <meta charset="utf-8" />
  <script src="wasm_exec.js"></script>
  <script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
      go.run(result.instance);
    });
  </script>
</head>
<body>
  <h1>Hello WebAssembly</h1>
</body>
</html>
```

```sh
# to serve the page
python3 -m http.server 6161
# open the page
open http://localhost:6161
```

![image](https://user-images.githubusercontent.com/1886786/232261981-31b1d546-d518-485f-b85c-53a9b9f9fd5d.png)


## A Counter App

```golang
// main.go
package main

import (
  "fmt"
  "syscall/js"
)

func main() {
  c := make(chan bool)
  fmt.Println("Hello from Go WebAssembly!")
  js.Global().Set("increment", js.FuncOf(increment))
  js.Global().Set("decrement", js.FuncOf(decrement))
  <-c
}

var count int

func increment(this js.Value, args []js.Value) interface{} {
  count++
  return count
}

func decrement(this js.Value, args []js.Value) interface{} {
  count--
  return count
}
```

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
# copy the `wasm_exec.js` with this command:
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

```html
<!-- index.html -->
<html>
<head>
  <meta charset="utf-8" />
  <script src="wasm_exec.js"></script>
  <script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
      go.run(result.instance);
    });
  </script>
</head>
<body>
  <h1>WebAssembly in Go</h1>
  <button id="increment" onclick="incrementCounter()">Increment</button>
  <span id="count">0</span>
  <button id="decrement" onclick="decrementCounter()">Decrement</button>
</body>
<script>
  const incrementCounter = () => {
    const count = increment()
    setCounterValue(count)
  }
  const decrementCounter = () => {
    const count = decrement()
    setCounterValue(count)
  }
  const setCounterValue = (value) => {
    const countEl = document.getElementById("count")
    countEl.textContent = value
  }
</script>
</html>
```

```sh
# to serve the page
python3 -m http.server 6161
# open the page
open http://localhost:6161
```

![image](https://user-images.githubusercontent.com/1886786/232261995-ee8d0815-776e-45a4-a972-868affb8d44e.png)


----

[^1]: [WebAssembly With Go, Create A Counter App](https://betterprogramming.pub/webassembly-with-go-create-a-counter-app-2bbe69d91845)
