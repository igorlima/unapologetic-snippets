---
layout: minimal
title: Algorithms and Data Structures
nav_order: 98
---

[Return to main website]({{site.baseurl}}/).

# Algorithms and Data Structures

## A/B test experiment

__What is an A/B Test?__ [^1]
An A/B test is a type of testing in which two variations of an app are released for testing to an isolated user segment before being deployed to the entire user base to collect metrics and feedback. One variation is designated variation A, and the other is designated variation B. Variation A usually represents the current version of the app and is considered a control or benchmark variation, while variation B represents the version of the app with the new feature to be tested.
__The role feature flags play.__
Feature flags are virtual switches that you can integrate into your app to control what features users can and cannot see.

## Data Races

Data Races are the toughest and the most time-consuming thing to debug. [^2]

![data race ilustration](https://github.com/igorlima/unapologetic-snippets/assets/1886786/a3866f1d-7506-4d2d-9ada-fc0229483187)
<sup style="display:block;">image from [Medium](https://mourya-g9.medium.com/golang-how-are-data-races-context-switching-and-multi-core-computing-related-cba8d17b542a) [^2]</sup>

## gRPC

gRPC is a high-performance, open-source framework for remote procedure calls (RPC). It uses Protocol Buffers (protobuf) as the interface definition language and supports various programming languages, including Go. gRPC enables communication between microservices using a strongly-typed contract, allowing for efficient and reliable service-to-service communication. [^3]

- [a sample in Go]({% link docs/languages/golang/other.grpc.md %})

## Artificial Intelligence (AI)

- [a knowledge base powered by ChatGPT in Python]({% link docs/languages/python/index.md %})
  - a Python implementation that, instead of bringing your documents to ChatGPT; instead, brings ChatGPT to your documents using apps like Google Drive, Notion, and Dropbox as your knowledge base.

## Coroutines, Threads, and Generators

At first glance, coroutines, threads, and generators look alike. All three provide concurrency in one form or another, but they differ in important ways. [^4]

- Coroutines provide concurrency without parallelism: when one coroutine is running, the one that resumed it or yielded to it is not.
Because coroutines run one at a time and only switch at specific points in the program, the coroutines can share data among themselves without races. The explicit switches serve as synchronization points, creating happens-before edges.
Because scheduling is explicit (without any preemption) and done entirely without the operating system, a coroutine switch takes at most around ten nanoseconds, usually even less. Startup and teardown is also much cheaper than threads.
- Threads provide more power than coroutines, but with more cost. The additional power is parallelism, and the cost is the overhead of scheduling, including more expensive context switches and the need to add preemption in some form. Typically the operating system provides threads, and a thread switch takes a few microseconds.
For this taxonomy, Go's goroutines are cheap threads: a goroutine switch is closer to a few hundred nanoseconds, because the Go runtime takes on some of the scheduling work, but goroutines still provide the full parallelism and preemption of threads. (Java's new lightweight threads are basically the same as goroutines.)
- Generators provide less power than coroutines, because only the top-most frame in the coroutine is allowed to yield. That frame is moved back and forth between an object and the call stack to suspend and resume it.

Coroutines are a useful building block for writing programs that want concurrency for program structuring but not for parallelism.

------ ------

[^1]: [How to conduct an A/B test experiment in Go using feature flags](https://medium.com/@chavezharris/how-to-conduct-an-a-b-test-experiment-in-go-using-feature-flags-660edef536f6)
[^2]: [Golang: How are Data Races, Context Switching, and Multi-Core Computing interlinked](https://mourya-g9.medium.com/golang-how-are-data-races-context-switching-and-multi-core-computing-related-cba8d17b542a)
[^3]: [Introduction to Microservices, Go-Kit, Grpc. Golang](https://medium.com/@edwinsiby/introduction-to-microservices-go-kit-grpc-golang-e187853d6c45)
[^4]: [Coroutines for Go](https://research.swtch.com/coro)
