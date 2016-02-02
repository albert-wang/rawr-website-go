{{ define "head" }}
	<link rel="stylesheet" href="/static/lib/simplemde.css">
	<script src="https://code.jquery.com/jquery-2.2.0.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/underscore.js/1.8.3/underscore-min.js"></script>
	<script src="/static/lib/simplemde.js"></script>

	<link rel="stylesheet" href="/static/css/post.css">
	<link rel="stylesheet" href="/static/lib/monokai-sublime.css">
	<script src="/static/lib/highlight.pack.js"></script>

	<title>Editing</title>

	<script>
	$(document).ready(function() {
		var contents = {{ .Post.Content }}

		var renderPreview = function(plainText, preview) { // Async method
			$.ajax("/admin/render", {
				"type" : "POST", 
				"data" : plainText, 
				"processData" : false,
				"success" : function(r) {
					currentPreview = r;
					$(preview).removeClass("postcontent post").addClass("postcontent post");
					preview.innerHTML = r;

					$('pre code').each(function(i, block) {
						hljs.highlightBlock(block);
					});
				}
			});

			return currentPreview;
		};

		renderPreview = _.debounce(renderPreview, 300);

		var currentPreview = "Loading...";
		var simplemde = new SimpleMDE({
			codeSyntaxHighlighting: true,
			blockStyles: {
				bold: "*",
				italic: "_"
			},
			initialValue: contents,
			previewRender: renderPreview,
		});


		$("#submit").on("click", function() {
			$.ajax("/admin/edit", {
				"type" : "POST", 
				"data" : JSON.stringify({
					"id" : {{ .Post.ID }}, 
					"content" : simplemde.value(),
					"title" : $("#title").val(), 
					"publish" : parseInt($("#publish").val()),
					"category" : parseInt($("#category").val()),
				}),
				"processData" : false
			});
		});
	});
	</script>
{{ end }}


{{ define "body" }}
<div class="row">
	<div class="col-xs-12">
		<textarea></textarea>
		<div class="row">
			<div class="col-xs-2">
				<span>Publish</span>
			</div>

			<div class="col-xs-9">
				<span>Title</span>
			</div>
		</div>
		<div class="row">
			<div class="col-xs-2">
				<input type="number" class="form-control" id="publish" placeholder="Publish Date" value="{{ if .Post.Publish }}{{ .Post.Publish.Unix }}{{ else }}0{{ end }}"/>
			</div>

			<div class="col-xs-6">
				<input type="text" class="form-control" id="title" placeholder="Title" value="{{ .Post.Title }}"/>
			</div>

			<div class="col-xs-3">
				<select class="form-control" id="category">
				{{ range .Categories }} 
					<option value={{.ID}} {{ if eq $.Post.CategoryID .ID }} selected {{ end }}>{{.Category}}</option>
				{{ end }}
					
				</select>
			</div>

			<div class="col-xs-1">
				<button id="submit" type="submit" class="pull-right btn btn-default">Submit</button>
			</div>
		</div>
	</div>
</div>
{{ end }}