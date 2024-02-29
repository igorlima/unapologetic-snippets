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

![image](https://github.com/ilima-kindle/unapologetic-snippets/assets/133927230/a5548a17-44f4-4d23-a96a-7cbed15f85f1)


This capability is crucial for creating search systems, recommendation engines, retrieval augmented generation (RAG) and any application that benefits from a deep understanding of content.

Embeddings are created through neural networks. They capture complex relationships and semantics into dense vectors which are more suitable for machine learning and data processing applications. They can then project these vectors into a proper high-dimensional space, specifically, a Vector Database.

![image](https://github.com/ilima-kindle/unapologetic-snippets/assets/133927230/6cf8bad2-5ce9-47b8-89ba-9aa5f5e33ac0)


The meaning of a data point is implicitly defined by its position on the vector space. After the vectors are stored, we can use their spatial properties to perform nearest neighbor searches.

These searches retrieve semantically similar items based on how close they are in this space.

The quality of the vector representations drives the performance. The embedding model that works best for you depends on your use case.

__Creating Vector Embeddings__

Embeddings translate the complexities of human language to a format that computers can understand. It uses neural networks to assign numerical values to the input data, in a way that similar data has similar values.

![image](https://github.com/ilima-kindle/unapologetic-snippets/assets/133927230/ec4b8f6b-6038-47b9-86c2-0dfead01faa7)

We can introduce more dimensions to capture more nuances.

For example, a third dimension could represent formality of the word, a fourth could indicate its emotional connotation (positive, neutral, negative), and so on.

![image](https://github.com/ilima-kindle/unapologetic-snippets/assets/133927230/4d693811-d2de-4e8d-85af-0df4d3cddb60)

![image](https://github.com/ilima-kindle/unapologetic-snippets/assets/133927230/f3e65d4e-d8e7-4d5b-afaf-11ea4296cfea)

__Integration with Embedding APIs__

Selecting the right embedding model for your use case is crucial to your application performance.

If you’re looking for NLP and rapid prototyping, including language translation, question-answering, and text generation, OpenAI is a great choice. Gemini is ideal for image search, duplicate detection, and clustering tasks.

[Fastembed](https://github.com/qdrant/fastembed), which we’ll use on the example below, is designed for efficiency and speed, great for applications needing lowlatency responses, such as autocomplete and instant content recommendations.


----

[^1]: [What are Vector Embeddings?](https://dev.to/qdrant/what-are-vector-embeddings-24pd)
