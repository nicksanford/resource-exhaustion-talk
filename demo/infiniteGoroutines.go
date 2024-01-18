package demo

import (
	"log"
	"os"
)

func InfiniteNumberOfGoRoutines(p string) {
	for fileNum := 0; fileNum < NUM_FILES; fileNum++ {
		go func() {
			logStartEvery(fileNum, fileNum, 1000)
			f, err := os.CreateTemp(p, "")
			if err != nil {
				log.Fatalf(err.Error())
			}
			f.Close()
			logEndEvery(fileNum, fileNum, 1000)
		}()
	}
}
