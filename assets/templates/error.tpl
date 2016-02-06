{{ define "head" }}
	<title>Rawr Productions - {{ .Code }}</title>
	<link rel="stylesheet" href="/static/css/error.css">
{{ end }}


{{ define "body" }}
	<h1>{{ .Code }}</h1>
	<img src="/static/img/penguindrum.png"/>
	<h2>{{ .Message }}</h2>
{{ end }}