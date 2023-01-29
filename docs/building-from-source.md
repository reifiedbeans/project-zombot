# Building from source

Start by downloading the source code using Git.

```shell
git clone https://github.com/reifiedbeans/project-zombot
```

Then compile the binary using the [Makefile](/Makefile).

```shell
make
```

You can pass `GOOS` and `GOARCH` variables as well to change the target OS and architecture, respectively.

```shell
GOOS=linux GOARCH=amd64 make
```

Running the above commands will create a binary in the `bin` folder.
