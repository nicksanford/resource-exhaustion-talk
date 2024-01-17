package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const NUM_FILES = 10000000000

func prompt(msg string) error {
	var answer string
	for {
		log.Printf("%s? y/n\n", msg)
		fmt.Scan(&answer)
		answer = strings.ToUpper(answer)
		if answer == "Y" {
			return nil
		}

		if answer == "N" {
			return errors.New("user said no")
		}
	}
}

func main() {
	log.Printf("PID: %d\n", os.Getpid())
	if err := prompt("start"); err != nil {
		return
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
			if err := prompt("continue"); err != nil {
				return
			}
		}
		// }()
	}
}
