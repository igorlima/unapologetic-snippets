---
layout: default
title: Miscellaneous
nav_order: 9
parent: Vim
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/vim/miscellaneous
---

# Miscellaneous

- [go back]({% link docs/languages/vim/index.md %})

Quick tips! I've got a list of handy shortcuts and commands that save the day when I need them. They're total time-savers!

- <details markdown="block"> <summary> Krisp AI </summary>

  - extract speaker and time from a line:
    - select the pattern: `SpeakerA | 00:00`
      - `:%s/^\([A-z0-9 ]\+| [0-9]\+\(:[0-9]\+\)\+\)$/- \1/gc`
      - `:%s/^\([A-z0-9 ]\+| [0-9]\+:[0-9]\+\)$/- \1/gc`
    - select lines that not start with the char `-`:
      - `:%s/\(^[^-].\+\)/  - \0/gc`
      - `:%s/\(^[^-].\+\)/  - \1/gc`
    - `\0` means the whole find and `\1` means a sub group under the find - the sub group is defined by parenthesis _"(subgroup)"_.
  </details>

- <details markdown="block"> <summary> diff </summary>


  ```sh
  # this is equivalent to calling vimdiff directly.
  vim -d file1 [file2 ...]
  ```
  
  **start** and **stop diff** of two files opened in 2 windows _(works in both, vertical and horizontal split)_:
  ```sh
  # You can also open vim in split-screen mode, with the -O option:-
  vim -O file1 [file2 ...]
  ```

  ```vim
  # to then turn on diff mode
  :windo diffthis
  # To then turn off diff mode
  :windo diffoff
  ```


  <details markdown="block"> <summary><strong><i>diff JSON files</i></strong></summary>

  ```sh
  vim -d <(git show HEAD~1:./snippets.json | jq . -) <(git show HEAD~2:./snippets.json | jq . -)
  vim -d <(git show f87d746:./snippets.json | jq . -) <(git show c147b39:./snippets.json | jq . -)
  
  diff -u <(git show f87d746:./snippets.json | jq . -) <(git show c147b39:./snippets.json | jq . -)
  vimdiff <(git show f87d746:./snippets.json | jq . -) <(git show c147b39:./snippets.json | jq . -)
  ```
  <a id="vimdiff-json-files"></a>
  </details>

  <a id="vimdiff"></a>
  </details>
