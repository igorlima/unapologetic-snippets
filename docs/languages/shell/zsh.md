---
layout: default
title: Z shell
nav_order: 2
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/zsh
---

`zsh`, colloquially known as the “Z shell,” is a souped-up version of the Bourne Shell (sh). Even without plugins, Zsh boasts of features missing in `sh`, such as:

- direct directory access: no need to type `cd`; just enter the directory name.
- advanced completion: for example, type `kill java`, hit tab, and Zsh will list all PIDs for Java processes.
- shared command history: Enter a command in one terminal, and it’s available in the history of another.

However, the true strength of Zsh lies in its extensibility. [^1]

If you’re on Ubuntu or another Debian-based system, you can use `apt`:

```sh
# Install zsh
$ sudo apt install zsh

# Check if everything works fine
$ zsh --version
zsh 5.9 (x86_64-ubuntu-linux-gnu)
```

set Zsh as your default shell and reboot your system:

```sh
chsh -s $(which zsh)
```

## ohmyzsh

- [GH link](https://github.com/ohmyzsh/ohmyzsh)

__Oh-My-Zsh__ is a renowned plugin framework for Zsh. Note: Oh-My-Zsh isn’t available via `apt`, so you'll need to run the following commands:

```sh
wget https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.dh
sh install.sh
```

For peace of mind, you can visit [GitHub](https://github.com/ohmyzsh/ohmyzsh) to review the installation script before executing it. With everything set up, it’s time to explore the world of Zsh plugins! [^1]



----

[^1]: [5 Plugins to Amp Up Your ZSH Journey](https://medium.com/itnext/zsh-essentials-5-plugins-to-amp-up-your-command-line-journey-47874b1b5740)
