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

	log.Println("starting")
	tmpDir := os.TempDir()
	log.Printf("tempdir: %s\n", tmpDir)
	p := path.Join(tmpDir, "resource-exhaustion")
	log.Println(p)

	log.Printf("rm -rf %s\n", p)
	if err := os.RemoveAll(p); err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("mkdir -p %s\n", p)
	if err := os.MkdirAll(p, 0o755); err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("creating temp file")
	_, err := os.CreateTemp(p, "")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for i := 0; i < NUM_FILES; i++ {
		log.Printf("creating temp file %d\n", i)
		if i%1000 == 0 {
			log.Printf("start %d\n", i)
		}
		// go func() {
		_, err := os.CreateTemp(p, "")
		if err != nil {
			log.Fatalf(err.Error())
		}
		if i%1000 == 0 {
			log.Printf("end %d\n", i)
		}
		// }()
	}
}
