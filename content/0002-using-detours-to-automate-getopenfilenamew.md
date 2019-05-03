+++
Title: using detours to automate getopenfilenamew
Category: 3
Hero: 
Publish:  Jan 23, 2012 3:12am (+0000) 
+++

One of the easier things to do with the Win32 API is to automate an application. For any
given application, the vast majority of your interactions with it will end up being through
the Win32 GUI, or through something wrapping the Win32 GUI. As a result, it is very easy to 
enumerate the windows within the interface and send a bunch of messages through SendMessage 
and similar methods. {:truncate}

This approach works very well, as long as you only have to do interactions within one parent window, 
and starts to break down when the application begins to start using the common dialog box library.
Since these create new windows, you need to have a consistent way to hook into the window creation, or
just guess, neither of which are easy or acceptable. After a bit of thinking, it came to me that you don't 
even have to let the window pop up. If you can overwrite the call to (in my case) GetOpenFileNameW, you
can elide window creation and just directly write the results to the output buffer. {:longtruncate}

In order to overwrite the call, I used the [Detours](http://research.microsoft.com/en-us/projects/detours/) 
library, which is a DLL Injection library made by Microsoft Research. While the professional version is required
for injecting into 64 bit processes and commercial development, the trial version is perfectly fine for 32 bit 
processes. After a quick installation, the actual usage of Detours is very easy. 

First, setup a dummy function that gets exported to ordinal 1. This is what detours requires as a target
for the injection process. To do this consistently, you need to write a def file - 

~~~~"cpp"
LIBRARY DETOURED
EXPORTS
	DummyFunction @1

~~~~

Then the implementation of the new method, 

~~~~"cpp"
//GetOpenFileNameW is BOOL WINAPI GetOpenFileNameW(LPOPENFILENAME)
//This method dosn't even have to be dll exported here.
BOOL WINAPI CustomGetOpenFileNameW(LPOPENFILENAME params)
{
	return TRUE;
}

~~~~

And then handle the attachment and detachment events by detouring and undetouring GetOpenFileName. 
These operations are done in a transaction to make sure that these changes are done atomically, so
nothing strange happens while rewriting the target application in memory.


~~~~"cpp"
BOOL WINAPI DllMain(HINSTANCE, DWORD reason, LPVOID)
{
	//Notice that error detection here is ommited for
	//berevity.
	switch(reason)
	{
		case DLL_PROCESS_ATTACH:
			DetourTransactionBegin();
			DetourUpdateThread(GetCurrentThread()):
			DetourAttach(&(PVOID&)GetOpenFileNameW, CustomGetOpenFileNameW);
			DetourTransactionCommit();
			break;

		case DLL_PROCESS_DETACH:
			DetourTransactionBegin();
			DetourUpdateThread(GetCurrentThread());
			DetourDetach(&(PVOID&)GetOpenFileNameW, CustomGetOpenFileNameW);
			DetourTransactionCommit();
			break;
	}
}
~~~~

This should be built into a DLL, which by itself is useless. The driver 
program needs to set up the data that the injected DLL uses, and then invoke `DetourCreateProcessWithDll` 
to actually create the process and inject the DLL.

Since the launcher process and the DLL are executing in separate processes, you also 
need to setup some sort of IPC to communicate intent. The way I set it up, the 
detoured DLL just overwrites GetOpenFileNameW with a function that returns the same 
data each time, so the launcher just sets up a shared memory region and waits
for the DLL to read from it.

