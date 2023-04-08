# tpl

[![build](https://github.com/schneidexe/tpl/actions/workflows/release.yml/badge.svg?branch=master)](https://github.com/schneidexe/tpl/actions/workflows/release.yml)
![GitHub](https://img.shields.io/github/schneidexe/tpl)
![GitHub Release Date](https://img.shields.io/github/release-date/schneidexe/tpl)

tpl is build for generating config files from templates using simple or complex (lists, maps, objects) shell environment
variables. Since the binary has zero dependencies it is build for docker but you can use it across all platform and
operating systems.

tpl uses [sprig](https://github.com/Masterminds/sprig) to extend golang's template capabilities.

Check the test section and have a look at `test.tpl` (template) and `text.txt` (result) in test folder for examples.

## setup

Just download the binary for your OS and arch from the [releases](https://github.com/schneidexe/tpl/releases) page.

If you want to use it inside your docker image you can add this to your `Dockerfile`:

```
ADD https://github.com/schneidexe/tpl/releases/download/v/tpl-linux-amd64 /bin/tpl
RUN chmod a+x /bin/tpl
```

## build

Local:
```
go install github.com/schneidexe/tpl@latest
```

X-Platform:
```
go install github.com/mitchellh/gox@latest
gox -osarch="darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm64 windows/386 windows/amd64"
```

## test
```
./test.sh
```
