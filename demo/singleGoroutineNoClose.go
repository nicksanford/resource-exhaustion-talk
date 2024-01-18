package demo

import (
	"log"
	"os"
)

func SingleGoroutineNoClose(p string) {
	for fileNum := 0; fileNum < NUM_FILES; fileNum++ {
		logStartEvery(1, fileNum, 1000)
		_, err := os.CreateTemp(p, "")
		if err != nil {
			log.Fatalf(err.Error())
		}
		logEndEvery(1, fileNum, 1000)
	}
}
