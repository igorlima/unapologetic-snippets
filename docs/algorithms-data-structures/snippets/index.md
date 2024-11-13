---
layout: default
title: Snippets
nav_order: 2
parent: Algorithms and Data Structures
has_children: false
has_toc: true
permalink: /docs/algorithms-and-data-structures/snippets
---

# Snippets

- [go back]({% link docs/algorithms-data-structures/index.md %}#snippets)

Snippets code refers to small, reusable pieces of source code or text that can
be easily inserted into a larger program or document. These snippets are
typically used to save time and reduce repetitive typing when coding or
writing.

- [script link](https://gist.githubusercontent.com/igorlima/90f67ef56912aa16306d9c4bd72b46c3/raw/snippets.sh)
  <sup>[+](https://gist.github.com/igorlima/90f67ef56912aa16306d9c4bd72b46c3#file-snippets-sh)</sup>
  <sup>[++](https://github.com/igorlima/unapologetic-snippets/tree/main/docs/algorithms-data-structures/snippets)</sup>

```sh
EDITOR=vim NAP_CONFIG=config.yaml nap
```
```sh
vim +"help ilima-help-snippets-code"
```

```sh
./snippets.sh <FOLDER>
./snippets.sh ~/.nap
./snippets.sh ~/workstation/github/unapologetic-snippets/docs/algorithms-data-structures/snippets
./snippets.sh ~/workstation/github/unapologetic-thoughts/snippets
```

-----

```sh
mkdir $(date +tmp-%Ya%mm%dd.%Hh%Mm%S)

# git clone --depth <depth> -b <branch> <repo_url>
git clone --depth 1 -b main https://github.com/igorlima/unapologetic-snippets.git

git pull origin main --depth=1
git fetch --depth=1

git diff --no-color > patch.patch
```

- KEY BINDINGS
  - | Action                              | Key         |
    | ----------------------------------- | :---------: |
    | Create a new snippet                | `n`         |
    | Edit selected snippet (in $EDITOR)  | `e`         |
    | Copy selected snippet to clipboard  | `c`         |
    | Paste clipboard to selected snippet | `p`         |
    | Delete selected snippet             | `x`         |
    | Rename selected snippet             | `r`         |
    | Set folder of selected snippet      | `f`         |
    | Set language of selected snippet    | `L`         |
    | Move to next pane                   | `tab`       |
    | Move to previous pane               | `shift+tab` |
    | Search for snippets                 | `/`         |
    | Toggle help                         | `?`         |
    | Quit application                    | `q ctrl+c`  |
- INSTALLATION
  - `go install github.com/maaslalani/nap@main`
    - Github repo:
      - `https://github.com/maaslalani/nap`
- JSON DIFF
  - https://semanticdiff.com/online-diff/json/
  - https://jsoncompare.org/
  - https://jsondiff.com/
- Compare JSON files
  - [how to `vimdiff`]({% link docs/languages/vim/miscellaneous.md %}#vimdiff)
    - `vim -d file1 [file2 ...]`
  - copy the current JSON
    - `cat snippets.json | pbcopy`
  - copy the previous JSON
    - `git show HEAD~1:./snippets.json | pbcopy`
