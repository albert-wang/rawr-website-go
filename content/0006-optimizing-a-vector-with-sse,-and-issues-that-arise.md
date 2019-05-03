+++
Title: optimizing a vector with sse, and issues that arise
Category: 2
Hero: 
Publish:  Mar 13, 2012 6:40am (+0000) 
+++

According to the most recent Steam Hardware Survey, 99% of steam users have SSE2 and SSE3
instruction support. The most immediately obvious thing to do with SSE is to try to optimize
4-vector math. However, a naiive reimplementation of vector methods doesn't come close to the theoretical 4x improvement that one would expect from an SSE implementation {:truncate}{:longtruncate}

Lets start with the naiive implementation of a vector4, and the naiive implementation of 
operator +=:

~~~~"cpp"
struct Vector4
{
	//Assume that this has constructors
	//and other reasonable methods on it,
	//like dot, operator[], other operators, 
	//implemented in the most straightforward way.


	Vector4& operator+=(const Vector4& other)
	{
		for (size_t i = 0; i < 4; ++i)
		{
			elements[i] += other.elements[i];
		}

		return this;
	}
	
	float elements[4];	
};

~~~~

Wonderful. Next, a test rig just to test this basic functionality:

~~~~"cpp"
Vector4 v = Vector4(r[0], r[1], r[2], r[3]);
Vector4 b = Vector4(r[4], r[5], r[6], r[7]);

LARGE_INTEGER begin;
QueryPerformanceCounter(&begin);

Vector4 c = Vector4::Zero;
for (size_t i = 0; i < 100000000; ++i)
{
	c += v + b;
	c -= b + v;
}

LARGE_INTEGER end;
QueryPerformanceCounter(&end);

LARGE_INTEGER freq;
QueryPerformanceFrequency(&freq);

double time = (double)(end.QuadPart - begin.QuadPart) / (freq.QuadPart);
std::cout << "Took: " << 1000 * time << "ms\n";
~~~~

The expected output is [0, 0, 0, 0], and r is populated with random data (no denormals). 

Using the naiive implementation, with /Ox (Full optimization) and no architecture flags, 
the generated code simply loads up every floating point value, converts to double, adds it, and
then converts backwards. The core of the operation is shown below - a bunch of flds, fadds and fstp. 

~~~~
009714AB  fadd        dword ptr [esp+5Ch]  
009714AF  fstp        dword ptr [esp+98h]  
009714B6  fld         dword ptr [esp+8Ch]  
009714BD  fadd        dword ptr [esp+7Ch]  
009714C1  fstp        dword ptr [esp+7Ch]  
009714C5  fld         dword ptr [esp+90h]  
009714CC  fadd        dword ptr [esp+80h]  
009714D3  fstp        dword ptr [esp+80h]  
009714DA  fld         dword ptr [esp+94h]  
009714E1  fadd        dword ptr [esp+84h]  
009714E8  fstp        dword ptr [esp+84h]  
009714EF  fld         dword ptr [esp+98h]  
009714F6  fadd        dword ptr [esp+88h]  
009714FD  fstp        dword ptr [esp+88h]  
~~~~

The problem is that the floating point model is set to precise, which forces 
80-bit precision on the FPU. Changing the floating point model to fast does 
not net any gains in this department - the generated code is functionally
identical between the versions.

The baseline performance on my i7 here is roughly 810 milliseconds with /fp:fast, and 1080
milliseconds with /fp:precise. This means that the target performance for this loop is 202.5
milliseconds, which will represent a 4x speedup. 

The least labor intensive way of preforming SSE optimizations is to let the compiler do it for
you. While Visual Studio 2010 doesn't allow targeting SSE3 architectures specifically, the 
SSE3 instructions do not really add very much - the only interesting instruction for 
short vector operations is hadd, the horizontal add. HADD can be used to implement
SSE dot products (ignoring the SSE4 dot product intrinsic, which is actually pretty slow), 
but only shows significant gains if you do 4 dot products at once. 

Either way, enabling /arch:SSE2 produces shockingly bad code. Basically, it loads every
element into its own XMM register with movss, and then adds them one by one. Since there
are not enough registers to load all 13 floating point values at one value per register,
the compiler generates a load of movss into memory, which incurs even more slowdown.

~~~~
;;; This loads up a bunch of elements into XMM registers 0-3.
0024139E  movss       xmm1,dword ptr [esp+80h]  
002413A7  movss       xmm2,dword ptr [esp+84h]  
002413B0  movss       xmm3,dword ptr [esp+88h]  
002413B9  addss       xmm1,dword ptr [esp+70h]  
002413BF  addss       xmm2,dword ptr [esp+74h]  
002413C5  addss       xmm3,dword ptr [esp+78h]  
002413CB  movss       dword ptr [esp+30h],xmm0  
002413D1  movss       xmm0,dword ptr [Math::Vector<float,4>::Zero+4 (24631Ch)]  
002413D9  movss       dword ptr [esp+34h],xmm0  
002413DF  movss       xmm0,dword ptr [Math::Vector<float,4>::Zero+8 (246320h)]  
002413E7  movss       dword ptr [esp+38h],xmm0  
002413ED  movss       xmm0,dword ptr [Math::Vector<float,4>::Zero+0Ch (246324h)]  
002413F5  movss       dword ptr [esp+3Ch],xmm0  
002413FB  movss       xmm0,dword ptr [esp+7Ch]  
00241401  addss       xmm0,dword ptr [esp+6Ch]  

