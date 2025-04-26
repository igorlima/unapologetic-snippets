---
layout: default
title: Artificial Intelligence
nav_order: 1
parent: Algorithms and Data Structures
has_children: true
has_toc: true
permalink: /docs/algorithms-and-data-structures/ai
---

# Artificial Intelligence

- [go back]({% link docs/algorithms-data-structures/index.md %}#artificial-intelligence-ai)

## How do you train an AI model?

In a sense, training a model really is like teaching a kid how to do something, like riding a bike. It's less about telling them how to do it, and more about giving them repetitions so they can figure out what's going on for themselves. With some well-timed guidance, of course.

A model is a decision-making machine. The way you train a model is by showing it many, many different situations and what the correct outcome is in those situations. The model uses some fancy math to learn the patterns in those situations and learns to apply them to new data.

## What are: Embeddings? Vector Databases? Vector Search? kNN? ANN? [^1]

**Embeddings: "Making Your Data AI-Friendly"**.
Embeddings convert raw data into numerical representations that capture relationships and distances within the data. This lower-dimensional representation enhances the efficiency of AI applications.

**Vector Databases: "A Hub for Embeddings and Raw Data"**.
Vector databases provide a centralized repository for embeddings and the corresponding raw data. They facilitate efficient storage, retrieval, and management of these resources.

**Vector Search: "Beyond Keyword-Based Queries"**.
Vector search enables similarity-based queries over vector representations of data. Unlike keyword search, vector search retrieves results based on distance metrics, providing more contextually relevant matches.

**Nearest Neighbors (k-NN): "Finding the Closest Matches"**.
k-NN is an algorithm that retrieves the k most similar data points to a given query embedding. It is commonly used in applications such as recommendation systems and anomaly detection.

**Approximate Nearest Neighbors (ANN): "A Faster Alternative to k-NN"**.
ANN algorithms provide approximate solutions to the nearest neighbor problem. They sacrifice accuracy in favor of significantly reduced computation time, making them suitable for large-scale data analysis and real-time applications.

------ ------

[^1]: [What are: Embeddings? Vector Databases? Vector Search? k-NN? ANN?](https://kozyrkov.medium.com/what-are-embeddings-vector-databases-vector-search-k-nn-ann-9eb35f715c94)
