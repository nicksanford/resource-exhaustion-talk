package main

import (
	"log"
	"os"
	"path"
)

const NUM_FILES = 1 //0000000000

func main() {
	tmpDir := os.TempDir()
	log.Printf("tempdir: %s", tmpDir)
	if err := os.MkdirAll(path.Join(tmpDir, "resource-exhaustion"), 0o650); err != nil {
		log.Fatalf(err.Error())
	}
	_, err := os.CreateTemp(path.Join(tmpDir, "resource-exhaustion"), "")
	if err != nil {
		log.Fatalf(err.Error())
	}
	// for i := 0; i < NUM_FILES; i++ {
	// 	go func() {
	// 		_, err := os.CreateTemp(path.Join(tmpDir, "resource-exhaustion"), "")
	// 		if err != nil {
	// 			log.Fatalf(err.Error())
	// 		}
	// 	}()
	// }
}
