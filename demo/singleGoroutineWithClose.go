package demo

import (
	"log"
	"os"
)

func SingleGoroutineWithClose(p string) {
	files := []*os.File{}
	for fileNum := 0; fileNum < NUM_FILES; fileNum++ {
		logStartEvery(fileNum, fileNum, 1000)
		f, err := os.CreateTemp(p, "")
		if err != nil {
			log.Fatalf(err.Error())
		}
		files = append(files, f)
		f.Close()
		logEndEvery(fileNum, fileNum, 1000)
	}
}
