undefined value default:
_{{ .undefined }}

simple string:
_{{ .foo }}

null value:
_{{ .null }}

special characters:
_{{ .baz }}
_{{ .money }}

number:
_{{ .number }}

list:
_{{ .bar }}

iterate over list:
{{ range .bar}}_{{ . }}
{{end}}

iterate over undefined list:
{{ range .undefined}}_{{ . }}
{{end}}

iterate over empty list:
{{ range .woot}}_{{ . }}
{{end}}

iterate over list with index:
{{ range $idx, $bar := .bar}}_{{ $idx }}:{{ $bar }}
{{end}}

access element in list:
_{{ index .bar 1 }}

map:
_{{ .foobar }}

empty map:
_{{ .whoa }}

iterate over map with key and value:
{{ range $foo, $bar := .foobar}}_{{ $foo }}:{{ $bar }}
{{end}}

access element in map:
_{{ .foobar.foo }}

access subelements:
_{{ index .foobaz.foo 1 }}

sprig:
_{{ .foo | upper }}
_{{ .bar | first }}