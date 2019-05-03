+++
Title: Constant speed on a spline
Category: 1
Hero: 
Publish: 2013-05-27 23:37:00.935539 +0000 +0000
+++

Whenever one needs a curve, a spline of some sort is usually involved. The most common forms are 
Catmull-Rom splines and Cubic Bezier Splines. Both of these splines have relatively simple polynomial 
forms, and are easily evaluated in the vertex shader. {:truncate}{:longtruncate}

A more complicated issue is the question of how to move along such a curve at a constant speed. 
Anyone looking for such a process is likely to stumble the root finding method, specified in the 
paper “Moving Along a Curve with Specified Speed”, by David Eberly. The process is 
straightforward – given a length, solve for the time that the curve reaches that length through
a modified Newton’s method. This method involves repeated computations of length, which is a 
slow operation. Calculating more than a few dozen reparameterizations is not possible on high end 
hardware, so the problem becomes one of acceleration.

The first step is to simply generate better initial guesses for Newton’s method, so that the number 
of iterations is minimized. At a cost of 4 bytes per spline, storing the previous result and feeding 
it (plus a small constant) to the next attempt to find a root reduces the number of iterations by 
two or three. This saves several integrals, but is still too slow. A more aggressive caching policy 
involves a preprocessing step. Instead of solving for time, a single integral can generate lengths 
for every time step. This generates a set of unevenly distributed `length, time` pairs, which can 
be linearly interpolated to approximate a reparameterization. 

The number of samples directly correlates to the error in the approximation, as shown in the 
image below. The numbers shown in the graph are derived from many random splines. A larger version can be found [here](http://img.rawrrawr.com/gallery/bfdf3f253b0f32153b747ebe358def94a838328d.png)

![graph](http://img.rawrrawr.com/gallery/5ba60d3f5fc68aeb5416fcfa732105000c5ecb3d.png)

4 samples will keep the difference below 30%, and each additional doubling of the number of samples 
halves the error level. 16 samples has roughly 5% error, which is good enough for most graphical 
applications. 16 samples also happens to fit into 4 float4 registers in a shader, so the linear 
fit interpolation can be done in hardware without any additional CPU work beyond computing the initial fit. 