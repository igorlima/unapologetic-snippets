---
layout: minimal
title: Google Drive as a Knowledge Base - main.py
nav_exclude: true
---

This snippet shows how to connect with Google Drive. The same idea can be applied to Notion, Salesforce, Zendesk, Dropbox, Slack, Gmail, Rippling, and so on... In doing so, the application can chat with a private data using natural language.

This API has four endpoints: [^2]
1. An authorize endpoint
  - initializes the process to connect to Google Drive
2. A callback endpoint
  - verifies user identity and creates credentials
    - a `gdrive_credentials.txt` file should be in the application folder after the authorization. In doing so, it reuses the credentials text file for future API calls without re-authenticating through the authorization flow again.
3. A load data endpoint
  - loads all of the documents from Google Drive into a vector database
    - it extracts all of the private documents under the given folder
    - and then it stores the text data inside a vector databse
4. A query endpoint
  - allows us to chat with our documents
    - it requires an OpenAI developer account for this step because it uses their embedding API.
    - it also needs a vector database - the idea is to upload all of the text from the private documents into a vector database. Whenever a user queries, it finds similar content in the databse and uses that as context to answer the user's question.
  - it performs a vector search to find the pieces of text that match a given query and retrieve only those. These pieces of text will serve as a __context__, which ChatGPT will use to answer the question.

----

- Google Drive credentials [^1]
  - Head over to `https://console.cloud.google.com/apis/credentials` and create a new OAuth Client ID credential
  - Select Web Application as the Application Type
  - Give the credential any name you want. You don’t need to add anything for Authorized JavaScript origins
  - Under Authorized redirect URIs, add the following URL: `http://127.0.0.1:5000/oauth/redirect`
  - Click Create
  - Download the credentials JSON file for the next step
    - `{"web":{"client_id":"72246653...", "project_id": ...}}`
