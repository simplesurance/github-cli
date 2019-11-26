# github-cli

github-cli is a commandline tool to interact with the Github API.

The tool only supports a few commands that we use in our CI system.

## Usage

see:

```sh
github-cli --help
```

## Installation

### From a Release

- Download a release from <https://github.com/simplesurance/github-cli/releases>
- Extract the .tar.xz archive via `tar xJf <filename>` and copy it into your `$PATH`

### Run with Docker

```shell
docker run -v $PWD:/repo -w /repo simplesurance/github-cli:latest
```

### Via go get

```shell
go get github.com:simplesurance/github-cli
```
