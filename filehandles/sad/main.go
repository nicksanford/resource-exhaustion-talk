package main

import (
	"log"
	"os"
)

const NUM_FILES = 10000000000

func main() {
	log.Printf("tempdir: %s", os.TempDir())

	for i := 0; i < NUM_FILES; i++ {
		_, err := os.CreateTemp("", "resource-exhaustion")
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
}
