# tpl
Lean go templates for command line

## build 
```
docker build -t tpl .
docker run --name tpl tpl && docker cp tpl:/go/src/github.com/schneidexe/tpl/bin . && docker rm tpl
```

## test
```
FOO=bar; BAR=foo,bar; FOBAR=foo:bar,bar:foo,sna:fu; bin/tpl-<OS>-<ARCH> -t test.tpl

```