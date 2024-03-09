---
layout: default
title: Docker
nav_order: 10
parent: Containerization
grand_parent: Programming Languages
permalink: /docs/languages/containerization/docker
---

__[go to *containerization* page]({% link docs/languages/containerization/index.md %})__

- [Docker samples]({% link docs/languages/containerization/docker.sample.md %})

__How does Docker work?__ [^1]
The Docker engine runs on top of the host operating system. Docker Engine includes a server process (`dockerd`) that manages docker containers on the host system. Docker containers are designed to isolate applications and their dependencies, ensuring that they can run consistently across different environments

To work with Docker, we need to understand three concepts, __Dockerfile__, __Docker image__, __Docker container__:
- What is a Docker file (Dockerfile)? _Docker file is a blueprint to build a docker image._
- What is a Docker image? _Docker image is a template for running docker containers._
  - Docker images contain all the dependencies needed to execute code inside a container.
- What is a Docker container? Container is just a running process.

__Docker volumes.__ A volume is a location in your local filesystem, managed by Docker. Volume doesn’t increase the size of the containers using it, and the volume’s contents exist outside the lifecycle of a given container.

[Docker Hub](https://hub.docker.com/) repository. Create an account if you haven’t already. Mine is [`igorlima`](https://hub.docker.com/u/igorlima).

To push the image to the docker hub.
```sh
# before you can push the image, run `docker login`
docker login
# docker push <tagged-image>
docker push igorlima/hello-world:0.1
```

# Docker Commands

The below process is beneficial when you need to transfer images between different environments without relying on a registry. [^2]

```sh
docker save -o myimage.tar yourimage
docker load -i myimage.tar
```

The docker history command provides a timeline of an image, revealing the commands used to create each layer. [^2]

```sh
docker history yourimage
```

The `docker system df` command gives a quick overview of Docker’s disk usage.
It summarizes the space occupied by containers, images, and other components. This command is handy for figuring out what’s taking up disk space in your Docker environment. [^2]

```sh
docker system df
```

The `docker container prune` command does a beautiful job of removing all stopped containers. It’s a handy way to tidy up your system by getting rid of those stopped containers that are no longer needed. [^2]

```sh
docker container prune
```

Use the prune command to delete images but instead of container you use the image. [^2]

```sh
docker image prune
```

----

[^1]: [Devops zero to hero #3 — Everything you need to know about Dockers](https://medium.com/illumination/devops-zero-to-hero-3-everything-you-need-to-know-about-dockers-7ff321b38e6b)
[^2]: [I Never Heard of these Docker Commands](https://medium.com/@arton.demaku/unknown-docker-commands-c8641c686843)
