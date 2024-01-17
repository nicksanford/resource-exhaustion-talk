package main

import (
	"log"
	"os"
	"path"
)

const NUM_FILES = 10000000000

func main() {
	tmpDir := os.TempDir()
	log.Printf("tempdir: %s", tmpDir)
	for i := 0; i < NUM_FILES; i++ {
		go func() {
			_, err := os.CreateTemp(path.Join(tmpDir, "resource-exhaustion"), "")
			if err != nil {
				log.Fatalf(err.Error())
			}
		}()
	}
}
