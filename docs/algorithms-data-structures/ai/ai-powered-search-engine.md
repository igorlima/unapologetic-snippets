---
layout: default
title: AI-Powered Search Engine
nav_order: 1
parent: Artificial Intelligence
grand_parent: Algorithms and Data Structures
permalink: /docs/algorithms-and-data-structures/ai/search-engine
---

- [go back]({% link docs/algorithms-data-structures/ai/index.md %})

- LlamaIndex Use cases: Good for tasks that require text search and retrieval, like information retrieval or content discovery Features: Excels in data indexing and language model enhancement Connectors: Provides connectors to access data from databases, external APIs, or other datasets
  - Ease of Search: Large chunks of data are harder to search over. Splitting data into smaller chunks therefore helps in better indexation.
  - Context Window Size: LLMs allow only a finite number of tokens in prompts and completions. The context therefore cannot be larger than what the context window permits.
  - When building RAG powered LLM apps, there is no right answer to "Which embeddings model to use?". However, you may notice particular embeddings working better for specific use cases (like summarization, text generations, classification etc.)


<details markdown="block">
  <summary> RAG in a nutshell </summary>
Retrieval Augmented Generation (RAG) [^2] — a software pattern that relies on the power of meaning-based search by turning raw data into vectors

Basics of RAG - RAG involves searching across a vast corpus of private data and retrieving results that are most similar to the query asked by the end user so that it can be passed on to the LLM as context.

A typical RAG architecture involves three steps:
- Step 1 — Create embeddings or vectors from a vast corpus of data.
- Step 2 — Store the vectors in a vector store as an index.
- Step 3 — Search through the vectors by comparing the query with the vector data and sending the retrieved content to the LLM.

Three categories of vector stores emerge — vector libraries, vector-only databases and enterprise databases that also added vector capabilities.

<details markdown="block">
  <summary><i> further more... </i></summary>

RAG is simple in theory (just add data to the context window!) but complex in practice. [^1]

At their core, LLMs are simple: you send them a prompt and you get back a response

Specifically, RAG seeks to give LLMs access to large quantities of additional knowledge in a way that circumvents the prohibitively expensive process of training new models

Embedding models, they are very big lists of numbers. The vectors that embedding models create are usually 768 or 1536 numbers (dimensions), but vectors of other sizes also exist.

Evaluating the quality of your RAG retrieval system, it is well worth your time to understand how the data is sourced and ingested before it ever hits the fancy AI parts of your RAG pipeline.

After data is ingested but before it can be run through an embedding model, it has to be split up into discrete pieces. So, how do you decide how to split the data? This is called your chunking strategy.

How large or how small of chunks are optimal? Should chunks overlap? Are there smarter ways to chunk than just dividing by page, paragraph, or fixed length? How should data in non-standard formats (code, JSON, etc.) be chunked?

These are the questions that a chunking strategy tries to answer, and there is no perfect solution. Different strategies have different trade-offs. Some are simple and fast to implement, giving passable results

Here are five categories of chunking: fixed-size, recursive, document-based, semantic, and agentic (using AI to chunk, nifty!)

The [Pinecone blog](https://www.pinecone.io/learn/chunking-strategies/) has more detail on many of these strategies.

There are many models that can be used to generate embeddings, and different models will perform better or worse in different situations. Some models come pretrained for general use, and some are fine-tuned for specific domains (i.e. medical records)

Some embedding models can only be accessed via an API (for example, OpenAI embeddings endpoint) while others are fully open sourced and can be downloaded and run locally or hosted in a cloud provider.

The TLDR of reranking: embedding models are optimized for speed, as they need to be run against a large number of documents. Other models called reranking models (or cross-encoders) are slower but optimized for accuracy. So a fast-inaccurate embedding model is used to generate the embeddings (that are saved in the VectorDB), and then the slow-accurate model is used to find the highest quality documents in this smaller set. The best matches from the slow-accurate search are prioritized in the context window.

Again, there’s a lot more to it than that, but this is the essence of reranking. The Pinecone blog has a great description of the process in more detail.

----
<!-- RAG further more -->
</details>

----
<!-- RAG in a nutshell -->
</details>

----

[^1]: [RAG for Quality Engineers](https://medium.com/slalom-build/rag-for-quality-engineers-c5f0828292b1)
[^2]: [The Ultimate Guide to Vector Database Landscape — 2024 and Beyond](https://medium.com/madhukarkumar/the-ultimate-guide-to-vector-databases-2024-and-beyond-16dfb15bef12)