;;; Integer conversion code omitted 

00241442  movss       xmm4,dword ptr [esp+10h]  
00241448  movaps      xmm5,xmm4  
0024144B  addss       xmm5,xmm6  
0024144F  addss       xmm5,dword ptr [esp+30h]  
00241455  movaps      xmm6,xmm4  
00241458  addss       xmm6,xmm7  
0024145C  movss       dword ptr [esp+54h],xmm6  
00241462  movaps      xmm6,xmm4  
00241465  movaps      xmm7,xmm2  
00241468  addss       xmm6,xmm7  
0024146C  movss       xmm7,dword ptr [esp+78h]  
00241472  addss       xmm6,dword ptr [esp+38h]  
00241478  addss       xmm7,xmm4  
0024147C  addss       xmm7,dword ptr [esp+3Ch]  
00241482  movss       dword ptr [esp+3Ch],xmm7  
~~~~

After all of this, the actual performance is actually a bit better - 700 milliseconds 
for the same loop, a 1.15x improvement on the naiive implementation, but still 
far from the idealized 210 millisecond runtime that an SSE implementation should aim for.

Before actually implementing the SSE version, certain constraints need to be added here.
SSE performs optimally when loading from 16 byte aligned addresses, which means that, by 
default, any heap allocation will likely be incompatible with the movps instructions. One can
specify that any stack allocation should be aligned with a declspec, but this causes 
certain problems. In particular, the vector implementation that ships with Visual Studio 
contains at least this troublesome line:

~~~~"cpp"
void resize(size_type _Newsize, _Ty _Val)
~~~~

When _Ty is an aligned type, the template will fail to compile, since it is not possible, 
in the general case, to push an aligned value onto the stack in preparation for a function call.
Changing the container to use boost::container::vector, which uses a constant reference instead
of a value, allows the container to compile. 

However, actually using the container causes memory problems - movps expects an aligned 
memory address, but the container doesn't guarantee any such alignment. This will result 
in a runtime crash. A solution here is to store pointers to your data, but this design 
will invalidate your cache like mad. A proper solution is to use an aligned allocator, 
but this causes template type issues. Without support for template template aliases, the 
next best solution is to use some type computation to generate a container with the 
properly aligned allocator, but this has maintainability problems. 

Another problem to consider here is that repeated calls to movps actually will do redundant
loads if you're not careful. The easiest way to make sure that move instructions are generated
only when they are required is to invoke a small bit of undefined behavior, and store both the
__m128 register and the floating point representation together: 

~~~~"cpp"
__declspec(align(16)) union storage
{
	__m128 sse
	float  fpu[4];
};
~~~~

The declspec above is not strictly required, since that can be inferred from the __m128 value,
but it doesn't harm the code generation at all, and it is always nice to be explicit. The actual
implementation lives in the [rmath](https://bitbucket.org/rraawwrr/rmath) repository, and not
be posted here. The generated code here is not quite optimal, but it is reasonable:

~~~~
00F71406  movaps      xmm2,xmmword ptr [esp+40h]  
00F7140B  movdqa      xmm1,xmmword ptr [Math::Vector<float,4>::Zero (0F766C0h)]  
00F71413  addps       xmm2,xmmword ptr [esp+10h]  

;;; Integer division and conversion ommitted ...

; Move (float)(i / 200) into xmm0, and shuffle it into all components, and 
; some more math operations 

00F71441  movss       xmm0,dword ptr [esp+0Ch]  
00F71447  fstp        dword ptr [esp+10h]  
00F7144B  shufps      xmm0,xmm0,0  
00F7144F  addps       xmm0,xmm2  
00F71452  addps       xmm0,xmm1  
00F71455  movss       xmm1,dword ptr [esp+10h]  
00F7145B  shufps      xmm1,xmm1,0  
00F7145F  addps       xmm1,xmm2  
00F71462  subps       xmm0,xmm1  
00F71465  movaps      xmm1,xmm0  
~~~~

This code executes in 280 milliseconds, an 2.8x improvement over the original, non-SSE code. 
Interestingly, the vast majority of the time is spent within the integer division and subsequent
floating point conversion. Removing the code shaves 200 milliseconds off the naiive implementation, 
to 600ms, and improves the sse implementation to 175ms, a 3.4x improvement. This improvement 
is significant, but only really affects vector computation heavy code, which should have been 
already instruction parallelized. Using an SSE version of a vector also requires special 
container considerations, which might require significant code rework. Considering the 
actual level of code that would benefit from this optimization, and the amount of code
that would be required to change out all the containers and allocators, it may not even be a good
cost/benefit ratio.