- OpenAI APY key
  - an OpenAI developer account is required because it uses their embedding API.
    - `https://platform.openai.com/account/api-keys`
      - [link](https://help.openai.com) for help if needed
    - `YOUR-OPENAI-API-KEY`

```python
# https://github.com/htrivedi99/chatgpt-gdrive-article
# https://medium.com/better-programming/using-google-drive-as-a-knowledge-base-for-your-chatgpt-application-805962812547

import json
import io
import webbrowser
from flask import Flask, request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from urllib.parse import urlparse
from urllib.parse import parse_qs
from collections import deque
from PyPDF2 import PdfReader
from flask_cors import CORS

import tiktoken
from qdrant_test import QdrantVectorStore
import openai
openai.api_key = "YOUR-OPENAI-API-KEY"


SCOPES = ['https://www.googleapis.com/auth/drive']
client_secrets = {"web":{"client_id":"72246653...", "project_id": ...}}

app = Flask(__name__)
CORS(app)

@app.route("/authorize", methods=['GET'])
def authorize_google_drive():
    flow = InstalledAppFlow.from_client_config(
        client_secrets,
        SCOPES,
        redirect_uri="http://127.0.0.1:5000/oauth/redirect"
    )

    authorization_url, state = flow.authorization_url(prompt='consent')
    webbrowser.open(authorization_url)
    return authorization_url

@app.route("/oauth/redirect", methods=['POST', 'GET'])
def redirect_callback():
    authorization_response = request.url
    print("authorization response: ", authorization_response)
    parsed_url = urlparse(authorization_response)
    auth_code = parse_qs(parsed_url.query)['code'][0]
    print("auth code: ", auth_code)

    flow = InstalledAppFlow.from_client_config(
        client_secrets,
        SCOPES,
        redirect_uri="http://127.0.0.1:5000/oauth/redirect"
    )

    flow.fetch_token(code=auth_code)
    credentials = flow.credentials
    credentials_string = credentials.to_json()
    with open("gdrive_credentials.txt", "w") as text_file:
        text_file.write(credentials_string)

    return "Google Drive Authorization Successful!"

def get_folder_id_from_url(url: str):
    url_path = urlparse(url).path
    folder_id = url_path.split("/")[-1]
    return folder_id

def list_files_in_folder(service, folder_id):
    query = f"'{folder_id}' in parents"
    results = service.files().list(q=query, fields="nextPageToken, files(id, name, mimeType, webViewLink)").execute()
    items = results.get("files", [])
    return items

def download_pdf(service, file_id):
    request = service.files().get_media(fileId=file_id)
    file = io.BytesIO(request.execute())
    return file

def extract_pdf_text(pdf_file):
    reader = PdfReader(pdf_file)
    text = ''
    for page_num in range(len(reader.pages)):
        text += reader.pages[page_num].extract_text()
    return text

def get_documents_from_folder(service, folder_id):
    folders_to_process = deque([folder_id])
    documents = []

    while folders_to_process:
        current_folder = folders_to_process.popleft()
        items = list_files_in_folder(service, current_folder)

        for item in items:
            mime_type = item.get("mimeType", "")

            if mime_type == "application/vnd.google-apps.folder":
                folders_to_process.append(item["id"])
            elif mime_type in ["application/vnd.google-apps.document", "application/pdf"]:
                # Retrieve the full metadata for the file
                file_metadata = service.files().get(fileId=item["id"]).execute()
                mime_type = file_metadata.get("mimeType", "")

                if mime_type == "application/vnd.google-apps.document":
                    doc = service.files().export(fileId=item["id"], mimeType="text/plain").execute()
                    content = doc.decode("utf-8")
                elif mime_type == "application/pdf":
                    pdf_file = download_pdf(service, item["id"])
                    content = extract_pdf_text(pdf_file)

                if len(content) > 0:
                    documents.append(content)

    return documents

def chunk_tokens(document: str, token_limit: int = 200):
    tokenizer = tiktoken.get_encoding(
        "cl100k_base"
    )

    chunks = []
    tokens = tokenizer.encode(document, disallowed_special=())

    while tokens:
        chunk = tokens[:token_limit]
        chunk_text = tokenizer.decode(chunk)
        last_punctuation = max(
            chunk_text.rfind("."),
            chunk_text.rfind("?"),
            chunk_text.rfind("!"),
            chunk_text.rfind("\n"),
        )
        if last_punctuation != -1:
            chunk_text = chunk_text[: last_punctuation + 1]
        cleaned_text = chunk_text.replace("\n", " ").strip()

        if cleaned_text and (not cleaned_text.isspace()):
            chunks.append(
                {"text": cleaned_text}
            )

        tokens = tokens[len(tokenizer.encode(chunk_text, disallowed_special=())):]

    return chunks

@app.route("/load", methods=['POST'])
def load_docs_from_drive():
    data = request.json
    google_drive_folder_path = data.get('folder_path')
    if not google_drive_folder_path:
        return {"msg": "A folder path must be provided in order to load google drive documents"}

    with open('gdrive_credentials.txt') as f:
        line = f.readline()
    credentials_json = json.loads(line)

    creds = Credentials.from_authorized_user_info(
        credentials_json
    )

    if not creds.valid and creds.refresh_token:
        creds.refresh(Request())
        credentials_string = creds.to_json()
        with open("gdrive_credentials.txt", "w") as text_file:
            text_file.write(credentials_string)

    service = build('drive', 'v3', credentials=creds)
    folder_id = get_folder_id_from_url(google_drive_folder_path)
    documents = get_documents_from_folder(service, folder_id)
    chunks = []
    for doc in documents:
        document_chunks = chunk_tokens(doc)
        chunks.extend(document_chunks)

    vector_store = QdrantVectorStore(collection_name="google-drive-docs")
    vector_store.upsert_data(chunks)

    return "docs loaded"

def chatgpt_answer(question, context):
    prompt = f"""

        Use ONLY the context below to answer the question. If you do not know the answer, simply say I don't know.

        Context:
        {context}

        Question: {question}
        Answer:"""

    res = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": "You are a question answering chatbot"},
            {"role": "user", "content": prompt}
        ]
    )
    return res['choices'][0]['message']['content']

@app.route("/query", methods=['POST'])
def query_knowledge_base():
    data = request.json
    query = data.get('query')
    vector_store = QdrantVectorStore(collection_name="google-drive-docs")
    results = vector_store.search(query)

    context = ""
    for entry in results:
        text = entry.get('text')
        context += text

    llm_answer = chatgpt_answer(query, context)
    print(llm_answer)
    return llm_answer


if __name__ == "__main__":
    app.run()
```

-----

[^1]: [Using Google Drive as a Knowledge Base For Your ChatGPT Application](https://medium.com/better-programming/using-google-drive-as-a-knowledge-base-for-your-chatgpt-application-805962812547)
[^2]: [Original full code](https://github.com/htrivedi99/chatgpt-gdrive-article)
