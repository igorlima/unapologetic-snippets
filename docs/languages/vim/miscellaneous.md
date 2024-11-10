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
