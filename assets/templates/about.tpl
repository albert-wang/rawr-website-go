{{ define "head" }}
	<title>Rawr Productions - About</title>
{{ end }}


{{ define "body" }}
	<div class="row feature">
		<div class="col-xs-12 col-sm-6 col">
			<h6 class="mini-header">about</h6>
			<h1>Rawr Productions</h1>
			<div class="postcontent">
				Rawr Productions is a one-man game-programming endevor. 

				I write games in my spare time, and talk about programming the rest
				of the time. I'm courrently working on a game called Araboth, a 
				high-intensity SHMUP modeled after Ikaruga.
			</div>
		</div>
	</div>

	<div class="row message"></div>

	<div class="row previews">
		<div class="col-xs-12 col-sm-6 col-md-3">
			<div class="no5">
				<h1>Programming</h1>
				<div class="postcontent">
					<p>
						I mostly work in C++ and DirectX11.
						Webdev stuff is done in golang, less, typescript
						and assorted tech as I see fit.
					</p>
					<p>
						This server is hosted in a docker instance, 
						backed up by postgres and cached by Redis, monitored
						by StatsD.
					</p>
				</div>
			</div>
		</div>

		<div class="col-xs-12 col-sm-6 col-md-3">
			<div class="no2">
				<h1>Gaming</h1>
				<div class="postcontent">
					<p>
						I'm a high-level Magic: The Gathering player, currently
						specializing in modern and legacy, running some sort of URx control
						and Miracles, respectively.
					</p>

					<p>
						I am also an avid player of fighting games. Mostly, I play
						Guilty Gear Xrd and Blazblue, but also dabble in Street Fighter
						and more obscure games, like UNIEL.
					</p>
				</div>
			</div>
		</div>

		<div class="col-xs-12 col-sm-6 col-md-3">
			<div class="no3">
				<h1>Education</h1>
				<div class="postcontent">
					<p>
						Graduated with a Bachelors of Science from the University of 
						Texas at Austin in Computer Science and a Batchelors of Business Administration
						in Finance from the same.
					</p>
				</div>
			</div>
		</div>

		<div class="col-xs-12 col-sm-6 col-md-3">
			<div class="no4">
				<h1>Social</h1>
				<div class="postcontent">
					<p>
						You can find my twitter feed <a href="https://twitter.com/rraawwrr">@rraawwrr</a>, and my twitch.tv at <a href="https://twitch.tv/rrrrawr">rrrrawr</a>
					</p>
				</div>
			</div>
		</div>
	</div>
{{ end }}