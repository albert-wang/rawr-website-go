{{ define "head" }}
	<title>Rawr Productions - {{ .Post.Title }}</title>
	<link rel="stylesheet" type="text/css" href="/static/css/post.css">
	<link rel="stylesheet" href="/static/lib/monokai-sublime.css">

	<script src="https://code.jquery.com/jquery-2.2.0.min.js"></script>
	<script src="/static/lib/highlight.pack.js"></script>

	<script>
		$(document).ready(function() {
			$('pre code').each(function(i, block) {
				hljs.highlightBlock(block);
			});

			$('.postcontent p').filter(function() {
				return this.innerHTML == "";
			}).remove();
		});
	</script>

	{{ if .Post.Hero }}
	<style>
		.feature.custom-hero {
			background-image: {{ darkenimg .Post.Hero }};
		}
	</style>
	{{ end }}
{{ end }}

{{ define "body" }}
	<div class="row feature {{ if .Post.Hero }}custom-hero{{else}}no-hero{{end}}" data-hero='{{ .Post.Hero }}'>
		<div class="col-xs-8 col-xs-offset-2">
			<h6 class="mini-header">{{ .Post.Category }}</h6>
			<h1>{{ .Post.Title }}</h1>
		</div>
	</div>

	<div class="row post">
		<div class="col-xs-8 col-xs-offset-2">
			<div class="postcontent">
				{{ blackfriday .Post.Full }}
			</div>
		</div>
	</div>
{{ end }} 
