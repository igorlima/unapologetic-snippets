---
layout: default
title: Docker - Samples - AI
nav_exclude: true
parent: Containerization
grand_parent: Programming Languages
permalink: /docs/languages/containerization/docker-samples-ai
---

__[back]({% link docs/languages/containerization/docker.sample.md %})__

AI copilot for knowledge base - to easily chat and search things  on top of personal notes and documents.

Merge AI with the brain.

<details markdown="block">
  <summary>
    khoj <sub><i>a copilot to search and chat (using RAG) with your knowledge base (pdf, markdown, org)</i></sub>
  </summary>

Khoj is the open-source, AI copilot for your knowledge base. Easily build on
top of your personal notes and documents.
- [khoj home page](https://khoj.dev/)
  - [online app](https://app.khoj.dev/)
- About
  - An AI second brain. A copilot to search and chat (using RAG) with your knowledge base (pdf, markdown, org). Use powerful, online (e.g gpt4) or private, offline (e.g mistral) LLMs. Self-host locally or have it always accessible on the cloud. Access from Obsidian, Emacs, Desktop app, Web or Whatsapp

`docker-compose.yml`
```bash
# based on:
# https://github.com/khoj-ai/khoj/blob/1da453306edda3c9590779c2f8e2480eed66ca75/docker-compose.yml
#
# https://docs.khoj.dev/
# Self-Host
# $> docker-compose up
# $> docker-compose down
#  Khoj should be running at http://localhost:42110.
#  You can see the web UI in your browser.
#
# Uninstall Khoj Server
# $> docker-compose down --volumes
version: "3.9"
services:
  database:
    container_name: khoj-db
    hostname: khoj-db
    image: ankane/pgvector
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: postgres
    volumes:
      - khoj_db:/var/lib/postgresql/data/
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 30s
      timeout: 10s
      retries: 5
  server:
    container_name: khoj-server
    hostname: khoj-server
    depends_on:
      database:
        condition: service_healthy
    # Use the following line to use the latest version of khoj. Otherwise, it will build from source.
    image: ghcr.io/khoj-ai/khoj:latest
    # Uncomment the following line to build from source. This will take a few minutes. Comment the next two lines out if you want to use the offiicial image.
    # build:
      # context: .
    ports:
      # If changing the local port (left hand side), no other changes required.
      # If changing the remote port (right hand side),
      #   change the port in the args in the build section,
      #   as well as the port in the command section to match
      - "42110:42110"
    working_dir: /app
    volumes:
      - khoj_config:/root/.khoj/
      - khoj_models:/root/.cache/torch/sentence_transformers
    # Use 0.0.0.0 to explicitly set the host ip for the service on the container. https://pythonspeed.com/articles/docker-connection-refused/
    environment:
      - POSTGRES_DB=postgres
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_HOST=database
      - POSTGRES_PORT=5432
      - KHOJ_DJANGO_SECRET_KEY=secret
      - KHOJ_DEBUG=False
      - KHOJ_ADMIN_EMAIL=username@example.com
      - KHOJ_ADMIN_PASSWORD=password
    command: --host="0.0.0.0" --port=42110 -vv --anonymous-mode

volumes:
  khoj_config:
  khoj_db:
  khoj_models:
```

----
<br/>
<!-- khoj -->
</details>

---

[^1]: [...](...)
