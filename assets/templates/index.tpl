{{ define "head" }}
	<title>Rawr Productions</title>
{{ end }}

{{ define "header" }}
{{ end }}

{{ define "body" }}
	<div class="row feature">
		<div class="col-xs-12 col-sm-6 col">
			<h6 class="mini-header">featured</h6>
			<h1>{{ .Featured.Title }}</h1>
			<div class="postcontent">
				{{ blackfriday .Featured.Short }}
			</div>
			<a class="more" href="/post/{{ .Featured.ID }}/{{ slug .Featured.Title }}">More...</a>
		</div>
	</div>

	<div class="row message">
		<div class="col-xs-12 col-sm-7">
			<h1>Want to see what I've been working on recently?</h1>
			<p>
				Check out the <a href="/projects">projects</a> page or see 
				what I have open sourced on <a href="https://bitbucket.org/rraawwrr/">Bitbucket</a>
			</p>
		</div>
	</div>

	<div class="row previews">
		{{ range $ind, $post := .Posts }}
			<div class="col-xs-12 col-sm-6 col-md-3">
				<div class="no{{ add $ind 2 }}">
					<a href="/post/{{ $post.ID }}/{{ slug $post.Title }}">
						<h1>{{ $post.Title }}</h1>
					</a>
					<div class="postcontent">
						{{ blackfriday $post.Short }}
					</div>
				</div>
			</div>
		{{ end }}
	</div>
{{ end }} 