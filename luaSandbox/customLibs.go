package luaSandbox

import (
	lua "github.com/yuin/gopher-lua"
	json "layeh.com/gopher-json"
)

var customLibs = []luaLib{
	{"fs", openFs},
	{"json", json.Loader},
}

func OpenSandboxedCustom(ls *lua.LState, opts SandboxOptions) {
	for _, lib := range customLibs {
		ls.Push(ls.NewFunction(lib.libFunc))
		ls.Push(lua.LString(lib.libName))
		ls.Call(1, 0)
	}
}
