+++
Title: Pipelining - Why it Matters
Category: 2
Hero: 
Publish: 2012-07-24 19:05:19 +0000 +0000
+++

While reading Hacker News, I came upon an article that purported that Java (or other high level languages) should be used in preference to C++ due to the quality of modern interpreters and JIT compilers. I will not discuss the conclusion, which is shaky at best, but rather the microbenchmark that was used to display the quality of the Java virtual machine.{:truncate}{:longtruncate}

The benchmark used was a cumulative sum method:

```cpp
void slowishSum(vector<int> & data) 
{
    if(data.size() == 0) 
        return;

    for (size_t i = 1; i < data.size(); ++i) 
    {
        data[i] += data[i - 1];
    }
}
```

On my machine, this does about 460~ million integers per second. This would be great, if it was all that the processor could do - but its not.

Many articles have been written on the benefits of cache coherency and data locality, with the idea that if you process your data in order, the CPU will load up the requested data (and a little more) into L1 cache, making further accesses to memory around there very fast. Cache awareness in programming can lead to significant performance improvements. Changing an algorithm from a AOS (Array of Structures) access pattern to a SOA (Structure of Arrays) access pattern can improve performance upwards of 200%. 

This does not apply to the cumulative sum function. All memory accesses here are sequential, allowing the CPU to have optimal cache use. So how do you make this method faster? One can unroll the loop, through usage of -funroll-loops (the best gcc flag name ever), or manually: 

```cpp
void sum(vector<int> & data) 
{
    if(data.size() == 0) 
        return;
 
    size_t i = 1;
    for (; i < data.size() - 1; i+=2) 
    {
        data[i] += data[i - 1];
        data[i+1] += data[i ];
    }

    for (; i != data.size(); ++i) 
    {
        data[i] += data[i - 1];
    }
}
```

This does not help, however, computing a similar 470~ million integers per second.

To further optimize this method, a basic understanding of CPU architecture is required. A very basic view of CPU operation is that it goes through several distinct steps while processing an instruction - fetching the instruction from memory, decoding the instruction, execution of the instruction, and writing the results back into memory. Each of these steps are independent of each other, so a basic CPU can pipeline four instructions at once - one doing writeback, one doing the execution, one decoding and the last fetching from memory. Modern CPUs are more complicated, and their pipelines deeper, but the general concept remains the same. However, pipelining introduces data hazards. Consider a program that executes 'INC A; STORE B A'. Without hazard mitigation, the increment may execute, and then while it is in the write-back stage, the store instruction executes, fetching the unincremented value of A and storing it into B. Hazards can be mitigated by hand, inserting additional instructions into the pipeline so that this situation never happens, or by the compiler, doing a similar operation. At the CPU level, the CPU can detect data dependencies and intentionally add in bubbles to the pipeline, so that most hazards are avoided.

A naive cumulative sum implementation, as in the first code block above, is almost the worst case for the CPU pipeline. Each iteration of the loop needs to load up two memory locations, one of which is dependent upon the operation in the previous loop. Even 
considering the instructions that are used to do loop maintenance (incrementing the counter, comparing against the size, jump if zero), the CPU must still insert a bubble to delay the dependent load. With processors having rather long pipelines, this can cause a tremendous slowdown. Therefore, removing this bubble should show some performance improvement. Without thinking too much, it is completely possible to compute several elements of the cumulative sum at the same time, without having any dependent data loads within the body of the loop - 

```cpp
void avoidStalls(vector<int>& data)
{
	if (data.size() == 0) 
	{
		return;
	}

	int remain = data[0];
	size_t i = 1;
	size_t size = data.size() - 2;
	for (i = 1; i < size; i += 3) 
	{
		int a = data[i + 0];
		data[i + 0] = remain + a;

		int b = data[i + 1];
		data[i + 1] = remain + a + b;

		int c = data[i + 2];
		data[i + 2] = remain + a + b + c;

		remain = data[i + 2];
	}

	for (; i < data.size(); ++i) 
	{
		data[i] += data[i - 1];
	}
}
```

At first glance, this code looks less efficient, since there is a lot of repeated work. However, doing the computation this way removes most of the data dependencies from the loop. A quick performance comparision shows that this version can compute 1040 million integers per second or more, an improvement of over 100%. This benchmark was compiled on MSVC 2012, with default settings for the 'Release' configuration. Other people, using GCC, have reported numbers similar (and some much higher) than those I obtained using the stall avoiding algorithm. It may be that they have better computers, or that GCC is capable of optimizing this dependent load into a better-for-the-pipeline algorithm, but I cannot make any definitive remarks on that.

Practically, I would like to say that it is always important to consider the processor and its memory while writing performance oriented code, but this is honestly a micro-optimization, and should be looked over for more fundamental changes for performance. In general, it may be hard to find patterns that are so obviously stalling the pipeline, and harder to change the algorithm to avoid such stalls. In any case, knowing this property of processors is always helpful while writing performance sensitive code, so beware.