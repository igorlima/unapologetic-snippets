---
layout: default
title: Compiling Go Application - Go Binary
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/commands-build-binary
---

# Compiling Go Application - Go Binary
__[back]({% link docs/languages/golang/commands.md %}#go-build)__
<br/>

To make the binary file even more suitable to run in a container, you can pass additional build options. For example, you’ll enable a statically linked binary by setting `CGO_ENABLED=0`, and you can pass additional linker options using the flag `-ldflags`. To reduce the binary size, use the options `-ldflags=”-s -w”` to strip the binary of debug symbols. Some of the build options that you’ll use[^1]:
- `CGO_ENABLED=0`: Enables statically linked binaries to make the application more portable. It allows you to use the binary with source images that don’t support shared libraries when building your container image.
- `GOOS=linux`: Since containers run Linux, set this option to enable repeatable builds even when building the application on a different platform.
- `-ldflags=”-s -w”`: The parameter `-ldflags` allows you to specify additional linker options that go build uses at the link stage of the build process. In this case, the option `-s -w` strips the binary of debugging symbols, decreasing its size. Without these symbols, it’s harder to debug the application, but this is usually not a major concern when running in a container. To see all linker options you can use, run go tool link.
- `-tags=containers`: This is specific to your own application.

Build your binary using these options:
```sh
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -tags=containers
```

Once you have the binary, you’ll create a container image by using a Dockerfile.

----

[^1]: [Compiling Your Go Application for Containers](https://medium.com/pragmatic-programmers/compiling-your-go-application-for-co-ntainers-b513190471aa)
