# unix2dos-go

This repository contains a simple cli written in go to convert DOS/Windows file newlines into Unix file newlines and vice-versa.

## Usage

```
unix2dos-go is a simple cli written in go to convert line breaks in a text file from Unix format (LF) to DOS format (CR+LF) and vice versa.
In DOS/Windows text files a line break, also known as newline, is a combination of two characters: a Carriage Return (CR) followed by a Line Feed (LF). In Unix text files a line break is a single character: the Line Feed
(LF). In Mac text files, prior to Mac OS X, a line break was single Carriage Return (CR) character. Nowadays Mac OS uses Unix style (LF) line breaks.

It is inspired by the unix2dos utility.


The source code is available at https://github.com/lescactus/unix2dos-go

Usage:
  unix2dos-go [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dos2unix    Convert DOS file format to Unix file format
  help        Help about any command
  unix2dos    Convert Unix file format to DOS file format

Flags:
  -h, --help            help for unix2dos-go
  -m, --mode uint32     Unix permission numeric mode in octal for the converted file. See chmod(1) for more informations. (default 420)
  -o, --output string   Name of the converted file. (default "./unix2dos.converted")
```

## Installation

### From source with go

You need a working [go](https://golang.org/doc/install) toolchain (It has been developped and tested with go 1.16 and go 1.16 only, but should work with go >= 1.12). Refer to the official documentation for more information (or from your Linux/Mac/Windows distribution documentation to install it from your favorite package manager).

```sh
# Clone this repository
git clone https://github.com/lescactus/unix2dos-go.git && cd unix2dos-go/

# Build from sources. Use the '-o' flag to change the compiled binary name
go build

# Default compiled binary is unix2dos-go
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./unix2dos-go --help
```

### From source with docker

If you don't have [go](https://golang.org/) installed but have docker, run the following command to build inside a docker container:

```sh
# Build from sources inside a docker container. Use the '-o' flag to change the compiled binary name
# Warning: the compiled binary belongs to root:root
docker run --rm -it -v "$PWD":/app -w /app golang:1.16 go build

# Default compiled binary is unix2dos-go
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./unix2dos-go --help
```

### From source with docker but built inside a docker image

If you don't want to pollute your computer with another program, this cli comes with its own docker image:

```sh
docker build -t unix2dos-go .

docker run --rm -v "$PWD":/tmp -w /tmp unix2dos-go unix2dos <file>
```
