undefined value:
_{{.undefined}}
_{{if .undefined}}{{.undefined}}{{end}}
_{{default "" .undefined}}

empty value:
_{{.tpl_test_empty}}

simple string:
_{{.tpl_test_foo}}

null value:
_{{.tpl_test_null}}

special characters:
_{{.tpl_test_baz}}
_{{.tpl_test_money}}
_{{.tpl_test_special}}

number:
_{{.tpl_test_number}}

list:
_{{.tpl_test_bar}}

iterate over list:
{{range .tpl_test_bar}}_{{.}}
{{end}}

iterate over undefined:
{{range .undefined}}_{{.}}
{{end}}

iterate over list with index:
{{range $idx, $bar := .tpl_test_bar}}_{{$idx}}:{{$bar}}
{{end}}

access element in list:
_{{index .tpl_test_bar 1}}

map:
_{{.tpl_test_foobar}}

empty map:
_{{.tpl_test_whoa}}

iterate over map with key and value:
{{range $foo, $bar := .tpl_test_foobar}}_{{$foo}}:{{$bar}}
{{end}}

access element in map:
_{{.tpl_test_foobar.foo}}

access subelements:
_{{index .tpl_test_foobaz.foo 1}}

sprig:
_{{.tpl_test_foo | upper}}
_{{.tpl_test_bar | first}}
_{{default "foo" .undefined}}
_{{default "foo" .tpl_test_empty}}

backslash:
_{{.tpl_test_backslash}}

urls:
_{{.tpl_test_urls}}

include:
_{{ include "foobar.txt" }}
_{{ mustInclude "test/include.txt" | trim }}

json:
_{{ .tpl_test_json }}
_{{ toJson .tpl_test_json }}
{
{{ range $idx, $key := (keys .tpl_test_json) }}
{{ if gt $idx 0 }},{{ end }}{{ quote $key }}:{{ get $.tpl_test_json $key | quote | upper }}
{{ end }}
}
