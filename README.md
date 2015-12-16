# tpl
Lean go templates for command line

## build 
```
docker build -t tpl .
docker run --name tpl tpl && docker cp tpl:/go/src/github.com/schneidexe/tpl/bin . && docker rm tpl
```

## test
```
export foo="bar"
export bar="[foo,bar]"
export foobar="{foo:bar,bar:foo,sna:fu}"
export snafu="{foo:[sna,fu]}" 

tpl -t test.tpl
```