---
layout: default
title: Docker Samples
nav_exclude: true
parent: Containerization
grand_parent: Programming Languages
permalink: /docs/languages/containerization/docker-sample
---

__[back]({% link docs/languages/containerization/docker.md %})__

A sample containers designed to experiment with the latest technologies,
without cluttering my computer with unnecessary installations.

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
# $> mkdir -p ~/Downloads/temp
docker run \
  --name linux-learning --rm \
  --mount src=`realpath ~/Downloads/temp`,target=/home/local,type=bind \
  --workdir /home/local \
  -it ubuntu:22.04 bash
  #
  # --mount src=`pwd`,target=/home/local,type=bind \
```

```sh
apt-get update
apt-get upgrade
apt-get install -y curl git vim

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
```

----
<br/>
<!-- Ubuntu -->
</details>

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
