+++
Title: Fundamentals of Graphic Programming
Category: 3
Hero: 
Publish: <nil>
+++

This is a fairly basic review of the fundamentals of graphic programming, in both terms and mathematics. Here, I will use the conventions that my game are done in, which is a right-handed coordinate system, with column vectors.

### Matrix operations.
Generally, a square matrix of size $$n$$ is represented in memory as a linear array of $$n*n$$ elements. Generally, I will only be concerned about $$4*4$$ matrices in this discussion. There are two major ways of storing a matrix $$[[11, 12, 13, 14],[21, 22, 23, 24],[31, 32, 33, 34],[41, 42, 43, 44]]$$, in the form `[11, 12, 13, 14, 21, 22, ... , 44]`, which is known as row-major, or in the form `[11, 21, 31, 41, 12, ..., 44]`, which is known as column major.

Generally, this doesn't affect 

