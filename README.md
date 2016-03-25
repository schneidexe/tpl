# tpl

tpl is build for genrating config files from templates using simple or complex (lists, maps, objects) shell environment variables. It's build for docker since the binary has zero dependecies but you can use it across all platform and opertating systems.

See test section and have a look at `test.tpl` (template) and `text.txt` (result) for examples.

## setup

Just download the binary for your OS and arch from the [releases](https://github.com/schneidexe/tpl/releases) page. 

If you want to use it inside your docker image you can add this to your `Dockerfile`:

```
RUN curl -sL https://github.com/schneidexe/tpl/releases/download/v0.3/tpl-linux-amd64 -o tpl && \
    chmod a+x tpl
```

## build 
```
docker build -t tpl . && \
docker run --name tpl tpl && \
docker cp tpl:/go/src/github.com/schneidexe/tpl/bin . && \
docker rm tpl && \
docker rmi tpl
```

## test
```
export foo="bar"
export bar="[foo,bar]"
export foobar="{foo:bar,bar:foo}"
export foobaz="{foo:[bar,baz]}" 
export baz="1.0-123"

tpl -t test.tpl | diff - test.txt && echo Tests succeeded! || echo Tests failed!
```
