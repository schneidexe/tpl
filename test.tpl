empty value default:
{{ .empty }}

empty value with check:
{{ if .empty }}{{ .empty }}{{ end }}

simple string:
{{ .foo }}

null value:
{{ .null }}

special characters:
{{ .baz }}
{{ .money }}

number:
{{ .number }}

list:
{{ .bar }}

iterate over list:
{{ range $bar := .bar}}{{ . }}
{{end}}

iterate over list with index:
{{ range $idx, $bar := .bar}}{{ $idx }}:{{ $bar }}
{{end}}

access element in list:
{{ index .bar 1 }}

map:
{{ .foobar }}

iterate over map with key and value:
{{ range $foo, $bar := .foobar}}{{ $foo }}:{{ $bar }}
{{end}}

access element in map:
{{ .foobar.foo }}

access subelements:
{{ index .foobaz.foo 1 }}
