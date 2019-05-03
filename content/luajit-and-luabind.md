+++
Title: Luajit and Luabind
Category: 1
Hero: 
Publish: 2012-05-26 19:24:50 +0000 +0000
+++

Just a quick aside for anyone that may be using [luajit](http://luajit.org/) with [luabind](http://www.rasterbar.com/products/luabind.html) on x86 platforms - be careful about the way that luabind throws exceptions. Luajit on x86 platforms has no exception interopability, which means that throwing exceptions through lua frames is not supported, and may corrupt the stack. This usually is not a problem, since luabind generally only throws exceptions after an error, or some other confluence of events conspires so that the exception does not corrupt anything. {:truncate}{:longtruncate}
However, the following code will create a buffer overflow in the lua internals -

```cpp
void show() {}

// ... Setup omitted. Assume show is bound to lua.
lua_loadstring(L, "show(1);");
luabind::object obj(luabind::from_stack(L, -1));
obj();
```

A workaround, without having to remove exception support from luabind, is to use the raw lua API to invoke a chunk, instead of turning it into a luabind::object - that is, invoke lua_pcall by hand and handle errors manually.