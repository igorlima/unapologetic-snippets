---
layout: default
title: mosh
nav_order: 9
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/mosh
---

# mosh

_basic things_
```bash
mosh opc@oracle-learn             # 1) learn.ilima.xyz
mosh opc@oracle-docker            # 2) docker.ilima.xyz

# upload: local -> remote
scp ~/Downloads/file.txt opc@oracle-learn:/home/opc/file.txt
# download: remote -> local
scp opc@oracle-learn:/home/opc/file.txt ~/Downloads/
```

_when an IDE instance die/stop/shutdown_
```bash
# start Docker
sudo systemctl start docker

# stop Docker
sudo systemctl stop docker
sudo systemctl stop docker.socket

# give it a go
sudo docker run hello-world
```

## Synchronize a local directory with a remote machine utilizing Secure Shell (SSH) protocol.

```bash
# synchronize folder to a remote machine using ssh

# sync the remote destionation to a local destination
rsync -urltv -e ssh opc@oracle-learn:/home/opc/tmp-20240309-2236/ `realpath ~/Downloads/tmp`

# from local destination to a remote destination
rsync -urltv -e ssh `realpath ~/Downloads/tmp` opc@oracle-learn:/home/opc/tmp-20240309-2236/
```

__Note__: the `rsync` options used here are:
- `-u`, `--update` skip files that are newer on the receiver
- `-r`, `--recursive` recurse into directories
- `-l`, `--links` copy symlinks as symlinks
- `-t`, `--times` preserve modification times
- `-v`, `--verbose` increase verbosity
- `--delete` delete extraneous files from dest dirs, acts as `--delete-during`
- `-a` option means _"archive mode"_ which can be used to sync directories, transfer files, and preserve ownership and permissions.
- lastly, `-e` is the option that allows you to specify your remote shell, in this case ssh
