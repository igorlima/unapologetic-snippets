---
layout: default
title: Tricks and Tips
nav_order: 50
parent: Containerization
grand_parent: Programming Languages
permalink: /docs/languages/containerization/trick-and-tips
---

__[go to *containerization* page]({% link docs/languages/containerization/index.md %})__

# Docker built-in way to download a file

__The Old Way: Installing curl__
```Dockerfile
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache curl
RUN curl -sS https://example.com/somefile.txt -o /app/somefile.txt

EXPOSE 8080
```

__The Better Way: Using Docker's ADD Instruction__
```Dockerfile
FROM alpine:latest
WORKDIR /app
ADD https://example.com/anotherfile.json /app/anotherfile.json

EXPOSE 8080
```

- __When to Use Each Approach__ [^1]
  - Use `ADD` when:
    - You're downloading a single file from a URL
    - The file doesn't require authentication
    - You want to keep your image minimal
  - Stick with `curl` or `wget` when:
    - You need more control over the download (headers, authentication, retries)
    - You're fetching multiple files in a complex workflow
    - You need to process or validate the downloaded content before using it


----

[^1]: [A Docker Trick I Wish I Knew Sooner](https://dev.to/joybtw/a-docker-trick-i-wish-i-knew-sooner-23of)
