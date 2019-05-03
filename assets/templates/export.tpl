+++
Title: {{ .Title }}
Category: {{ .CategoryID }}
Hero: {{ .Hero }}
Publish: {{ if .Publish }} {{ .Publish.Format "Jan 2, 2006 3:04pm (MST)" }} {{ end }}
+++

{{ .Content }}
