# tpl

tpl is build for generating config files from templates using simple or complex (lists, maps, objects) shell environment 
variables. Since the binary has zero dependencies it is build for docker but you can use it across all platform and 
operating systems.

tpl uses [sprig](https://github.com/Masterminds/sprig) to extend golang's template capabilities.

Check the test section and have a look at `test.tpl` (template) and `text.txt` (result) in test folder for examples.

## setup

Just download the binary for your OS and arch from the [releases](https://github.com/schneidexe/tpl/releases) page. 

If you want to use it inside your docker image you can add this to your `Dockerfile`:

```
ADD https://github.com/schneidexe/tpl/releases/download/v0.6.0/tpl-linux-amd64 /bin/tpl
RUN chmod a+x /bin/tpl
```

## build 

Local:
```
go get github.com/schneidexe/tpl
```

X-Platform:
```
go get github.com/mitchellh/gox
gox
```

## test
```
./test.sh
```
