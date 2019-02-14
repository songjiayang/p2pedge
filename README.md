## p2pedge

Edge compute with p2p (peer to peer) network, base on IPFS & IPFS pubsub.

![edage](https://user-images.githubusercontent.com/1459834/52759500-0408af00-3047-11e9-89c7-4ec774f69c4e.png)

### Why p2p edge compute?

- With the rise of the Internet of Things, the emergence of 5G networks will generate more data and more connections. The IDC construction of traditional cloud solutions is time-consuming and labor-intensive, and it is impossible to achieve large-area coverage, so that a large amount of data is transmitted over long distances. There will be a large delay in the process.

- P2p networks have a large number (smart hardware, sensors, personal computers, etc.). Most of them have low resource utilization. If we can use some way to integrate these resources, it can solve the geographical coverage in edge computing not enough problems, and it has a very big economic advantage.

### Requirements

- [ipfs](https://github.com/ipfs/go-ipfs/releases)
- p2pedge [tools](https://github.com/songjiayang/p2pedge/releases)
- p2pedge [examples](https://github.com/songjiayang/p2pedge/releases/download/v0.1.0/examples.tar.gz)

If everything is ok the folder will look like:

```
drwxr-xr-x   5 songjiayang  1603212982      160 Feb 14 15:16 .
drwx------+ 19 songjiayang  1603212982      608 Feb 14 15:00 ..
-rwxr-xr-x   1 songjiayang  1603212982  6709700 Feb 14 14:57 edge-ctl
-rwxr-xr-x   1 songjiayang  1603212982  6951396 Feb 14 14:54 edge-manager
drwxr-xr-x   4 songjiayang  1603212982      128 Feb 14 11:44 examples
```

### Usage

- start ipfs enable pubsub

```
ipfs daemon --enable-namesys-pubsub 
```

- start edge-manager

```
./edge-manager 
```

- add task's data to ipfs

```
ipfs add -r examples

added QmXtTrmq1cMWvvBRHFj46D6sMpzud2zaZw8EBLMDNW31jW examples/echo/task.json
added QmQtNxAtDZe1RJcmwngQBPdGMnAyufpxPUfWUXRPqzJ4wS examples/echo
added QmY7Ruumkyyg1NkbXJthnfhJ3jPfoDCfHDGNiKmMHcQ5Jk examples
 6.22 MiB / 6.22 MiB [==========================================
```

- add `echo` task to edge node

```
./edge-ctl add QmWeAxGEuZzmBWP2okuSgzKUiPH4yyMVB4otdSbzxUYBsq QmXtTrmq1cMWvvBRHFj46D6sMpzud2zaZw8EBLMDNW31jW

2019/02/14 15:22:42 Add successful, you can use `edge-ctl data QmXtTrmq1cMWvvBRHFj46D6sMpzud2zaZw8EBLMDNW31jW` to get the result.
```

Afer add, you can get print info at edge-manager console:

```
2019/02/14 15:22:42 New task with cid: QmXtTrmq1cMWvvBRHFj46D6sMpzud2zaZw8EBLMDNW31jW 
```

>> `QmWeAxGEuZzmBWP2okuSgzKUiPH4yyMVB4otdSbzxUYBsq` is ipfs node id, `QmXtTrmq1cMWvvBRHFj46D6sMpzud2zaZw8EBLMDNW31jW` is task's CID.

- get task's result

```
./edge-ctl data QmXtTrmq1cMWvvBRHFj46D6sMpzud2zaZw8EBLMDNW31jW
echo
echo
```