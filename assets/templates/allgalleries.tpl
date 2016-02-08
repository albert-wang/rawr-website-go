{{ define "head" }} 
	<title>Rawr Productions - Galleries</title>

	<link rel="stylesheet" href="/static/css/gallery.css">
{{ end }} 

{{ define "body" }}
	{{ range .Galleries }} 
		<div class="row title-row {{ if .Hero }}custom-hero{{else}}no-hero{{end}}" {{ if .Hero }}style="background-image: {{ darkenimg .Hero }};"{{end}}>
			<div class="col-xs-12 col-sm-6 col-sm-offset-3 gallery">
				<a href="/gallery/{{ .ID }}/{{ slug .Name }}">
					<h1>{{ .Name }}</h1>
				</a>
			</div>
		</div>

		<div class="row">
			<div class="col-xs-12 col-sm-6 col-sm-offset-3">
				<div class="postcontent">
				
				</div>
			</div>
		</div>
	{{ end }}
{{ end }}