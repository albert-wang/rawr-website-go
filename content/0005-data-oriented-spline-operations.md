+++
Title: data oriented spline operations
Category: 3
Hero: 
Publish:  Feb 14, 2012 5:10am (+0000) 
+++

The GPU is the most powerful processing unit in a modern computer, and sadly, much of its processing
power ends up wasted as it waits for data from the CPU to draw. It is common to move linear algebra 
tasks like model skinning and physics to the GPU, but other tasks are often still left for the GPU. {:truncate}

In Araboth, one of the most expensive operations was computing the vertex positions of a laser. 
Since the common case was about 40 lasers of 32 subdivisions each, this led to something like 5000 vertex 
(and therefore 5000 spline evaluations and 5000 velocity evaluations) every frame. THe sheer volume of computations
was overwhemling the GPU. {:longtruncate}

Approaching this performance from a higher level, it becomes clear that the speed (or lack thereof) 
is due to 3 major factors - Main memory to Video memory transfers of large amounts of data, 
recomputing the splines repeitively on the CPU, and a lack of cache coherency. The only one that
could have been fixed trivially might have improved performance a bit, but it would not have been
enough to guarantee a consistent 30 frames per second at high load. Thankfully, this operation
has an obvious GPU accelerated solution.

The first step is to rethink the way that data is used and applied to each vertex. Originally, the 
spline would be used to compute its position and velocity at each time step, and a vertex
position was derived from these positions. However, GPU acceleration works much better
if the data access is turned around - the vertex computs its position from the spline parameters. This
allows for vertices to be static, and then transformed in the vertex shader. Since the majority of 
the vertex attributes can be computed in the vertex shader, the only attributes that are required 
for computation have to be transfered from main memory to video memory. My implementation takes
4 floats per vertex.

However, this change requires that the spline for each laser be passed in through as well. Since
each vertex gets its own copy to work on, this at first glance may force us to either pass in 
4 float3 control points per vertex, which will destroy our memory transfer optimization, 
or pass in our control points through a shader uniforms (which limits how many lasers we can show), 
or pass them through a texture lookup table, which also destroys our memory transfer rates. Thankfully,
DirectX supports geometry instancing through stream multiplexing, so it is possible to instance one collection
of control points to an entire laser, saving a tremendous amount of memory bandwidth.

After the initial choices of data representation, the rest of the implementation is mostly detail. 
Unrolling the spline function that was used in the original implementation and passing it through
wolfram alpha for minimization, we get

~~~~"cpp"
float3 evaluate(float3 a, float3 b, float3 c, float3 d, float t)
{
	float t3 = t * t * t;
	float t2 = t * t;

	//-a t^3+3 a t^2-3 a t+a+3 b t^3-6 b t^2+3 b t-3 c t^3+3 c t^2+d t^3
	//This expansion is derived from wolfram alpha.
	float3 p = 
		 -1 * a * t3 + 3 * a * t2 - 3 * a * t + a
		+ 3 * b * t3 - 6 * b * t2 + 3 * b * t 
		- 3 * c * t3 + 3 * c * t2 
		+ d * t * t3;
	return p;
}
~~~~

The rest of the vertex shader is basically expanding the vertex attributes into its components,
marshalling the spline vectors into the evaluate method, and then computing its texture coordinates
and colors, which is irrelevant here.

This design netted a huge improvement in performance. Prior to this, it was possible to render a max 
of about 20 lasers at 30 frames per second. Now, with instancing and spline computations done on the 
GPU, it is possible to render over 200 lasers at over 30 frames per second. A combination of moving
computations to the GPU from the CPU and reducing the huge memory transfer overhead netted a
10x speedup. 

This optimization is not trivial, and is very much application dependent. However, if it is possible
to move a comptuation from the processor to the graphics card, it is almost always worth it.
