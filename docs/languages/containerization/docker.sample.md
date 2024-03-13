---
layout: default
title: Docker - Samples
nav_exclude: true
parent: Containerization
grand_parent: Programming Languages
permalink: /docs/languages/containerization/docker-samples
---

__[back]({% link docs/languages/containerization/docker.md %})__

- [AI samples]({% link docs/languages/containerization/docker.sample.ai.md %})

A sample containers designed to experiment with the latest technologies,
without cluttering my computer with unnecessary installations.

```bash
mkdir -p ~/workstation/gists

# VIM
git clone https://gist.github.com/90f67ef56912aa16306d9c4bd72b46c3.git ~/workstation/gists/my-vimrc
ln -s -f ~/workstation/gists/my-vimrc/.vimrc ~/.vimrc
git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim
# :PluginInstall

# TMUX
git clone https://gist.github.com/1def8b371fd4e9cff3c69bed35647dbc.git ~/workstation/gists/my-tmux-conf
ln -s -f ~/workstation/gists/my-tmux-conf/.tmux.conf ~/.tmux.conf
```

## ubuntu

<details markdown="block">
  <summary>
    Ubuntu
  </summary>

A Linux Docker container designed to experiment with the latest technologies,
without cluttering my computer with unnecessary installations.

```bash
# a docker for playing with new linux tools
# https://hub.docker.com/_/ubuntu
# $> docker image ls
#
# Automatically REMOVE THE CONTAINER when it exits
# $> mkdir -p ~/Downloads/temp
docker run \
  --name linux-learning --rm \
  --mount src=`realpath ~/Downloads/temp`,target=/home/local,type=bind \
  --workdir /home/local \
  -it ubuntu:22.04 bash
  # -p 8086:8086 \
  # --mount src=`pwd`,target=/home/local,type=bind \

# DETACH
# $> mkdir -p ~/Downloads/temp
docker run \
  --detach \
  --name linux-learning \
  --mount src=`realpath ~/Downloads/temp`,target=/home/local,type=bind \
  --workdir /home/local \
  -it ubuntu:22.04 bash
  # -p 8086:8086 \
  # --mount src=`pwd`,target=/home/local,type=bind \
# ...
docker exec -it linux-learning bash
docker stop linux-learning
docker start linux-learning
docker rm linux-learning
```

```sh
apt-get update
apt-get upgrade
apt-get install -y curl wget git vim

# gvm
# https://github.com/moovweb/gvm
apt-get install bison
apt-get install -y bsdmainutils binutils gcc make
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
# Please restart your terminal session or to get started right away run
#  `source /root/.gvm/scripts/gvm`
#
# GVM USAGE
# gvm listall
# gvm install go1.22.1 -B
# gvm use go1.22.1
# gvm list

# pyenv
# Simple Python version management
# https://github.com/pyenv/pyenv
apt-get install -y zlib1g-dev
apt-get install -y make build-essential libssl-dev zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev libncursesw5-dev xz-utils tk-dev libffi-dev liblzma-dev git 
curl https://pyenv.run | bash
# Load pyenv-virtualenv
#  `export PYENV_ROOT="$HOME/.pyenv"`
#  `[[ -d $PYENV_ROOT/bin ]] && export PATH="$PYENV_ROOT/bin:$PATH"`
#  `eval "$(pyenv init -)"`
#
# PYENV USAGE
# pyenv install --list
# pyenv install -v 3.10.10
# pyenv install -v 3.12.2
# pyenv local 3.12.2
# pyenv versions

# nvm
# Node Version Manager - POSIX-compliant bash script to manage multiple active node.js versions
# https://github.com/nvm-sh/nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
# start using nvm or run the following to use it now
#  `export NVM_DIR="$HOME/.nvm"`
#  `[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm`
#  `[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion`
#
# NVM USAGE
# nvm ls-remote
# nvm install v21.7.1
# nvm use v21.7.1

# rvm
# Ruby Version Manager (RVM)
# RVM is a command-line tool which allows you to easily install, manage, and
# work with multiple ruby environments from interpreters to sets of gems.
apt-get install -y bash make gnupg2 autoconf bison build-essential libssl-dev libyaml-dev libreadline6-dev zlib1g-dev libncurses5-dev libffi-dev libgdbm6 libgdbm-dev libdb-dev
# https://rvm.io/
# https://rvm.io/rvm/install
gpg2 --keyserver keyserver.ubuntu.com --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB
curl -sSL https://get.rvm.io | bash -s stable
#  `source /usr/local/rvm/scripts/rvm`
#  `echo "source /usr/local/rvm/scripts/rvm" >> ~/.bash_profile`
# RVM USAGE
# https://rvm.io/rvm/basics
# rvm install 3.1.3
# rvm use 3.1.3
#
# Jekyll
# bundle install
# bundle exec jekyll serve --config _config_dev.yml --host 0.0.0.0 --port 8086
#  http://localhost:8086/

# rust
# Install Rust and Cargo
# https://doc.rust-lang.org/cargo/getting-started/installation.html
curl https://sh.rustup.rs -sSf | sh
# SSL peer certificate
apt-get install -y libssl-dev
# run one of the following (note the leading DOT):
#  `. "$HOME/.cargo/env"            # For sh/bash/zsh/ash/dash/pdksh`
#  `source "$HOME/.cargo/env.fish"  # For fish`
```

----
<br/>
<!-- Ubuntu -->
</details>

## wiki

<details markdown="block">
  <summary>
    Github Wiki - Gollum
  </summary>

```bash
# https://github.com/gollum/gollum/wiki/Gollum-via-Docker
# https://github.com/gollum/gollum#configuration
# https://github.com/gollum/gollum
#
# tells Docker to use the current working directory as the git repository you
# want to serve your wiki from
docker run \
  --name gh-wiki --rm \
  -p 4567:4567 \
  -v $(pwd):/wiki \
  gollumwiki/gollum:v5.3.0

docker stop gh-wiki
```

----
<br/>
<!-- Github Wiki - Gollum -->
</details>

## OSX

<details markdown="block">
  <summary>
    Docker-OSX
  </summary>

Run macOS VM in a Docker! Run near native OSX-KVM in Docker! X11 Forwarding! CI/CD for OS X Security Research! Docker mac Containers.
- [Docker-OSX github](https://github.com/sickcodes/Docker-OSX)
  - _https://github.com/sickcodes/Docker-OSX_
- [Docker HUB](https://hub.docker.com/r/sickcodes/docker-osx)
  - _https://hub.docker.com/r/sickcodes/docker-osx_

----
<br/>
<!-- Docker-OSX -->
</details>

---

[^1]: [...](...)
