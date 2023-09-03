---
layout: default
title: gRPC Server and Client in Golang
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other-grpc-sample-b
---

__[back]({% link docs/languages/golang/other.grpc.md %})__

## gRPC Server and Client in Golang

- [what is gRPC?]({% link docs/algorithms-data-structures/index.md %}#grpc)


<details markdown="block">
  <summary>
    code snippet
  </summary>

A sample from a Medium post [^1].

```bash
# init
mkdir remote_debugging && cd remote_debugging
go mod init github.com/igorlima/remote_debugging
go mod tidy
```

<details markdown="block">
  <summary>
    <sup>go.mod</sup>
  </summary>
```go.mod
module github.com/igorlima/remote_debugging

go 1.20

require (
  github.com/golang/protobuf v1.3.5
  google.golang.org/grpc v1.29.1
)
```
----
<br/>
<!-- go.mod -->
</details>

```bash
# main.go
mkdir src && cd src && touch main.go
vi main.go
```

<details markdown="block">
  <summary>
    <sup>main.go</sup>
  </summary>
```golang
package main

import (
  "log"
  "net"
  "strconv"

  "google.golang.org/grpc"
)

func main() {
  port := 4867
  // Establishing a tcp connection.
  listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
  if err != nil {
    log.Fatalf("Unable to establish gRPC connection on: %d Error: %s", port, err.Error())
  }
  server := grpc.NewServer()
  log.Println("Started listening on port", port)
  if err := server.Serve(listen); err != nil {
    log.Fatalf("Serve: Error at gRPC Serve Error: %s", err.Error())
  }
}
```
----
<br/>
<!-- main.go -->
</details>

Unlike REST which communicates via plain text using HTTP 1.0 protocol, GRPC on the hand communicates via Protocol Buffers(Rawly it a set of fields in a struct that maintains the order of fields) using HTTP 2.0 protocol. For that, we need each message request and response struct and an RPC method that gets invoked by the client. __*Let’s create the request, response, and method interface.*__

```bash
# RPC method
mkdir pkg && mkdir pkg/user_registration && mkdir proto && cd proto && touch user_registration.proto
vi user_registration.proto
```

<details markdown="block">
  <summary>
    <sup>proto file</sup>
  </summary>
```proto
syntax = "proto3";
package main;
option go_package = "./user_registration";
// The request message containing the user's name.
message UserRegistrationRequest {
    string user_name = 1;
    string password = 2;
    string gender = 3;
    int64 age = 4;
    // Can keep going.
}
// The response message containing the greetings
message UserRegistrationResponse {
    bool success = 1;
}
service RemoteDebugPost {
    // Sends a greeting
    rpc Register (UserRegistrationRequest) returns (UserRegistrationResponse) {}
}
```

The above file specifies what had to be expected in the user registration request, what will be sent in response, and an RPC method named __Register__ which accepts the request and returns the response.

----
<br/>
<!-- proto file -->
</details>

Now that we created the basic proto file, we need to make sure that our Go Compiler understands this format. For that, we need to run a command that generates a `.pb.go` file. The below command generates the `user_registration.pb.go` file in `/pkg/user_registration` directory relative to root directory.

<details markdown="block">
  <summary>
    <sup>in case: <code>protoc: command not found</code></sup>
  </summary>
```bash
# in case: protoc: command not found
# to install the protocol buffer compiler
# https://grpc.io/docs/protoc-installation/
# https://grpc.io/docs/languages/go/quickstart/
brew install protobuf
protoc --version  # Ensure compiler version is 3+

# something to give a try
go mod tidy
# something else to give a go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go mod download google.golang.org/grpc
go mod download github.com/golang/protobuf
```
----
<br/>
<!-- in case: protoc: command not found -->
</details>

<details markdown="block">
  <summary>
    <sup>in case: <code>undefined: grpc.SupportPackageIsVersion7</code> or <code>undefined: grpc.ServiceRegistrar</code></sup>
  </summary>
```bash
# in case: `undefined: grpc.ServiceRegistrar` or `undefined: grpc.ServiceRegistrar`
# https://stackoverflow.com/questions/63662787/undefined-grpc-supportpackageisversion7-grpc-serviceregistrar
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
```
----
<br/>
<!-- in case: `undefined: grpc.ServiceRegistrar` or `undefined: grpc.ServiceRegistrar` -->
</details>

```bash
# generate Go code from the protobuf definition
# protoc --go_out=plugins=grpc:../pkg/user_registration user_registration.proto
# protoc --go_out=../pkg/user_registration --go_opt=Muser_registration.proto=github.com/igorlima/remote_debugging user_registration.proto

# run it inside the proto folder
# protoc --go_out=../pkg/user_registration --go_opt=Muser_registration.proto=. user_registration.proto
# protoc --go_out=../pkg --go_opt=Muser_registration.proto=. user_registration.proto
protoc --go_out=../pkg --go-grpc_out=../pkg --go_opt=Muser_registration.proto=. user_registration.proto
# when running inside the src folder
# protoc --go_out=./pkg/user_registration --go_opt=Muser_registration.proto=./proto proto/user_registration.proto
# protoc --go_out=./pkg --go_opt=Muser_registration.proto=./proto proto/user_registration.proto
protoc --go_out=./pkg --go-grpc_out=./pkg --go_opt=Muser_registration.proto=./proto proto/user_registration.proto
# Note: The M before user_registration.proto
# https://stackoverflow.com/questions/72107940/solving-protoc-gen-go-unable-to-determine-go-import-path-problem-by-adding-a-m
```

```bash
# To check the file `user_registration.pb.go` file in `/pkg/user_registration` directory relative to root directory.
ls ../pkg/user_registration/
ls src/pkg/user_registration/
```

Once all these steps are done, the repo should look like below.

```
 |
 |-go.mod
 |-go.sum
 |-src
 | |-proto
 | | |-user_registration.proto
 | |-main.go
 | |-pkg
 | | |-user_registration
 | | | |-user_registration_grpc.pb.go
 | | | |-user_registration.pb.go
```



<details markdown="block">
  <summary>
    <sup>server.go</sup>
  </summary>
```go
package main

import (
  "context"
  "fmt"
  "log"
  "net"
  "strconv"

  user_registration "github.com/igorlima/remote_debugging/src/pkg/user_registration"
  "google.golang.org/grpc"
)

// server is responsible to implement all the RPC Methods defined.
type rpc struct {
  user_registration.UnimplementedRemoteDebugPostServer
}

func (s *rpc) Register(ctx context.Context, request *user_registration.UserRegistrationRequest) (
  response *user_registration.UserRegistrationResponse, err error) {

  fmt.Println("%v\n", request)
  return &user_registration.UserRegistrationResponse{
    Success: true,
  }, nil
}

func registerRPCMethods(server *grpc.Server) {
  rpcImplementer := new(rpc)
  user_registration.RegisterRemoteDebugPostServer(server, rpcImplementer)
}

func main() {
  port := 4867
  // Establishing a tcp connection.
  listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
  if err != nil {
    log.Fatalf("Unable to establish gRPC connection on: %d Error: %s", port, err.Error())
  }
  server := grpc.NewServer()
  registerRPCMethods(server)
  log.Println("Started listening on port", port)
  if err := server.Serve(listen); err != nil {
    log.Fatalf("Serve: Error at gRPC Serve Error: %s", err.Error())
  }
}
```
------
<!-- server.go -->
</details>

<details markdown="block">
  <summary>
    <sup>client.go</sup>
  </summary>
```golang
package main

import (
  "context"
  "fmt"

  user_registration "github.com/igorlima/remote_debugging/src/pkg/user_registration"
  "google.golang.org/grpc"
)

func main() {
  conn, err := grpc.Dial("localhost:4867", grpc.WithInsecure())
  if err != nil {
    panic("Not able to make client connection")
  }
  defer conn.Close()
  client := user_registration.NewRemoteDebugPostClient(conn)

  reqContext := context.Background()
  request := &user_registration.UserRegistrationRequest{
    UserName: "mourya",
    Password: "secretpls",
    Gender:   "M",
    Age:      24,
  }
  response, err := client.Register(reqContext, request)
  fmt.Println(response, err)
}
```
------
<!-- client.go -->
</details>

After creating the file `server.go` and `client.go`, the repo should look like below.

```
 |
 |-/go.mod
 |-/go.sum
 |-/src
 | |-server.go
 | |-proto
 | | |-user_registration.proto
 | |-client.go
 | |-main.go
 | |-pkg
 | | |-user_registration
 | | | |-user_registration_grpc.pb.go
 | | | |-user_registration.pb.go
```

```bash
# run the gRPC server
go run src/server.go
```

```bash
# in a separate terminal, run the gRPC client:
go run src/client.go
```


</details>


----

[^1]: [Remote Debugging a Dockerized Golang-GRPC Server using Goland Editor, Delve](https://medium.com/swlh/remote-debugging-golang-grpc-service-using-goland-delve-and-docker-4df71be21201)
