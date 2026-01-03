package luaSandbox

import (
	"os"

	lua "github.com/yuin/gopher-lua"
)

func openFs(L *lua.LState) int {
	fsMod := L.RegisterModule("fs", fsFuncs)
	L.Push(fsMod)
	return 0
}

var fsFuncs = map[string]lua.LGFunction{
	"read":  readFile,
	"write": writeFile,
}

// TODO: REAL SANDBOXING HERE!!!

// readFile reads the contents of a file specified by the first argument and returns the data and a potential error.
// on success, returns the data as a string. on error, returns nil and the error message.
func readFile(L *lua.LState) int {
	name := L.CheckString(1)

	data, err := os.ReadFile(name)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(data))
	return 1
}

// writeFile writes the contents of the second argument to the file specified by the first argument.
// on success, returns nothing. on error, returns the error message.
func writeFile(L *lua.LState) int {
	name := L.CheckString(1)
	data := L.CheckString(2)

	if err := os.WriteFile(name, []byte(data), 0644); err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	return 0
}

// TODO: create fs.scandir

// TODO: ENTIRE FILESYSTEM LIBRARY NEEDS TO ONLY ACCEPT THESE FOUR TYPES OF PATHS:
// cache:file.txt -> .sklair/cache/file.txt
// project:file.txt (only this one allows READONLY access to one level above the project directory, using project:../file.txt)
// temporary:file.txt -> .sklair/tmp/file.txt
// generated:file.txt -> .sklair/generated/file.txt -> build -> build/_sklair/generated/file.txt
