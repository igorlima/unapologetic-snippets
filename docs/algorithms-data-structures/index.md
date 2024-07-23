---
layout: default
title: Algorithms and Data Structures
nav_order: 4
has_children: true
has_toc: false
permalink: /docs/algorithms-and-data-structures
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

- [a sample in Go]({% link docs/languages/golang/other.grpc.md %})
- [a sample building a gRPC server and client in Node.js](https://semaphoreci.medium.com/an-introduction-to-grpc-building-distributed-systems-with-efficiency-and-scalability-in-mind-c13024e4b0d3)
- [a implementation of Protocol Buffers in TypeScript](https://github.com/bufbuild/protobuf-es)

__What are Protocol Buffers?__
In a nutshell, Protocol Buffers have two main functions:
- They are a language for writing schemas for your data.
- They define a binary format for serializing your data.

These two independent traits functions work together to allow your project and everyone who interacts with it to define messages, fields, and service APIs in the exact same way.
Protocol Buffers also allow you to serialize this structured data.

__How does gRPC work?__[^9] gRPC utilizes Remote Procedure Call (RPC) technology to facilitate communication between disparate services using a protocol buffer. _The protocol buffer serves as a conduit for structured data exchange between systems_.

__gRPC API service methods__[^9] are functionalities that can be remotely called by clients in gRPC API are defined by service methods. These methods are specified in an _Interface Definition Language_ (IDL) file written in Protocol Buffer format (`.proto`). This IDL file is utilized to create the server and client code for gRPC.

__Advantages of gRPC__:
- gRPC offers strongly typed contracts that use Protocol Buffers to define the service contract. This ensures that the data being transmitted is consistent and validated, which reduces the chance of errors and bugs.
- gRPC supports bi-directional streaming, allowing both the client and server to send and receive data simultaneously, which enables real-time communication and reduces latency.

__Design Your API.__[^10] The `.proto` files for the service defines the source of truth for your API spec. Well commented ones can even be used to [generate documentation](https://github.com/pseudomuto/protoc-gen-doc).



## [Artificial Intelligence (AI)]({% link docs/algorithms-data-structures/ai/index.md %})

- [a knowledge base powered by ChatGPT in Python]({% link docs/languages/python/other.chatgpt-knowledge-base.md %})
  - a Python implementation that, instead of bringing your documents to ChatGPT; instead, brings ChatGPT to your documents using apps like Google Drive, Notion, and Dropbox as your knowledge base.

### Prompts

In AI writing, prompts are what the user says to the AI writing tool to create the context they want. Prompts are essentially commands, which specify the parameters of the assignment such as the topic, length, style, and which words to use. [^5]

### Text embeddings

The foundation technology for modern-day natural language processing techniques is _text empowering_. It's a step we need to convert input into the correct representation. A text embedding[^7], also called a _vector_, is a piece of text that's projected into a high-dimensional space.

### Similarity

Matching is about finding similarities, and there are many ways of doing it. One popular way of doing this whit text data is with _cosine similarity_. _Cosine similarity_[^7] is a measure of how similar 2 vectors are.

A _vector database_ is a database that can store vectors along with other information.

![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/db077170-411b-47cf-b9dc-84a82297ae4c)

### Retrieval Augmented Generation (RAG)

__What is a RAG?__
Retrieval Augmented Generation (RAG)[^8]. Simply put, RAGs help LLMs by giving them access to external data so that they can generate a response with additional context. This context can be anything from recent news, audio transcripts of a lecture, or in my case — fantasy football news.

__How do they work?__
You can think of RAG as an LLM with vector search attached. Here are the 4 key steps that take place:
1. Load a vector database with encoded documents.
1. Encode the query into a vector using a sentence transformer.
1. Based on the inputted query, retrieve relevant context from thevector database.
1. Leverage context along with the query to prompt the LLM.

Retrieval Augmented Generation (RAG) is a powerful NLP technique that combines large language models with selective access to knowledge - by providing the relevant pieces of the context from our documents. [^12]
- <details markdown="block"><summary><i>image ilustrating RAG</i></summary>
   
  ![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/ee7c67e7-9dc8-4f96-8802-233303d3bc1a)
  _<sup style="display:block;">image from Medium [^13]</sup>_
  </details>


## Functional Programming

Functional programming principles is about empowering the developer to work with immutable data, compose functions, and bring a declarative flair to the code.

- [Functional Paradigms in Go]({% link docs/languages/golang/other.functional-paradigms.md %})

## Coroutines, Threads, and Generators

At first glance, coroutines, threads, and generators look alike. All three provide concurrency in one form or another, but they differ in important ways. [^4]

- Coroutines provide concurrency without parallelism: when one coroutine is running, the one that resumed it or yielded to it is not.
Because coroutines run one at a time and only switch at specific points in the program, the coroutines can share data among themselves without races. The explicit switches serve as synchronization points, creating happens-before edges.
Because scheduling is explicit (without any preemption) and done entirely without the operating system, a coroutine switch takes at most around ten nanoseconds, usually even less. Startup and teardown is also much cheaper than threads.
- Threads provide more power than coroutines, but with more cost. The additional power is parallelism, and the cost is the overhead of scheduling, including more expensive context switches and the need to add preemption in some form. Typically the operating system provides threads, and a thread switch takes a few microseconds.
For this taxonomy, Go's goroutines are cheap threads: a goroutine switch is closer to a few hundred nanoseconds, because the Go runtime takes on some of the scheduling work, but goroutines still provide the full parallelism and preemption of threads. (Java's new lightweight threads are basically the same as goroutines.)
- Generators provide less power than coroutines, because only the top-most frame in the coroutine is allowed to yield. That frame is moved back and forth between an object and the call stack to suspend and resume it.

Coroutines are a useful building block for writing programs that want concurrency for program structuring but not for parallelism.

## Bitwise Operations

Bitwise operators are characters that represent actions (bitwise operations) to be performed on single bits. They operate at the binary level and perform operations on bit patterns that involve the manipulation of individual bits.

- [bitwise operations in Python]({% link docs/languages/python/other.bitwise-operations.md %})

## Observability

Observability is how businesses know what’s going on with their different systems and operations. [^15]

- Businesses are made up of **different software systems**, from apps and infrastructure to data pipelines
- Observability helps teams **monitor when things go wrong** and **profile how things change** over time; basically, knowing what’s going on under the hood
- There are **4 major types of observability**, from infrastructure to your whole business, and each carries different tools and responsible teams When teams can clearly see what’s happening in their systems – whether it’s an app, server, or part of the business – then they can make better decisions and ultimately get better outcomes

**Observability basics: what the hell is going on?** Observability is a complicated word for a pretty simple thing: companies want to know **what’s going on** with their systems and their business. Startups, and big companies, have a lot happening, from software to operations to why my campaigns aren’t converting this week, and the teams responsible for those things some visibility into them

In practice, observability usually manifests itself as a piece of software – one that you write yourself, or buy from someone else – that monitors the signals of your and spits some things out. The consumers of that observability are usually interested in **alerts** when things go wrong, but also **profiling and understanding** their better so they can make improvements.

- here’s one example:
  - **who**: your backend engineering team
  - **what**: how fast the database is running queries
  - **how**: query plans and times
  - **using**: built-in Postgres functions or Datadog
- here’s another example:
  - **who**: your data team
  - **what**: how clean a new data pipeline is
  - **how**: missing and null values
  - **using**: custom queries or Datafold
- and yet another example:
  - **who**: your product team
  - **what**: how larger customers are using a new feature
  - **how**: feature adoption and user growth
  - **using**: product analytics tools or Rupert

Examples abound. For a more conceptual backing though, you can pretty clearly split observability into **4 distinct categories**.

- Application observability
  - who: your frontend and backend engineering teams
  - what: how your application is performing for users
  - how: APM monitoring tools like DataDog or New Relic or Grafana
- Infrastructure observability
  - who: your backend engineering / DevOps engineering / SREs
  - what: the utilization and performance of your servers (and services, and other infra)
  - how: infrastructure monitoring tools like Datadog or Prometheus or Grafana
- Data observability
  - who: your data team, data engineers, data scientists
  - what: the quality of your data and systems built on top of it
  - how: custom queries and tools like Datafold
- Business observability
  - who: your functional teams (customer success, marketing, sales, etc), plus data owners
  - what: what's going on with your metrics and business events
  - how: dashboards and monitoring/alerting tools like Rupert

## Data Science Notebook

Data Science notebooks – or just “code notebooks”. [^6]

The typical use case for a data team’s code is much more __iterative__ than a software engineer’s. _For every line of code you write, you want to see the output to check if it did what you intended._

And another thing: data science and analytics work is often __exploratory__: the goal isn’t to create some sort of artifact (an app, a file). Instead, it’s to see what’s up – ask a few questions, make a chart, make a transformation. _Exploratory work is by definition iterative!_

Jupyter Notebooks are common place in data science. They allow a mixture of __“repeat, evaluation, loop”__ (REPL) code writing and documentation in a single place. They’re most commonly used for analysis and brainstorming purposes. [^14]

### Breaking down the notebook

The idea behind the notebook is that you write a block of code, run it, and see what the output is. Then you move on to the next block, and so on and so forth. We can break down what makes notebooks great into a few broad categories:

1. Notebooks are iterative: instead of writing all your code together, you can run it piece by piece and see what the output looks like.
2. Notebooks are visual: integrated tables, charts, and text make it easier to visualize what’s going on in your code.
3. Notebooks are organized: a great notebook is like a presentation. You can name each cell, reorder them, and add markdown to build a story.

Notebooks are pretty much the standard for how data teams do exploratory work today. Data teams use notebooks to:
- Clean up and categorize data
- Visualize data in charts and tables
- Build prototypes of ML models
- Create presentations and data stories for their stakeholders Report regularly on metrics with dashboards
- ...and a lot more

*__links:__*
- [Python Notebook]({% link docs/languages/python/notebook.md %})

## Webhook

__What Are Webhooks?__ [^11] Webhooks works like a notification system. Imagine giving your phone number to a friend who will call you whenever they have news. In the digital world, Webhooks do this through URLs instead of phone numbers. You give a third-party software your URL which they immediately notify when something happens.
- For example, if something changes in a GitHub repository, like a failed pipeline. GitHub will instantly send a message (webhook) to a specified URL of your choice. This URL should lead to your service waiting to hear the news.

What’s needed to use WebHooks is **an API that listens on the URL provided to the service**, to accept the information sent.

## [Authentication]({% link docs/algorithms-data-structures/authentication/index.md %})

Authentication in computer science refers to the process of verifying the identity of a user or system attempting to access a network, device, or application. It ensures that the entity requesting access is who they claim to be. In simple terms, authentication is like showing your ID to confirm your identity before entering a secure area. It involves providing credentials, such as a username and password, a fingerprint, or a security token, to prove that you are authorized to access the system. This verification step is crucial for maintaining security and preventing unauthorized access to sensitive information or resources.

## [Snippets]({% link docs/algorithms-data-structures/snippets/index.md %})

Snippets generally refer to small pieces or fragments of something larger.
In programming and software development, it means:
- Short reusable pieces of code
- Templates for commonly used code structures
- Code examples or samples

------ ------

[^1]: [How to conduct an A/B test experiment in Go using feature flags](https://medium.com/@chavezharris/how-to-conduct-an-a-b-test-experiment-in-go-using-feature-flags-660edef536f6)
[^2]: [Golang: How are Data Races, Context Switching, and Multi-Core Computing interlinked](https://mourya-g9.medium.com/golang-how-are-data-races-context-switching-and-multi-core-computing-related-cba8d17b542a)
[^3]: [Introduction to Microservices, Go-Kit, Grpc. Golang](https://medium.com/@edwinsiby/introduction-to-microservices-go-kit-grpc-golang-e187853d6c45)
[^4]: [Coroutines for Go](https://research.swtch.com/coro)
[^5]: [17 Writing Tasks You Can Do With AI](https://www.grammarly.com/blog/ai-writing-tasks/)
[^6]: [What’s a data science notebook?](https://technically.substack.com/p/whats-a-data-science-notebook)
[^7]: [Matching resumes with job postings using LLMs and Go](https://medium.com/@sausheong/matching-resumes-with-job-postings-using-llms-and-go-8ad9f0dfce6a)
[^8]: [Build your own RAG with Mistral-7B and LangChain](https://medium.com/@thakermadhav/build-your-own-rag-with-mistral-7b-and-langchain-97d0c92fa146)
[^9]: [An Introduction to gRPC: Building Distributed Systems With Efficiency and Scalability in Mind](https://medium.com/@semaphoreci/an-introduction-to-grpc-building-distributed-systems-with-efficiency-and-scalability-in-mind-c13024e4b0d3)
[^10]: [Go Services over gRPC](https://medium.kieranwarwick.dev/go-services-over-grpc-2fdc3bd7daeb)
[^11]: [Webhook Testing Without the Headache: A Developer’s Sanity-Saving Tutorial](https://itnext.io/webhook-testing-without-the-headache-a-developers-sanity-saving-tutorial-d6aea887f582)
[^12]: [Building, Evaluating and Tracking a Local Advanced RAG System : Mistral 7b + LlamaIndex + W&B](https://towardsdatascience.com/building-evaluating-and-tracking-a-local-advanced-rag-system-mistral-7b-llamaindex-w-b-5c9c69059f92)
[^13]: [Generative AI Terminology — An Evolving Taxonomy To Get You Started](https://pub.towardsai.net/generative-ai-terminology-an-evolving-taxonomy-to-get-you-started-4ca487bfe2d8)
[^14]: [Using IPython Jupyter Magic Commands to Improve the Notebook Experience](https://towardsdatascience.com/using-ipython-jupyter-magic-commands-to-improve-the-notebook-experience-f2c870cab356)
[^15]: [What is Observability?](https://read.technically.dev/p/what-is-observability)
