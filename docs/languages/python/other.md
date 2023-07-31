---
layout: default
title: Python Other
nav_order: 10
parent: Python
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/python/other
---

## How to Use Google Drive as a Knowledge Base For a ChatGPT Application.

Plenty of articles have been about using ChatGPT to chat with a PDF document.

So, instead of bringing your documents to ChatGPT; instead, bring ChatGPT to your documents using apps like Google Drive, Notion, and Dropbox as your knowledge base.

> "You don’t need to build a new knowledge base. You just need to use the knowledge base you already have." [^1]

Breakdown of what is been covering: [^1]

1. Connecting to Google Drive via Python

2. Extracting all of the documents under a specific folder

3. Storing the text data inside a vector database

4. Querying the data using ChatGPT [^2]


<details markdown="block">
  <summary>
    <sup>install dependencies</sup>
  </summary>
```sh
# pip show Flask
# pip show Flask-Cors
# pip show google-api-python-client
# pip show google-auth
# pip show google-auth-httplib2
# pip show google-auth-oauthlib

pip install Flask==2.3.2
pip install Flask-Cors==4.0.0
pip install google-api-python-client==2.93.0
pip install google-auth==2.22.0
pip install google-auth-httplib2==0.1.0
pip install google-auth-oauthlib==1.0.0
pip install PyPDF2==3.0.1
pip install tiktoken==0.3.3
pip install openai==0.27.0

# pip show qdrant_client
# pip show matplotlib
# pip show plotly
# pip show pandas
# pip show scipy
# pip show scikit-learn
pip install qdrant_client==1.3.1
pip install matplotlib==3.7.2
pip install plotly==5.15.0
pip install pandas==2.0.3
pip install scipy==1.11.1
pip install scikit-learn==1.3.0
```
----
<!-- install dependencies -->
</details>

- [`main.py`]({% link docs/languages/python/code/google-drive-oauth-app/main.md %})
- [`qdrant_test.py`]({% link docs/languages/python/code/google-drive-oauth-app/qdrant_test.md %})


```sh
# authorize
curl http://127.0.0.1:5000/authorize
```

```sh
# load data from GDocs
curl -X POST 'http://127.0.0.1:5000/load' \
-H 'Content-Type: application/json' \
-d '{"folder_path": "https://drive.google.com/drive/folders/1cAa_F9CnyCALEscyuBFimiLo9JsPZxkN"}'
```

```sh
# query
curl -X POST 'http://127.0.0.1:5000/query' \
-H 'Content-Type: application/json' \
-d '{"query": "what the difference between structured data and unstructured data"}'
```


-----

[^1]: [Using Google Drive as a Knowledge Base For Your ChatGPT Application](https://medium.com/better-programming/using-google-drive-as-a-knowledge-base-for-your-chatgpt-application-805962812547)
[^2]: [Original full code](https://github.com/htrivedi99/chatgpt-gdrive-article)
