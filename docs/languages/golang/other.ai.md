---
layout: default
title: Other AI
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/other-ai
---

__[back]({% link docs/languages/golang/other.md %})__

# AI models in Go

## Using Gemini models in Go with LangChainGo [^1]

<details markdown="block"><summary>LangChainGo examples with GoogleAI...</summary>
```golang
package main

import (
  "context"
  "fmt"
  "log"
  "os"

  "github.com/tmc/langchaingo/llms"
  "github.com/tmc/langchaingo/llms/googleai"
)

func main() {
  ctx := context.Background()
  apiKey := os.Getenv("API_KEY")
  llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
  if err != nil {
    log.Fatal(err)
  }

  prompt := "What is the L2 Lagrange point?"
  answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(answer)
}
```
</details>

----

[^1]: [Using Gemini models in Go with LangChainGo](https://eli.thegreenplace.net/2024/using-gemini-models-in-go-with-langchaingo/)
