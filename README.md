# tpl

tpl is built for generating config files from templates using simple or complex (lists, maps, objects) shell environment
variables. Since the binary has zero dependencies, it is built for Docker, but you can use it across all platforms and
operating systems.

tpl uses [sprig](https://github.com/Masterminds/sprig) to extend Go's template capabilities.

Check the test section and have a look at `test.tpl` (template) and `test.txt` (result) in the project directory for examples.

## Setup

Just download the binary for your OS and architecture from the [releases](https://github.com/schneidexe/tpl/releases) page.

If you want to use it inside your Docker image, you can add this to your `Dockerfile`:

```
ADD https://github.com/schneidexe/tpl/releases/download/v0.7.6/tpl-linux-amd64 /bin/tpl
RUN chmod a+x /bin/tpl
```

## Build

Local:
```
go build
```

Cross-platform:
```
GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o tpl-darwin-amd64 .
GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -o tpl-darwin-arm64 .
GOOS=linux   GOARCH=386   CGO_ENABLED=0 go build -o tpl-linux-386 .
GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o tpl-linux-amd64 .
GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -o tpl-linux-arm64 .
GOOS=windows GOARCH=386   CGO_ENABLED=0 go build -o tpl-windows-386.exe .
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o tpl-windows-amd64.exe .
```

## Test
```
./test.sh
```
