simple string:
{{ .FOO }}

truncate string (only works for truncating the tail):
{{ .FOO | printf "%.2s" }}

list:
{{ .BAR }}

iterate over list:
{{ range $bar := .BAR}}{{ . }}
{{end}}

iterate over list with index:
{{ range $idx, $bar := .BAR}}{{ $idx }}:{{ $bar }}
{{end}}

access element in list:
{{ index .BAR 1 }}

map:
{{ .FOOBAR }}

iterate over map with key and value:
{{ range $foo, $bar := .FOOBAR}}{{ $foo }}:{{ $bar }}
{{end}}

access element in map:
{{ .FOOBAR.foo }}
