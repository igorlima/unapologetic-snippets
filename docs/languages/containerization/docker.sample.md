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

## programming languages

<details markdown="block">
  <summary>
    Programming Languages
  </summary>

A Docker container to experiment and learn new things without cluttering my computer with unnecessary installations.

```sh
mkdir -p ~/Downloads/temp
```

### Python

<details markdown="block">
  <summary>
    Python
  </summary>

```bash
# a docker for playing with Python
# https://hub.docker.com/_/python
# $> docker image ls
#
# Automatically REMOVE THE CONTAINER when it exits
docker run \
  --name python-learning --rm \
  --mount src=`realpath ~/Downloads/temp`,target=/home/local,type=bind \
  --workdir /home/local \
  -it python:3.12.2 bash
  # -p 8086:8086 \
  # --mount src=`pwd`,target=/home/local,type=bind \

# DETACH
docker run \
  --detach \
  --name python-learning \
  --mount src=`realpath ~/Downloads/temp`,target=/home/local,type=bind \
  --workdir /home/local \
  -it python:3.12.2 bash
  # -p 8086:8086 \
  # --mount src=`pwd`,target=/home/local,type=bind \
# ...
docker exec -it python-learning bash
docker stop python-learning
docker start python-learning
docker rm python-learning
```
<!-- Python -->
</details>

-----

<!-- Programming Languages -->
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

## Oracle Cloud - IDE

<details markdown="block"><summary>CentOS</summary>

### Jupyter Notebook

<details markdown="block"><summary>CentOS - Jupyter</summary>

- [My notebooks]({% link docs/languages/python/notebook.md %}#centos)

```sh
#!/bin/bash

# chmod +x exec-docker-bash.sh
sudo docker exec -it python-learning bash

# see the ip address of the container
# $> curl ifconfig.me

# Run Jupyter Notebook
# $> jupyter notebook --allow-root --ip=0.0.0.0 --port=8888
```
```sh
#!/bin/bash

# chmod +x exec-docker-notebook-d.sh
sudo docker exec -d \
  -e GOPATH=/root/go \
  -e PATH=/root/go/bin:/usr/local/bin:/usr/local/sbin:/usr/sbin:/usr/bin:/sbin:/bin \
  -it python-learning \
  jupyter notebook \
    --allow-root --ip=0.0.0.0 --port=8888 \
    --NotebookApp.token='0fd8d941-9f43-49ef-b987-3efb21966bb9'

# generating UUID
# $> echo `uuidgen | tr '[:upper:]' '[:lower:]'`

# See the IP address of the container
# http://155.248.192.51:8088/
# $> curl ifconfig.me

# Run Jupyter Notebook
# $> jupyter notebook --allow-root --ip=0.0.0.0 --port=8888 --NotebookApp.token='xxxxxx'
```
```sh
#!/bin/bash

# chmod +x run-docker-d.sh
sudo docker run \
  --name python-learning --detach \
  --mount src=`pwd`,target=/home/local,type=bind \
  -p 8088:8888 \
  --workdir /home/local \
  -it python:3.12.2 bash

# ...
# docker exec -it python-learning bash
# docker stop python-learning
# docker start python-learning
# docker rm python-learning
```
```sh
#!/bin/bash

# chmod +x run-docker.sh
sudo docker run \
  --name python-learning --rm \
  --mount src=`pwd`,target=/home/local,type=bind \
  -p 8088:8888 \
  --workdir /home/local \
  -it python:3.12.2 bash
```

----
<br/>
<!-- CentOS - Jupyter -->
</details>

### Others

<details markdown="block"><summary>CentOS - Others</summary>

```sh
#!/bin/bash

# chmod +x run-docker-ttyd.sh
sudo docker run -d --name ttyd \
  --restart unless-stopped \
  -p 8083:8083 \
  -e TZ="America/Sao_Paulo" \
  --mount src="$(pwd)/workspace/ttyd",target=/root/workspace,type=bind \
  --privileged \
  -it tsl0922/ttyd:1.7.3 ttyd -p 8083 -c user:password bash

# docker exec -it --privileged ttyd bash
# docker exec -d --privileged ttyd dockerd
```

```sh
#!/bin/bash

# chmod +x exec-docker-ttyd.sh
# sudo docker run -d --name ttyd \
#   --restart unless-stopped \
#   -p 8083:8083 \
#   -e TZ="America/Sao_Paulo" \
#   --mount src="$(pwd)/workspace/ttyd",target=/root/workspace,type=bind \
#   --privileged \
#   -it tsl0922/ttyd:1.7.3 ttyd -p 8083 -c user:password bash

sudo docker exec -it --privileged ttyd bash
# docker exec -d --privileged ttyd dockerd
```

```sh
#!/bin/bash

# chmod +x run-docker-sftpgo.sh
sudo docker run --rm --name my-sftpgo \
   -p 8080:8080 \
   -e TZ="America/Sao_Paulo" \
   -e SFTPGO_DATA_PROVIDER__CREATE_DEFAULT_ADMIN=true \
   -e SFTPGO_DEFAULT_ADMIN_USERNAME=user \
   -e SFTPGO_DEFAULT_ADMIN_PASSWORD=password \
   --mount type=bind,source="$(pwd)/workspace/ttyd",target=/srv/sftpgo \
   -d drakkan/sftpgo:2.5.x

# docker exec -it my-sftpgo bash
```

![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/deaa0b2c-2451-48f2-87ab-064f3f7231cc)


----
<br/>
<!-- CentOS - Others -->
</details>

----
<br/>
<!-- CentOS -->
</details>


## Google Cloud CLI

<details markdown="block">
  <summary>
    gcloud
  </summary>

```bash
# Google Cloud CLI Docker
# https://cloud.google.com/sdk/docs/downloads-docker
#
# Artifact Registry
# https://gcr.io/google.com/cloudsdktool/google-cloud-cli
docker pull gcr.io/google.com/cloudsdktool/google-cloud-cli:468.0.0

# $> mkdir -p ~/Downloads/temp
docker run \
  --name google-cloud-cli --rm \
  --mount src=`realpath ~/Downloads/temp`,target=/home/local,type=bind \
  --workdir /home/local \
  -it \
  gcr.io/google.com/cloudsdktool/google-cloud-cli:468.0.0 \
  bash

# https://github.com/sigoden/aichat/blob/601288029dd7affa2115547a70c74b21d2003b66/config.example.yaml#L83
# https://igorlima.github.io/unapologetic-snippets/docs/languages/shell/cli-ai#aichat
# https://igorlima.github.io/unapologetic-snippets/docs/languages/containerization/docker-samples-ai#aichat
gcloud auth application-default login
```

----
<br/>
<!-- gcloud -->
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
