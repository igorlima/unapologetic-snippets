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

------ ------

[^1]: [How to conduct an A/B test experiment in Go using feature flags](https://medium.com/@chavezharris/how-to-conduct-an-a-b-test-experiment-in-go-using-feature-flags-660edef536f6)
[^2]: [Golang: How are Data Races, Context Switching, and Multi-Core Computing interlinked](https://mourya-g9.medium.com/golang-how-are-data-races-context-switching-and-multi-core-computing-related-cba8d17b542a)
[^3]: [Introduction to Microservices, Go-Kit, Grpc. Golang](https://medium.com/@edwinsiby/introduction-to-microservices-go-kit-grpc-golang-e187853d6c45)
