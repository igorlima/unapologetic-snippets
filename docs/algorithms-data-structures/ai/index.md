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

## Neurons: the building blocks of AI

What a neuron is and how researchers and engineers piece these neurons together to build complex systems capable of incredibly challenging tasks. [^2]

You can think of these neurons as the mathematical functions that are doing this hard investigative work, getting into the weeds of the data and figuring out what’s going on.

Researchers can combine neurons in two ways. First, neurons can be lined up in a sequence, so the output of one becomes the input of the next. Neurons can also be stacked in layers, where they don’t interact directly but take the same input values. Some special neurons can even accept their own output and use it to update their internal function, in a kind of simulated memory.

## Predictive AI vs. Generative AI

Generative AI - The LLMs (like ChatGPT) that spit out text, images, and deepfakes. Predictive AI–The machine learning models that forecast trends and detect patterns.

Predictive AI figures out what's about to happen. Generative AI automates a response.

## MCP

MCP stands for Model Context Protocol. It's a relatively new open standard for connecting AI models to the world outside–the datasets, tools and systems they need to access, which helps them produce better, more contextually relevant responses.

MCP provides a standardized 'connector'–a single protocol that integrates multiple tools and services, which is a significant advancement on the fragmented, complex system of separate connectors for different data sources that developers have had to use until now.

MCP is open and model-agnostic, which means any AI model can use it and any developer can create an MCP integration without permission, making it an attractive prospect for the AI community.

The protocol defines how clients and servers communicate, keeping everything in sync.

MCP (Model Context Protocol) is a groundbreaking open standard for connecting AI models to external data and tools.

------ ------

[^1]: [What are: Embeddings? Vector Databases? Vector Search? k-NN? ANN?](https://kozyrkov.medium.com/what-are-embeddings-vector-databases-vector-search-k-nn-ann-9eb35f715c94)
[^2]: [The beginner’s guide to AI model architectures](https://read.technically.dev/p/all-about-ai-model-architectures)
