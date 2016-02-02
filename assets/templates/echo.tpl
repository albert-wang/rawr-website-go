{{ define "head" }}
	<title>Rawr Productions</title>
{{ end }}

{{ define "header" }}
{{ end }}

{{ define "body" }}
	<div class="row feature">
		<div class="col-xs-12 col-sm-6 col">
			<h6 class="mini-header">featured</h6>
			<h1>Pipelining - Why It Matters</h1>
			<div class="postcontent">
			<p>
			While reading Hacker News, I came upon an article that purported that Java (or other high level languages) should be used in preference to C++ due to the quality of modern interpreters and JIT compilers. I will not discuss the conclusion, which is shaky at best, but rather the microbenchmark that was used to display the quality of the Java virtual machine.
			</p>
			</div>
			<span class="more">More...</span>
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
        <div class="col-xs-12 col-sm-6 col-md-3">
        	<div class="no2">
                <a href="/post/13/constant-speed-on-a-spline">
                    <h1>Constant Speed On A Spline</h1>
                </a>
                <div class="postcontent"> <p>Whenever one needs a curve, a spline of some sort is usually involved. The most common forms are 
Catmull-Rom splines and Cubic Bezier Splines. Both of these splines have relatively simple polynomial 
forms, and are easily evaluated in the vertex shader. </p> </div>
			</div>
        </div>
        
        <div class="col-xs-12 col-sm-6 col-md-3">
        	<div class="no3">
                <a href="/post/12/bad-gateway-problems">
                    <h1>Bad Gateway Problems</h1>
                </a>
                <div class="postcontent"> <p>There appear to be some major issues with the node application that runs this website. It crashes after some time, without a visible stack trace or anything. I'm not sure why this happens - the application is running under a daemon that should restart it if it dies, but this appears also not to work.</p>

				<p>Either way, the golang application thats running on the same server is running just fine after several months. More investigation is needed, but I can't help but think that a better language wouldn't have these problems.</p> </div>
			</div>
        </div>
        
        <div class="col-xs-12 col-sm-6 col-md-3">
        	<div class="no4">
                <a href="/post/11/pipelining---why-it-matters">
                    <h1>Pipelining - Why It Matters</h1>
                </a>
                <div class="postcontent"> <p>While reading Hacker News, I came upon an article that purported that Java (or other high level languages) should be used in preference to C++ due to the quality of modern interpreters and JIT compilers. I will not discuss the conclusion, which is shaky at best, but rather the microbenchmark that was used to display the quality of the Java virtual machine.</p> </div>
			</div>
        </div>
        
        <div class="col-xs-12 col-sm-6 col-md-3">
        	<div class="no5">
                <a href="/post/10/luajit-and-luabind">
                    <h1>Luajit And Luabind</h1>
                </a>
                <div class="postcontent"> <p>Just a quick aside for anyone that may be using <a href="http://luajit.org/">luajit</a> with <a href="http://www.rasterbar.com/products/luabind.html">luabind</a> on x86 platforms - be careful about the way that luabind throws exceptions. Luajit on x86 platforms has no exception interopability, which means that throwing exceptions through lua frames is not supported, and may corrupt the stack. This usually is not a problem, since luabind generally only throws exceptions after an error, or some other confluence of events conspires so that the exception does not corrupt anything. </p> </div>
            </div>
        </div>
	</div>
{{ end }} 