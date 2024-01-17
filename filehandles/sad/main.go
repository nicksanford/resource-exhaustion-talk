package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const NUM_FILES = 10000000000

func main() {
	log.Printf("PID: %d\n", os.Getpid())

	var start string
	for {
		log.Println("start? y/n")
		fmt.Scan(&start)
		start = strings.ToUpper(start)
		if start == "Y" {
			break
		}

		if start == "N" {
			return
		}
	}

	tmpDir := os.TempDir()
	log.Printf("tempdir: %s\n", tmpDir)
	p := path.Join(tmpDir, "resource-exhaustion")
	log.Println(p)

	if err := os.RemoveAll(p); err != nil {
		log.Fatalf(err.Error())
	}

	if err := os.MkdirAll(p, 0o755); err != nil {
		log.Fatalf(err.Error())
	}

	_, err := os.CreateTemp(path.Join(tmpDir, "resource-exhaustion"), "")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for i := 0; i < NUM_FILES; i++ {
		// go func() {
		_, err := os.CreateTemp(path.Join(tmpDir, "resource-exhaustion"), "")
		if err != nil {
			log.Fatalf(err.Error())
		}
		// }()
	}
}
