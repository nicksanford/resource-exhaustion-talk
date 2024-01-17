package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
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
	p := path.Join(os.TempDir(), "resource-exhaustion")
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

	c := make(chan int)
	wg := sync.WaitGroup{}
	defer wg.Wait()
	workFunc := func() {
		defer wg.Done()
		for {
			tmpI := <-c
			if tmpI%1000 == 0 {
				log.Printf("start %d\n", tmpI)
			}
			f, err := os.CreateTemp(p, "")
			if err != nil {
				log.Fatalf(err.Error())
			}
			f.Close()
			if tmpI%1000 == 0 {
				log.Printf("end %d\n", tmpI)
				// if err := prompt("continue"); err != nil {
				// 	return
				// }
			}
		}
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go workFunc()
	}

	for i := 0; i < NUM_FILES; i++ {
		c <- i
	}
}
