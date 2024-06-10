---
layout: default
title: Regex
nav_order: 6
parent: Programming Languages
has_children: false
has_toc: true
permalink: /docs/languages/regex
---

# Regular Expressions

Regular expressions are a powerful tool for pattern matching and text manipulation. They are used in many programming languages, text editors, and command-line tools.

- testing regex with specialized tools
  - [rexi](https://github.com/royreznik/rexi)
    - Terminal UI for Regex Testing. An interactive CLI for testing regex from terminal.
    - <details markdown="block"><summary><sup><i>more</i></sup></summary>

      ```sh
      # installation
      pip3 install rexi
      # usage
      cat /etc/hosts | rexi
      ```
      </details>
  - [sig](https://github.com/ynqa/sig)
    - Interactive grep (for streaming)
    - <details markdown="block"><summary><sup><i>more</i></sup></summary>
      
      `sig` is an interactive grep tool designed for streaming data. It offers
      real-time search updates as data streams in, making it easy to navigate
      large amounts of dynamic data.

      ```sh
      brew install ynqa/tap/sigrs
      ```
      
      ```sh
      # press `<ctrl-f>` to enter Archived mode
      curl https://raw.githubusercontent.com/ynqa/sig/0db6dda7ddc6e9b9197e0d249e3f31e73b3d0d50/README.md |& sig
      # press `<ctrl-f>` to enter Archived mode
      sig --cmd "curl https://raw.githubusercontent.com/ynqa/sig/0db6dda7ddc6e9b9197e0d249e3f31e73b3d0d50/README.md"
      # archived mode
      sig -a --cmd "ls -la"
      ls -la |& sig -a
      ```
       
      This tool also features keyboard shortcuts, an archived mode for dealing
      with backwards searching issues related to piping processes. In archived
      mode, `sig` saves the latest N entries and allows searching through them
      based on given key inputs. This mode also enables users to search through
      static data, such as files.
        
      `sig` works best when used in environments where data is streamed, such as
      the output of kubernetes logs or the output of websocket data.
      </details>
  - [RegExr](https://regexr.com/)
    - RegExr is an online tool to learn, build, & test Regular Expressions (RegEx / RegExp).
- **programming languages**
  - JavaScript
    - [MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions)
    - internal link
      - [browser JS]({% link docs/languages/js/regex.md %})

----

[^1]: [...](...)
