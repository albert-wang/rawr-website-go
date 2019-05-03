+++
Title: Fundamentals of Graphic Programming
Category: 3
Hero: 
Publish: <nil>
+++

This is a fairly basic review of the fundamentals of graphic programming, in both terms and mathematics, along with possible performance characteristics.{:truncate}{:longtruncate} Here, I will use the conventions that my game are done in, which is a right-handed coordinate system, with column vectors.

### Basic Linear Algebra
Consider the matrix $$A = [[11, 12, 13, 14],[21, 22, 23, 24],[31, 32, 33, 34],[41, 42, 43, 44]]$$ and the vector $$vec v = (x, y, z, w)$$

Generally, a square matrix of size $$n$$ is represented in memory as a linear array of $$n\*n$$ elements. There are two obvious ways of storing a matrix, in the form `[11, 12, 13, 14, 21, 22, ... , 44]`, which is known as row-major, or in the form `[11, 21, 31, 41, 12, ..., 44]`, which is known as column major.

Since these arrays are linear in memory, a programmer can load a $$4\*4$$ matrix into 4 SIMD registers. Doing matrix math with 4-element wide registers can be significantly faster than doing them with standard FPU operations, though the speedup depends on your memory layout. A row-major matrix with column vectors multiplied as $$vec vA$$ 

