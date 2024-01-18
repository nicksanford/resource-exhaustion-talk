package demo

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Prompt() {
	var answer string
	for {
		fmt.Println("Paused\nquit? y/n")
		fmt.Scan(&answer)
		answer = strings.ToUpper(answer)
		if answer == "N" {
			break
		}

		if answer == "Y" {
			log.Fatal("quitting")
		}
	}
}

func CreateTempDir(p string) {
	log.Printf("mkdir -p %s\n", p)
	if err := os.MkdirAll(p, 0o755); err != nil {
		log.Fatalf(err.Error())
	}
}

func CleanTempDir(p string) {
	log.Println("clearing temp dir tempDir: ", p)
	log.Println("running: rm -rf ", p)
	if err := os.RemoveAll(p); err != nil {
		log.Fatalf(err.Error())
	}
}

func logStartEvery(goRoutineNumber, fileNumber, mod int) {
	if fileNumber%mod == 0 {
		log.Printf("start of GoRoutine: %d, file: %d\n", goRoutineNumber, fileNumber)
	}
}

func logEndEvery(goRoutineNumber, fileNumber, mod int) {
	if fileNumber%mod == 0 {
		log.Printf("end of GoRoutine: %d, file: %d\n", goRoutineNumber, fileNumber)
	}
}
