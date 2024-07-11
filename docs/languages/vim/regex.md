---
layout: default
title: Regex
nav_order: 2
parent: Vim
grand_parent: Programming Languages
has_children: true
permalink: /docs/languages/vim/regex
---

# Regex Vim

- [go back]({% link docs/languages/vim/index.md %})

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>


## How to do an inverse search in Vim?

_(i.e., get lines not containing pattern)_

<details markdown="block">
<summary>
  search for the lines not containing a pattern
</summary>

To search for the lines not containing `Person A` at the beginning:
- `/^\(Person A\)\@!`
- `/\v^(Person A)@!`
- replace it to something else:
  - `:%s/^\(Person A\)\@!/- \0/gc`
  - <details markdown="block">
    <summary><sub><i>text to play with</i></sub></summary>
  
    ```txt
    Person A: Hey, how's it going?
    Person B: Not too bad, thanks. How about you?
    Person A: I'm doing well! Have you finished that book you were reading?
    Person B: Yes, I have. It was quite an interesting read.
    Person A: That's great! Do you recommend it?
    Person B: Absolutely, it's a must-read if you're into mystery novels.
    Person A: Sounds intriguing. I'll definitely check it out. Thanks for the recommendation!
    Person B: You're welcome! Let me know what you think once you've read it.
    ```
    </details>

To search for the lines not containing `Person A` and `Speaker B` at the beginning:
- `/^\(Person A\|Person B\)\@!`
- `/\v^(Person A|Person B)@!`
- replace it to something else:
  - `:%s/^\(Person A\|Person B\)\@!/- \0/gc`
  - <details markdown="block">
    <summary><sub><i>text to play with</i></sub></summary>

    ```txt
    Person A: Hi guys, how are you both doing today?
    Person B: I'm doing well, thanks for asking. How about you, Person C?
    Person C: I'm good too. Thanks, Person B. How about you, Person A?
    Person A: I'm great, thank you! Have either of you seen the new movie that just came out?
    Person B: Not yet, but I've heard good things about it. What about you, Person C?
    Person C: I actually saw it yesterday. It was really good!
    Person A: That's awesome! I'll have to check it out soon.
    Person B: Sounds like a plan. Maybe we can all go together next time.
    Person C: That sounds like a great idea!
    ```
    </details>

------

</details>

