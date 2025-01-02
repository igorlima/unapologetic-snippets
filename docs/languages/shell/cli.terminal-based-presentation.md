---
layout: minimal
title: Shell - Terminal-Based Presentation
nav_exclude: true
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/cli-terminal-based-presentation/
---

[..]({% link docs/languages/shell/cli.md %})

# Terminal-Based Presentation

__[back]({% link docs/languages/shell/cli.md %})__

- <details markdown="block"><summary>Diagram Tools</summary>
   
  <a id="diagram-tools"></a>
  - For ASCII diagrams, see: [^1]
    - [PlantUML](#plant-uml) for sequence diagrams
    - [Graph::Easy](#graph-easy) for component diagrams
  <hr>
  <!-- Diagram Tools -->
  </details>

# Slides

Slides is a terminal-based tool that brings the art of presentations straight to the console[^1]. Let's explore how to get started and maximize this tool!

  <details markdown="block"><summary>installation...</summary>

  Here's how you can install it:
  - **MacOS**: `brew install slides`.
  - **Golang**: `go install github.com/maaslalani/slides@latest`.
  - **from source**:
    ```sh
    git clone https://github.com/maaslalani/slides.git
    cd slides
    go install
    ```
  <hr>
  <!-- Installation -->
  </details>

  <details markdown="block"><summary>example...</summary>
  
  Here's a quick glimpse of what Slides can do:
  
  ````md
  ---
  author: Igor Lima
  date: MMMM dd, YYYY
  paging: Slide %d / %d
  ---

  # Welcome to Slides
  A terminal-based presentation tool
  
  ---
  
  ## Everything is markdown
  This entire presentation is a markdown file.
  
  ---
  
  ## Everything happens in your terminal
  Create slides and present them without ever leaving your terminal.
  
  ---
  
  ## Code execution
  ```go
  package main
  
  import "fmt"
  
  func main() {
    fmt.Println("Execute code directly inside the slides")
  }
  ```
  
  You can execute code inside your slides by pressing `<C-e>`, and the output will appear at the end of the current slide.
  
  ---
  
  ## Pre-process slides
  
  Add a code block with three tildes (`~`) to run a command *before* displaying the slides. The text inside the block is passed as `stdin` to the command, and it gets replaced with the command's `stdout`.
  
  ```txt
  ~~~graph-easy --as=boxart
  [ A ] - to -> [ B ]
  ~~~
  ```
  
  The above will transform into:
  
  ```txt
  ┌───┐  to   ┌───┐
  │ A │ ────> │ B │
  └───┘       └───┘
  ```
  
  For security, ensure the file has execution permissions. Use `chmod` to set this up:
  
  ```bash
  chmod +x file.md
  ```
  ````
  <hr>
  <!-- Example -->
  </details>

To start the presentation, run:
- `slides presentation.md`
- `curl http://example.com/slides.md | slides`

Enjoy creating and sharing presentations easily on the terminal!

  <details markdown="block"><summary>references...</summary>
  
  - [GitHub](https://github.com/maaslalani/slides)
  - [Some examples](https://github.com/maaslalani/slides/tree/c6eea3330053045cede8d65ee1086fb5d4db6c43/examples)
  - [Slides Homepage](https://maaslalani.com/slides/)
  - [Terminal-Based Presentations](https://itnext.io/terminal-based-presentations-66c9f0c9b4a3) [^1]
  - [Diagram Tools](#diagram-tools)
  <hr>
  <!-- references -->
  </details>


<br>
<br>

# Other Tools

- <details markdown="block"><summary> Graph-Easy: a CLI Tool for Rendering <strong><i>diagrams in Terminal</i></strong> </summary>
  
  <a id="graph-easy"></a>
  The `graph-easy` tool is an utility that can render and convert graphs in various formats. Its ability to create diagrams directly in the terminal using ASCII art is fascinating. Additionally, it can convert and render graphs in formats like HTML, SVG, PNG, and so forth.
  
  This tool is crafted in Perl and can be installed via CPAN.
  
  - <details markdown="block"><summary><strong>installation instructions</strong></summary>
    
    1. **download and extract:**
    
       ```sh
       # https://metacpan.org/pod/Graph::Easy
       # https://metacpan.org/pod/Graph::Easy::As_svg
       # https://github.com/shlomif/Graph-Easy-As_svg
       # https://metacpan.org/dist/Graph-Easy
       curl -O https://cpan.metacpan.org/authors/id/S/SH/SHLOMIF/Graph-Easy-0.76.tar.gz
       tar -xzf Graph-Easy-0.76.tar.gz
       cd Graph-Easy-0.76
       ```
    
    2. **build and test:**
    
       ```sh
       perl Makefile.PL
       make test
       ```
    
    3. **install the package _(as root)_:**
    
       ```sh
       sudo make install
       ```
    <hr>
    <!-- installation instructions -->
    </details>
  
  - <details markdown="block"><summary>SVG rendering setup</summary>
     
    For rendering diagrams as SVG, an additional package is required. Here's how to install it:
    
    1. **install the SVG package:**
    
       ```sh
       # http://bloodgate.com/perl/graph/
       # https://metacpan.org/pod/Graph::Easy#as_svg()
       perl -MCPAN -e 'shell'
       > install Graph::Easy::As_svg
       > q
       ```
    
    2. **try rendering an SVG:**
    
     ```sh
     # https://linux.die.net/man/1/graph-easy
     # https://metacpan.org/dist/Graph-Easy/view/bin/graph-easy
     echo "[ Bonn ] - car -> [ Berlin ]" | graph-easy --as=svg --output=output.svg
     ```
    <hr>
    <!-- SVG rendering setup -->
    </details>
  
  - <details markdown="block"><summary><strong>rendering examples</strong></summary>
     
    Once you have the package installed, you can give it a go with rendering simple graphs:
    
    1. **render as ASCII:**
    
       ```sh
       echo "[ Bonn ] - car -> [ Berlin ]" | graph-easy
       ```
       ```txt
       +------+  car   +--------+
       | Bonn | -----> | Berlin |
       +------+        +--------+
       ```
    
    2. **render as PNG:**
    
       ```sh
       echo "[ Bonn ] - car -> [ Berlin ]" | graph-easy --as=png
       open graph.png
       ```
    <hr>
    <!-- rendering examples -->
    </details>
   
  <hr>
  <!-- Graph-Easy: a CLI Tool for Rendering diagrams in Terminal/Console -->
  </details>

- <details markdown="block"><summary> PlantUML: another CLI Tool for rendering diagrams <sub><sup><i>(but not in terminal)</i></sup></sub></summary>

  <a id="plant-uml"></a>
  PlantUML is a tool used to create a wide variety of diagrams with a simple language. It's useful for generating diagrams in ASCII format for display in terminals. With PlantUML, you can create UML diagrams, sequence diagrams, flowcharts, and more.
  
  <details markdown="block"><summary>References...</summary>
    
  - **Home Page:**
    - [PlantUML Home](https://plantuml.com/)
    - [Getting Started](https://plantuml.com/starting)
    - [Sequence Diagrams](https://plantuml.com/sequence-diagram)
  
  - **GitHub:**
    - [PlantUML Online Server](https://github.com/plantuml/plantuml-server)
  
  - **Online Editors:**
    - [PlantUML Editor](https://www.plantuml.com/plantuml/uml/SoWkIImgAStDuNBCoKnELT2rKt3AJrAmKiX8pSd9vt98pKi1IG80)
    - [Text Editor](https://www.plantuml.com/plantuml/txt/SoWkIImgAStDuNBCoKnELT2rKt3AJrAmKiX8pSd9vt98pKi1IG80)
    - [Editor](https://editor.plantuml.com/)
  <hr>
  <!-- references -->
  </details>
  
  <details markdown="block"><summary>How to run and use PlantUML locally...</summary>
  
  To get started with PlantUML locally, run the PlantUML server container using Docker:
  
  ```sh
  # run the PlantUML server container
  docker run \
    -d --name plant-uml --rm \
    -p 3005:8080 \
    plantuml/plantuml-server:jetty
  
  # stop the container when you're done
  docker ps
  docker stop [CONTAINER_ID]
  docker stop plant-uml
  ```
  <hr>
  <!-- how to run and use PlantUML locally -->
  </details>
  
  <details markdown="block"><summary><strong>Examples...</strong></summary>
  
  - **[Sequence Diagram](https://plantuml.com/sequence-diagram):**
    - _input_
      ```txt
      @startuml
      Alice -> Bob: Authentication Request
      Bob --> Alice: Authentication Response
      
      Alice -> Bob: Another authentication Request
      Alice <-- Bob: Another authentication Response
      @enduml
      ```
    - _output_
      ```txt
       ┌─────┐                           ┌───┐
       │Alice│                           │Bob│
       └──┬──┘                           └─┬─┘
          │    Authentication Request      │  
          │───────────────────────────────>│  
          │                                │  
          │    Authentication Response     │  
          │<───────────────────────────────│  
          │                                │  
          │Another authentication Request  │  
          │───────────────────────────────>│  
          │                                │  
          │Another authentication Response │  
          │<───────────────────────────────│  
       ┌──┴──┐                           ┌─┴─┐
       │Alice│                           │Bob│
       └─────┘                           └───┘
      ```
  - Explore more diagram types:
    - [JSON Data Diagram](https://plantuml.com/json)
    - [YAML Data Diagram](https://plantuml.com/yaml)
    - [Visualizing Regex](https://plantuml.com/regex)
    - [Project Schedule Diagram](https://plantuml.com/gantt-diagram)
    - [Chronology Diagram](https://plantuml.com/chronology-diagram)
    - [Mind Map Diagram](https://plantuml.com/mindmap-diagram)
    - [Activity Diagram](https://plantuml.com/activity-diagram-beta)
      - [Legacy Activity Diagrams](https://plantuml.com/activity-diagram-legacy)
    - [State Diagram](https://plantuml.com/state-diagram)
    - [Timing Diagram](https://plantuml.com/timing-diagram)
    - [Math Notation](https://plantuml.com/ascii-math)
    - [Hierarchy Diagram](https://plantuml.com/wbs-diagram)
      - [Entity Relationship Diagrams](https://plantuml.com/er-diagram)
    - [UI MockUp - Wireframe](https://plantuml.com/salt)
  <hr>
  <!-- Examples -->
  </details>
  
  PlantUML is a tool that makes diagramming easy.
     
  <hr>
  <!-- PlantUML: a CLI Tool for Rendering diagrams -->
  </details>

----

[^1]: [Terminal-Based Presentations](https://itnext.io/terminal-based-presentations-66c9f0c9b4a3)
