---
layout: minimal
title: RAG Search
nav_order: 99
---

[Return to main website]({{site.baseurl}}/).

**Harnessing AIChat and RAG with a Public GitHub Repository**

This guide demonstrates how to leverage AIChat's built-in vector database and full-text search capabilities to create a powerful knowledge base from an entire public GitHub repository. This enables Retrieval-Augmented Generation (RAG), allowing you to use AIChat to answer questions and provide assistance based on the repository's content.

**Creating a RAG Knowledge Base from a GitHub Repository:**

The following steps outline how to create a RAG knowledge base using AIChat, using the `ilima-snippets` example:

1.  **Initialize RAG:** Use the `.rag` command followed by a name for your RAG knowledge base.

    ```
    > .rag ilima-snippets
    ```

2.  **Configure Embedding Model:** Select an appropriate embedding model. The example uses `gemini:text-embedding-004`. Note the model's token limits, batch size, and pricing (if applicable).

    ```
    > Select embedding model: gemini:text-embedding-004 (max-tokens:2048; max-batch:100; price:0)
    ```

3.  **Set Chunk Size and Overlay:** Configure the chunk size and overlay for optimal retrieval.  A chunk size of 1500 and an overlay of 100 are used in the example.  Experiment with these values to find what works best for your repository.

    ```
    > Set chunk size: 1500
    > Set chunk overlay: 100
    ```

4.  **Add Documents:** Specify the GitHub repository's document path using a wildcard to include all files within the specified directory.  In this example, all documents within the `docs` directory of the `unapologetic-snippets` repository are added.

    ```
    > Add documents: https://github.com/igorlima/unapologetic-snippets/tree/main/docs/**
    ```

5.  **Exit RAG Configuration:**  Use the `.exit rag` command to finalize the RAG creation process.

    ```
    @ilima-snippets> .exit rag
    ```

**Managing Your RAG Knowledge Base:**

AIChat provides several commands for managing your RAG knowledge base:

*   **View Citation Sources:**  Use `.sources rag` to display the sources used in the last query. This helps verify the accuracy and relevance of the information.

    ```
    > .sources rag
    ```
    *   *Show citation sources used in last query*

*   **Edit Documents:** Use `.edit rag-docs` to add or remove documents from the existing RAG knowledge base.

    ```
    > .edit rag-docs
    ```

*   **Rebuild RAG:**  Use `.rebuild rag` to rebuild the knowledge base after making changes to the documents. This ensures that the RAG reflects the latest content.

    ```
    > .rebuild rag
    ```

*   **Show RAG Information:** Use `.info rag` to display information about the RAG knowledge base, such as the number of documents and the embedding model used.

    ```
    > .info rag
    ```

**Further Resources:**

*  **AIChat RAG Guide:**  For a more in-depth explanation, refer to the official [AIChat RAG Guide](https://github.com/sigoden/aichat/wiki/RAG-Guide) <sup>[+](https://github.com/sigoden/aichat/wiki/RAG-Guide/900d5644a72b33a0adba0e420bc3e645177a9f68)</sup>.
*  Leveraging *AIChat RAG* with Your [RC File Configuration Assistance Made Easy](https://igorlima.github.io/unapologetic-thoughts/pages/dots-mapping#aichat-rag-rc-knowledge-base)
*  Leveraging *AIChat RAG* with [ILIMA THOUGHTS](https://igorlima.github.io/unapologetic-thoughts/pages/search#aichat-rag-rc-knowledge-base)
