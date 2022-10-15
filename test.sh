#!/usr/bin/env bash
export foo="bar"
export bar="[foo,bar]"
export foobar="{foo:bar,bar:foo}"
export foobaz="{foo:[bar,baz]}"
export baz="1.0-123"
export number="59614658972"
export null="null"
export empty=
export money="500€"
export special="?&>=:/"
export woot="[]"
export whoa="{}"
export backslash="\.\/"
export urls="{google:[https:://google.com,http:://google.de],github:https:://github.com}"
export json='{"abc":123,"def":["a","b","c"],"ghi":"[{,!?!,}]"}'

go get -v

echo
echo

go run tpl.go -t test.tpl | diff -y test.txt - && echo Tests succeeded! || echo Tests failed!

echo
echo

go run tpl.go -t test.tpl -o test.out && diff -y test.txt test.out && echo Tests succeeded! || echo Tests failed!
