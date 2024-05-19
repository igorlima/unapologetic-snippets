---
layout: default
title: termux
nav_order: 9
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/termux
---

# Termux

## Uploading and Downloading Files to/from a Remote Server

```sh
# upload: local -> remote
cp storage/downloads/my-clippings-20240519.txt workstation
scp ~/workstation/my-clippings-20240519.txt oracle-learn:/home/opc/
sudo docker cp my-clippings-20240519.txt ttyd:/root/workspace/

# archive file(s)
tar -cf 20240519-1433.tar -T /dev/null              # create an empty tar archive
tar -tf 20240519-1433.tar                           # show the archive file
tar -rf 20240519-1433.tar output/clips.json         # add file(s) to an existing archive
tar -uf 20240519-1433.tar output/clips.json         # update file(s) to an existing archive
tar --delete -f 20240519-1433.tar output/clips.json # delete file(s) to an existing archive
tar -xf 20240519-1433.tar -C ./tmp                  # extract the content of an archive file into specific folder
tar -czvf 20240519-1433.tar.gz 20240519-1433.tar    # compress an existing archive
tar -xzvf 20240519-1433.tar.gz -C ./tmp             # uncompress an existing archive

# download: remote -> local
sudo docker cp ttyd:/root/workspace/20240519-1433.tar .
scp oracle-learn:/home/opc/20240519-1433.tar ~/workstation/
cp workstation/20240519-1433.tar storage/downloads/
```


----

[^1]: [...](...)
