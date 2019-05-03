+++
Title: Bad Gateway Problems
Category: 1
Hero: 
Publish:  Jan 1, 2013 6:27pm (+0000) 
+++

There appear to be some major issues with the node application that runs this website. It crashes after some time, without a visible stack trace or anything. I'm not sure why this happens - the application is running under a daemon that should restart it if it dies, but this appears also not to work.

Either way, the golang application thats running on the same server is running just fine after several months. More investigation is needed, but I can't help but think that a better language wouldn't have these problems.
