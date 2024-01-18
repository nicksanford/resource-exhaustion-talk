package main

import (
	"fmt"
	"os"
	"path"

	"github.com/nicksanford/resource-exhaustion-talk/demo"
)

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	demo.Prompt()
	p := path.Join(os.TempDir(), "resource-exhaustion")
	demo.CreateTempDir(p)

	// debug.SetMaxThreads(1000000000000000)
	// demo.InfiniteNumberOfGoRoutines(p)
	// demo.FiniteNumberOfGoRoutines(p)
	// demo.SingleGoroutineNoClose(p)
	// demo.SingleGoroutineWithClose(p)
	// demo.RunOutOfMemoryImmediately()
	demo.RunOutOfMemory()
}
