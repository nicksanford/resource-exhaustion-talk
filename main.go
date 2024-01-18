package main

import (
	"log"
	"os"
	"path"

	"github.com/nicksanford/resource-exhaustion-talk/demo"
)

func main() {
	log.Printf("PID: %d\n", os.Getpid())
	demo.Prompt()
	// demo.RunOutOfMemoryImmediately()
	// demo.RunOutOfMemory()
	p := path.Join(os.TempDir(), "resource-exhaustion")
	demo.CleanTempDir(p)
	demo.CreateTempDir(p)

	demo.SingleGoroutineNoClose(p)
	// demo.SingleGoroutineWithClose(p)
	// demo.InfiniteNumberOfGoRoutines(p)
	// demo.FiniteNumberOfGoRoutines(p)
}
