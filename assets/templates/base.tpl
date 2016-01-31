<!DOCTYPE html>
<html>
	<head>
		<link rel="icon" type="image/png" href="/favicon-16x16.png" sizes="16x16"/>
		<link rel="icon" type="image/png" href="/favicon-32x32.png" sizes="32x32"/>
		<link rel="icon" type="image/png" href="/favicon-96x96.png" sizes="96x96"/>
		<link rel="icon" type="image/png" href="/favicon-192x192.png" sizes="192x192"/>

		<meta name="msapplication-TileColor" content="#343d5b">
		<meta name="msapplication-square310x310logo" content="/mstile-310x310.png"> 

		<script type="text/javascript" async="" src="https://ssl.google-analytics.com/ga.js"></script>
		<script type="text/javascript">
			var _gaq = _gaq || [];
			_gaq.push(['_setAccount', 'UA-28905432-1']);
			_gaq.push(['_trackPageview']);

			(function() {
			var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
			ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
			var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
			})();
		</script>

		<link rel="stylesheet" type="text/css" href="/static/lib/bootstrap.css">
		<link rel="stylesheet" type="text/css" href="/static/css/core.css">
		{{ template "head" . }}
	</head>

	<body>
		<div class="header">
			<div class="container">
				<img class="pull-right" src="/static/img/rawrproductions_light.png"/>
			</div>
		</div>

		<div class="background">
			<img src="{{ s3img "web/light.png" }}"/>
		</div>

		<div class="container">
			<img class="pull-right" src="/static/img/indie_light.png"/>

			<!-- Navigation -->
			<div class="row navigation">
				<ul>
					<li>
						<h1>Home</h1>
						<h6>the frontpage</h6>
					</li>

					<li>
						<h1>Blog</h1>
						<h6>read the blog</h6>
					</li>

					<li>
						<h1>Gallery</h1>
						<h6>browse the gallery</h6>
					</li>

					<li>
						<h1>Projects</h1>
						<h6>what i'm workign on</h6>
					</li>

					<li>
						<h1>About</h1>
						<h6>about rawr productions</h6>
					</li>
				</ul>
			</div>

			{{ template "body" . }} 
		</div>

	
		
		
		
	</body>
</html>