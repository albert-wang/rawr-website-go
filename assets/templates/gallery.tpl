{{ define "head" }} 
	<title>Rawr Productions - Galleries</title>

	<link rel="stylesheet" href="/static/css/gallery.css">
{{ end }} 

{{ define "body" }}
<div class="row">
	<div class="col-xs-12">
		<h1>{{ .Gallery.Name }}</h1> <h6> {{ .Gallery.Description }} </h6>
	</div>
</div>

<div class="row">
	{{ range .Images }} 
		<div class="col-xs-4">
			<a href="{{s3img .Orig }}">
				<img src="{{s3img .Thumb }}"/>
			</a>
		</div>
	{{ end }}
</div>
{{ end }}