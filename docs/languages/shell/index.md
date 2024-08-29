---
layout: default
title: Shell
nav_order: 1
parent: Programming Languages
has_children: true
permalink: /docs/languages/shell
---

# Shell

A shell is a program that provides an interface for users to interact with the operating system. It interprets user commands and passes them to the system for execution. Shells can be command-line (CLI) or graphical (GUI) based. In software development, it’s often used as a terminal where developers type commands, making it a crucial tool for them. Examples include Bash, Zsh, TCSH, and KSH.

- **Bash:**
  - <details markdown="block"><summary><i>futher info...</i></summary>

    __Bash__, short for Bourne Again SHell, is a Unix shell and command language. It was created by Brian Fox for the GNU Project as a free software replacement for the Bourne shell1. It’s an sh-compatible shell that incorporates useful features from the Korn shell (ksh) and the C shell (csh).
    
    Here are some key features of Bash:
    
    - **Command-line editing**: Allows you to move around the command prompt and edit your command.
    - **Unlimited size command history**: You can access any command you’ve previously entered.
    - **Job control**: You can control multiple processes.
    - **Shell functions and aliases**: You can create your own commands and shortcuts.
    - **Indexed arrays of unlimited size**: You can create and manipulate arrays.
    - **Integer arithmetic in any base from two to sixty-four**: You can perform arithmetic operations.

    Most sh scripts can be run by Bash without modification. It’s widely used due to its power and versatility, and it’s the default shell for most Linux systems. It’s also highly customizable, which makes it a favorite among many developers.

    </details>

- **Zsh:**
  - <details markdown="block"><summary><i>futher info...</i></summary>

    __Zsh__, also known as the Z shell, is an extended version of the Bourne Shell (sh), with many enhancements and support for plugins and themes. It’s similar to Bash, but offers additional features such as:
      - **Automatic cd**: Just type the name of the directory.
      - **Recursive path expansion**: For example, `/u/lo/b` expands to `/usr/local/bin`.
      - **Automatic Syntax Highlighting**: Words, file extensions, and other things visible in the Terminal are color-coded.
      - **Spelling correction and approximate completion**: If you make a minor mistake typing a directory name, ZSH will fix it for you.
      - **Plugin and theme support**: ZSH includes many different plugin frameworks.

    Zsh is popular in the developer community due to its configurability and compatibility with Bash. It’s often used with a framework like Oh My Zsh2 to make it even more powerful and user-friendly.
    -------

    </details>


![image](https://github.com/user-attachments/assets/9eef088e-d426-4bc3-8136-94c9087c1bfb)
_<sup style="display:block;">image from Medium [^1]</sup>_

------ ------

[^1]: [5 Must-Have Command Line AI Tools](https://piotrzan.medium.com/5-must-have-command-line-ai-tools-839b0cf95c97)
