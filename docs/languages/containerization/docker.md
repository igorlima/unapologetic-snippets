---
layout: default
title: Docker
nav_order: 10
parent: Containerization
grand_parent: Programming Languages
permalink: /docs/languages/containerization/docker
---

__[go to *containerization* page]({% link docs/languages/containerization/index.md %})__

__How does Docker work?__ [^1]
The Docker engine runs on top of the host operating system. Docker Engine includes a server process (`dockerd`) that manages docker containers on the host system. Docker containers are designed to isolate applications and their dependencies, ensuring that they can run consistently across different environments

To work with Docker, we need to understand three concepts, __Dockerfile__, __Docker image__, __Docker container__:
- What is a Docker file (Dockerfile)? _Docker file is a blueprint to build a docker image._
- What is a Docker image? _Docker image is a template for running docker containers._
  - Docker images contain all the dependencies needed to execute code inside a container.
- What is a Docker container? Container is just a running process.

__Docker volumes.__ A volume is a location in your local filesystem, managed by Docker. Volume doesn’t increase the size of the containers using it, and the volume’s contents exist outside the lifecycle of a given container.

[Docker Hub](https://hub.docker.com/) repository. Create an account if you haven’t already. My is [`igorlima`](https://hub.docker.com/u/igorlima).

To push the image to the docker hub.
```sh
# before you can push the image, run `docker login`
docker login
# docker push <tagged-image>
docker push igorlima/hello-world:0.1
```


----

[^1]: [Devops zero to hero #3 — Everything you need to know about Dockers](https://medium.com/illumination/devops-zero-to-hero-3-everything-you-need-to-know-about-dockers-7ff321b38e6b)
