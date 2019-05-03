+++
Title: writing native node.js in windows
Category: 2
Hero: 
Publish: 2012-01-13 06:35:06.510752 +0000 +0000
+++

As of Node.js 0.7.0-alpha, there is no blessed path for writing native Node.js extensions on windows. 
However, it is not impossible to build fully compatible native node modules using Visual Studio. Doing 
so properly, however, requires a bit of finnesse.{:truncate}{:longtruncate}

The first step is relatively easy - building Node.js from source. The node.js installation packages
include the node binary, but no libraries. Clone the repository and run vcbuild.bat from the
visual studio command line and you should have the build results in Release/ or Debug/. There are
utility libraries in $(CONFIGURATION)/lib, but those seem to be built under different settings
than the main node.lib static library.

Writing a native library is slightly harder. The entry point for node modules is the init function.
This function should be dllexported and extern C'ed to ensure that the linker can see your method. 
Since the dllexport is very platform specific, node.js actually provides a macro to expediate this process - 

````
extern "C"
{
	NODE_EXPORT void init(v8::Handle<v8::Object> env);
}
````{: class="brush: cpp"}

This function is called when your extension is loaded, and where you should register
all your prototype methods, your classes, and so forth. The internals of libuv and the v8 engine
are complicated, so in this example we won't bind anything and do a simple operation in init.

````
void init(v8::Handle<v8::Object> env)
{
	std::cout << "Initialized native library\n";
}
````{: class="brush: cpp"}

Make sure that your output settings are set to build a dynamic link library, and your 
linker settings are set correctly - only add in node.lib as linker input, similar
to the image below.

<a href='//img.rawrrawr.com/gallery/ad89bde4275e21b002ed544add13560e9f3c1a62.png' class='imglink'>
	<img src='//img.rawrrawr.com/gallery/med-ad89bde4275e21b002ed544add13560e9f3c1a62.png' />
</a>

Build your application, and rename the output dll so that it ends in .node. In your node application, 
simply require the library name, and you should see "initialized native library" in the console.