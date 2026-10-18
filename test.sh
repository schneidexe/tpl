#!/usr/bin/env bash
export tpl_test_foo="bar"
export tpl_test_bar="[foo,bar]"
export tpl_test_foobar="{foo:bar,bar:foo}"
export tpl_test_foobaz="{foo:[bar,baz]}"
export tpl_test_baz="1.0-123"
export tpl_test_number="59614658972"
export tpl_test_null="null"
export tpl_test_empty=
export tpl_test_money="500€"
export tpl_test_special="?&>=:/"
export tpl_test_woot="[]"
export tpl_test_whoa="{}"
export tpl_test_backslash="\.\/"
export tpl_test_urls="{google:[https:://google.com,http:://google.de],github:https:://github.com}"
export tpl_test_json='{"abc":123,"def":["a","b","c"],"ghi":"[{,!?!,}]"}'

env | grep tpl_test

echo
echo

go run tpl.go -t test.tpl | diff -y test.txt - && echo Tests succeeded! || echo Tests failed!

echo
echo

go run tpl.go -t test.tpl -o test.out && diff -y test.txt test.out && echo Tests succeeded! || echo Tests failed!
