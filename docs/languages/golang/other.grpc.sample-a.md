---
layout: default
title: gRPC Server and Client in Golang
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other-grpc-sample-a
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
go mod init github.com/igorlima/grpc-sample
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

<details markdown="block">
  <summary>
    <sup>go.mod</sup>
  </summary>
```go.mod
module github.com/igorlima/grpc-sample

go 1.18

require (
  github.com/golang/protobuf v1.5.3 // indirect
  golang.org/x/net v0.9.0 // indirect
  golang.org/x/sys v0.7.0 // indirect
  golang.org/x/text v0.9.0 // indirect
  google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
  google.golang.org/grpc v1.56.2 // indirect
  google.golang.org/protobuf v1.31.0 // indirect
)
```
<br/>
</details>

```proto
syntax = "proto3";

package main;
option go_package = "./grpcsample";

service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
```

```bash
# generate Go code from the protobuf definition
protoc --go_out=. --go-grpc_out=. hello.proto
```

<details markdown="block">
  <summary>
    <sup>server.go</sup>
  </summary>
```go
package main

import (
  "context"
  "log"
  "net"

  pb "github.com/igorlima/grpc-sample/grpcsample"
  "google.golang.org/grpc"
)

type server struct {
  pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
  return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
  lis, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterHelloServiceServer(s, &server{})
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
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
  "log"

  pb "github.com/igorlima/grpc-sample/grpcsample"
  "google.golang.org/grpc"
)

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewHelloServiceClient(conn)

  name := "John"
  res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Response: %s", res.Message)
}
```
------
<!-- client.go -->
</details>

```bash
# run the gRPC server
go run server.go
```

```bash
# in a separate terminal, run the gRPC client:
go run client.go
```


</details>


----

[^1]: [Introduction to Microservices, Go-Kit, Grpc. Golang](https://medium.com/@edwinsiby/introduction-to-microservices-go-kit-grpc-golang-e187853d6c45)
