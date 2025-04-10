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
- [script link](https://gist.githubusercontent.com/igorlima/90f67ef56912aa16306d9c4bd72b46c3/raw/snippets.sh)
  <sup>[+](https://gist.github.com/igorlima/90f67ef56912aa16306d9c4bd72b46c3#file-snippets-sh)</sup>
  <sup>[++](https://github.com/igorlima/unapologetic-snippets/tree/main/docs/algorithms-data-structures/snippets)</sup>

Snippets code refers to small, reusable pieces of source code or text that can
be easily inserted into a larger program or document. These snippets are
typically used to save time and reduce repetitive typing when coding or
writing.


<details markdown="block"><summary>nap</summary>

```sh
# cd ~/workstation/github/unapologetic-snippets/docs/algorithms-data-structures/snippets
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
- Compare JSON files
  - how to [`vimdiff`]({% link docs/languages/vim/miscellaneous.md %}#vimdiff) or [`diff` JSON files]({% link docs/languages/vim/miscellaneous.md %}#vimdiff-json-files)
    - `vim -d file1 [file2 ...]`
      ```sh
      # raw diff
      vim -d <(cat snippets.json | jq . -) <(git show HEAD~1:./snippets.json | jq . -)
      # sorting by date
      vim -d <(cat snippets.json | jq ". | sort_by(.date)" -) <(git show HEAD~1:./snippets.json | jq ". | sort_by(.date)" -)
      ```
      ```sh
      # advanced diff
      vim -d <(git show f87d746:./snippets.json | jq . -) <(git show c147b39:./snippets.json | jq . -)
      ```
  - <details markdown="block"> <summary> JSON diff <i>online</i> </summary>
  
    - copy the current JSON
      - `cat snippets.json | pbcopy`
    - copy the previous JSON
      - `git show HEAD~1:./snippets.json | pbcopy`
    - links:
      - [https://semanticdiff.com/online-diff/json/](https://semanticdiff.com/online-diff/json/)
      - [https://jsoncompare.org/](https://jsoncompare.org/)
      - [https://jsondiff.com/](https://jsondiff.com/)
      - <details markdown="block"> <summary>Diff Text</summary>
         
        [Diff Text](https://difftext.com/): A Simple Online Diff Checker Tool — Lets you compare by words, characters, or lines and you can toggle show all lines or only changed lines.
        </details>
    </details>

--------
<!-- nap -->
</details>


<details markdown="block"><summary>git-remote-s3</summary>

<a id="git-remote-s3"></a>

- [ilima vim help file](https://gist.githubusercontent.com/igorlima/90f67ef56912aa16306d9c4bd72b46c3/raw/ilima-help-snippets.txt) <sup>[+](https://gist.github.com/igorlima/90f67ef56912aa16306d9c4bd72b46c3/#file-ilima-help-snippets-txt)</sup>
- `git-remote-s3`
  - [setup](https://igorlima.github.io/unapologetic-thoughts/snippets/#git-remote-s3)
  - [tmuxp yaml](https://igorlima.github.io/unapologetic-thoughts/technical/2013/01/01/tmuxp.html#my-git-remote-s3-tmuxp-yaml)

--------
<!-- git-remote-s3 -->
</details>

💡 Got a spark or a bright idea? There's [**a place**](https://igorlima.github.io/unapologetic-thoughts/snippets/#git-remote-s3k) _to plant that seed and watch it grow!_ 🌱
