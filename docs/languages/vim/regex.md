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

<!-- search for the lines not containing a pattern -->
</details>
<!-- H2: How to do an inverse search in Vim? -->

## lookahead and lookbehind

<sup><sub>If you want to search for a pattern only when it occurs next to another pattern, use the regex features “lookahead” and “lookbehind” (collectively “lookaround”). If you want to search for a pattern only when it doesn't occur next to another, use their complements, “negative lookahead” and “negative lookbehind” (“negative lookaround”).</sub></sup>

<details markdown="block"> <summary> intro </summary>

Lookahead and lookbehind are two types of zero-width assertions in regular expressions (regex). They do not match characters but instead assert whether a match is possible.
- **Lookahead**:
  - Lookahead assertions check if a pattern matches without including the matched text in the result. There are two types of lookahead:
    - Positive lookahead: Ensures that the pattern matches.
    - Negative lookahead: Ensures that the pattern does not match.
- **Lookbehind**:
  - Lookbehind assertions check if a pattern matches before the current position without including the matched text in the result. There are two types:
    - Positive lookbehind: Ensures that the pattern matches.
    - Negative lookbehind: Ensures that the pattern does not match.
- **Key differences**:
  - Direction: Lookahead checks the string to the right of the current position, while lookbehind checks the string to the left.
  - Length: Lookbehind requires a fixed-length pattern, while lookahead can handle variable-length patterns.
- <details markdown="block"> <summary> Examples </summary>
   
  1. **Lookahead**: Match "hello" only if followed by "world".
    - ...
      ```
      /hello(?=world)/
      ```
  2. **Lookbehind**: Match "world" only if preceded by "hello".
    - ...
      ```
      /(?<=hello)world/
      ```
  3. **Negative lookahead**: Match "hello" only if not followed by "world".
    - ...
      ```
      /hello(?!world)/
      ```
  4. **Negative lookbehind**: Match "world" only if not preceded by "hello".
    - ...
      ```
      /(?<!hello)world/
      ```
   
  <!-- Example -->
  </details>

---------
<!-- intro -->
</details>

<details markdown="block"> <summary> expand </summary>

- <details markdown="block"> <summary> check the <strong>vim help</strong> </summary>
  
  - ...
    ```vim
    " positive lookahead
    :h \@=
    " negative lookahead
    :h \@!
    " positive lookbehind
    :h \@<=
    " negative lookbehind
    :h \@<!
    ```
  - ...
    ```vim
    \@=     Matches the preceding atom with zero width. {not in Vi}
            Like "(?=pattern)" in Perl.
            Example             matches
            foo\(bar\)\@=       "foo" in "foobar"
            foo\(bar\)\@=foo    nothing
    ```
  - without wildcards:
    ```
    Positive Lookahead:  \(find this\)\(followed by this\|or that\)\@=
    Negative Lookahead:  \(find this\)\(not followed by this\|or that\)\@!
    Positive Lookbehind: \(preceded by this\|or that\)\@<=\(find this\)
    Negative Lookbehind: \(not preceded by this\|or that\)\@<!\(find this\)
    ```
  - with wildcards:
    ```
    Positive lookahead:  \(find this\)\(.*\(eventually followed by this\|or that\)\)\@=
    Negative lookahead:  \(find this\)\(.*\(not eventually followed by this:\|or that\)\)\@!
    Positive lookbehind: \(\(eventually preceded by this\|or that\).*\)\@<=\(find this\) 
    Negative lookbehind: \(\(not eventually preceded by this\|or that\).*\)\@<!\(find this\)
    ```
    <sup>**Note**: For the wildcard versions, the extra parentheses are required so that the wildcard is excluded from the alternatives group, but is included in the lookaround group. This prevents duplicating the wildcards for every alternative. <sup>[+](https://stackoverflow.com/questions/37166743/lookarounds-positive-negative-lookbehind-lookahead)</sup></sup>
   
   
  <!-- check the vim help -->
  </details>
- regular expression to find lines containing multiple specific words or patterns in any arbitrary order
  - to find the word `support` and `mail` at once:
    - ...
      ```vim
      /^\(.*support\)\@=\(.*mail\)\@=
      ```
    - ...
      ```vim
      " for insensitive case
      /^\(.*support\)\@=\(.*mail\)\@=\c
      ```
  - additional hint:
    - but do not contain the word `service` in the same line:
      - ...
        ```
        /^\(.*support\)\@=\(.*mail\)\@=\(.*service\)\@!
        ```

---------
<!-- expand -->
</details>
<!-- H2: lookahead and lookbehind -->

## Miscellaneous

<details markdown="block"> <summary> expand </summary>

case insensitive search in Vim:
  - You can use the `\c` escape sequence anywhere in the pattern. For example:
    - `/\ccopyright` or `/copyright\c` or even `/copyri\cght`
  - To do the inverse (case sensitive matching), use `\C` (capital C) instead.

---------
<!-- expand -->
</details>
<!-- H2: Miscellaneous -->
