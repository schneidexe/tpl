# tpl
Lean go templates for command line

## build 
```
docker build -t tpl .
docker run --name tpl tpl && docker cp tpl:/go/src/github.com/schneidexe/tpl/bin . && docker rm tpl
```

## test
```
foo="bar; bar="[foo,bar]"; foobar="{foo:bar,bar:foo,sna:fu}"; snafu="{foo:[sna,fu]}"; bin/tpl-<OS>-<ARCH> -d -t test.tpl

```