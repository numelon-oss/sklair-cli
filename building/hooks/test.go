package hooks

import (
	"sklair/luaSandbox"
)

// TODO
func hi() {
	osExitChan := make(chan int)
	defer close(osExitChan)

	L := luaSandbox.NewSandbox(luaSandbox.SandboxOptions{
		ExitChannel: osExitChan,
	})
	defer L.Close()

	if err := L.DoString(`
print(fs.read("go.mod"))
`); err != nil {
		panic(err)
	}
}
