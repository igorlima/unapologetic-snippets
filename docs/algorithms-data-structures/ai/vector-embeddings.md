---
layout: default
title: Vector Embeddings
nav_order: 1
parent: Artificial Intelligence
grand_parent: Algorithms and Data Structures
permalink: /docs/algorithms-and-data-structures/ai/vector-embeddings
---

- [go back]({% link docs/algorithms-data-structures/ai/index.md %})

# What are Vector Embeddings? [^1]

Embeddings are numerical machine learning representations of the semantic of the input data.

They capture the meaning of complex, high-dimensional data, like text, images, or audio, into vectors. Enabling algorithms to process and analyze the data more efficiently.

__How do embeddings work?__

The nuances of natural language or the hidden meaning in large datasets of images, sounds, or user interactions are hard to fit into a table. Traditional relational databases can't efficiently query most types of data being currently used and produced, making the retrieval of this information very limited.

At their core, vector embeddings are about semantics. They take the idea that "a word is known by the company it keeps" and apply it on a grand scale.

----------

SAVE IMAGE

----------

This capability is crucial for creating search systems, recommendation engines, retrieval augmented generation (RAG) and any application that benefits from a deep understanding of content.

Embeddings are created through neural networks. They capture complex relationships and semantics into dense vectors which are more suitable for machine learning and data processing applications. They can then project these vectors into a proper high-dimensional space, specifically, a Vector Database.

----------

SAVE IMAGE

----------

The meaning of a data point is implicitly defined by its position on the vector space. After the vectors are stored, we can use their spatial properties to perform nearest neighbor searches.

These searches retrieve semantically similar items based on how close they are in this space.

The quality of the vector representations drives the performance. The embedding model that works best for you depends on your use case.

__Creating Vector Embeddings__

Embeddings translate the complexities of human language to a format that computers can understand. It uses neural networks to assign numerical values to the input data, in a way that similar data has similar values.

----------

SAVE IMAGE

----------

We can introduce more dimensions to capture more nuances.

For example, a third dimension could represent formality of the word, a fourth could indicate its emotional connotation (positive, neutral, negative), and so on.

----------

SAVE IMAGE

----------

__Integration with Embedding APIs__

Selecting the right embedding model for your use case is crucial to your application performance.

If you’re looking for NLP and rapid prototyping, including language translation, question-answering, and text generation, OpenAI is a great choice. Gemini is ideal for image search, duplicate detection, and clustering tasks.

Fastembed, which we’ll use on the example below, is designed for efficiency and speed, great for applications needing lowlatency responses, such as autocomplete and instant content recommendations.


----

[^1]: [What are Vector Embeddings?](https://dev.to/qdrant/what-are-vector-embeddings-24pd)
