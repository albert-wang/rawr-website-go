{{ define "head" }} 
	<title>Rawr Productions - {{ .Category }}</title>

	<link rel="stylesheet" href="/static/css/archive.css">
{{ end }} 

{{ define "body" }}
	{{ range .Posts }} 
		<div class="row title-row {{ if .Hero }}custom-hero{{else}}no-hero{{end}}" {{ if .Hero }}style="background-image: {{ darkenimg .Hero }};"{{end}}>
			<div class="col-xs-12 col-sm-6 col-sm-offset-3">
				<h6><a href="/blog/{{ .Category }}/1">{{.Category}}</a></h6>
				<a href="/post/{{ .ID }}/{{ slug .Title }}">
					<h1>{{ .Title }}</h1>
				</a>
			</div>
		</div>

		<div class="row">
			<div class="col-xs-12 col-sm-6 col-sm-offset-3">
				<div class="postcontent">
					{{ blackfriday .Short }}
				</div>
			</div>
		</div>
	{{ end }}
{{ end }}