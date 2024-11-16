---
layout: default
title: ssh
nav_order: 9
parent: Shell
grand_parent: Programming Languages
permalink: /docs/languages/shell/ssh
---

# SSH

## file transfer - download and upload

- [mosh]({% link docs/languages/shell/mosh.md %})
  - `mosh user@oracle-server`
    ```sh
    # upload: local -> remote
    scp ~/Downloads/file.txt user@oracle-learn:/home/user/file.txt
    # download: remote -> local
    scp user@oracle-learn:/home/user/file.txt ~/Downloads/
    ```

<details markdown="block"> <summary>sftp</summary>

```sh
sftp user@oracle-server
```

```sh
# to get available commands
sftp> help

# change local and remote directory
# check your local and remote server directory
sftp> !pwd                      # check current directory on local system
sftp> pwd                       # check current directory on remote system
sftp> lcd /home/                # navigate between directories on local system
sftp> cd Uploads                # navigate between directories on remote system

# interactive sFTP commands
sftp> ls                        # list directory
sftp> pwd                       # print working directory on remote host
sftp> lpwd                      # print working directory on local host
sftp> mkdir uploads             # create a new directory

# to upload files and folders recursively
# use the flag -r
sftp> put -r  Tecmint.com-articles
sftp> mkdir Tecmint.com-articles
sftp> put -r Tecmint.com-articles

# to preserve the modification times, access times, and modes from the original files transferred
# use the -p flag.
sftp> put localfile.txt
sftp> put -r localdir
sftp> put -pr Tecmint.com-articles

# to download a whole directory from remote Linux host to local machine
# use the get command with the -r flag
sftp> get remotefile.txt
sftp> get -r remotedir
sftp> get -r fstools-0.0

sftp> bye
sftp> exit
```

</details>

<details markdown="block"> <summary>scp</summary>

```sh
# upload: local -> remote
scp local_file user@remote_host:remote_file

# download: remote -> local
scp user@remote_host:remote_file local_file
```
</details>

<details markdown="block"> <summary>reverse tunnel port on ssh server</summary>

how to send data to local clipboard from a remote SSH session

```sh
# REFERENCE
# https://stackoverflow.com/questions/1152362/how-to-send-data-to-local-clipboard-from-a-remote-ssh-session
# https://gist.github.com/dergachev/8259104

# find the IP address of the client in an SSH session
$> who am i
$> echo $SSH_CLIENT
$> echo $SSH_CONNECTION

# client (ssh session startup)
$> ssh username@server.com -R 2000:localhost:2000
$> ssh user@oracle-server -R 2000:localhost:2000
# usage of -N for no interactive command line
$> ssh user@oracle-server -N -R 2000:localhost:2000

# client (another tab)
$> nc -l 2000 | pbcopy

# server (inside SSH session)
$> cat some_useful_content.txt | nc localhost 2000

# troubleshooting
$> netstat -anpt | grep 2000   # CentOS
$> lsof -i :2000               # OSX
```

- <details markdown="block"> <summary>pbcopy/pbpaste through scp</summary>

  ```sh
  # upload: local -> remote
  scp clipboard.txt user@oracle-server:clipboard.txt
  # download: remote -> local
  scp user@oracle-server:clipboard.txt clipboard.txt
  ```
  </details>
- <details markdown="block"> <summary>from local -> to remote</summary>

  ```sh
  # The `-L` option tells the tunnel to answer on the local side of the tunnel (the host running your client).
  # https://unix.stackexchange.com/questions/46235/how-does-reverse-ssh-tunneling-work

  # client (ssh session startup)
  $> ssh user@oracle-server -N -L 2000:localhost:2000

  # server (inside SSH session)
  $> nc -l 2000 | vim -
  $> nc -l 2000 | cat
  $> nc -l 2000 | cat > text.txt   # to create/overwrite
  $> nc -l 2000 | cat >> text.txt  # to append

  # client (local machine)
  # -c, --close      close connection on EOF from stdin
  $> echo "hello world" | nc localhost 2000 -c
  $> pbpaste | nc localhost 2000 -c
  $> cat some_useful_content.txt | nc localhost 2000 -c
  ```
  </details>
- <details markdown="block"> <summary>from remote -> to local</summary>

  ```sh
  # the `-R` option tells the tunnel to answer on the remote side (the SSH server).
  # https://unix.stackexchange.com/questions/46235/how-does-reverse-ssh-tunneling-work

  # client (ssh session startup)
  $> ssh user@oracle-server -N -R 2000:localhost:2000

  # client (another tab)
  $> nc -l 2000 | pbcopy

  # server (inside SSH session)
  $> echo "hello world" | nc localhost 2000
  $> cat some_useful_content.txt | nc localhost 2000
  ```
  </details>
- <details markdown="block"> <summary>troubleshooting</summary>

  ```sh
  # https://gist.github.com/jasperf/c189ae315719f1e5ddd5
  $> netstat -anpt | grep 2000   # CentOS
  $> lsof -i :2000               # OSX
  ```
  </details>

<!-- reverse tunnel port on ssh server -->
</details>
