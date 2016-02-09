# tpl
Lean go templates for command line. 

See test section and have a look at `test.tpl` (template) and `text.txt` (result) for examples.

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
export foobar="{foo:bar,bar:foo,sna:fu}"
export foobaz="{foo:[sna,fu]}" 
export baz="1.0-123"

tpl -t test.tpl | diff - test.txt && echo Tests succeeded! || echo Tests failed!
```
