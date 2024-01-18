package demo

import (
	"log"
	"os"
)

func SingleGoroutineNoClose(p string) {
	files := []*os.File{}
	for fileNum := 0; fileNum < NUM_FILES; fileNum++ {
		logStartEvery(1, fileNum, 1000)
		f, err := os.CreateTemp(p, "")
		if err != nil {
			log.Fatalf(err.Error())
		}
		// To prevent the garbage collector from cleaning up files
		files = append(files, f)
		logEndEvery(1, fileNum, 1000)
	}
}
