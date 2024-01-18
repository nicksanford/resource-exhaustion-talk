package demo

import (
	"log"
	"os"
	"runtime"
	"sync"
)

func FiniteNumberOfGoRoutines(p string) {
	fileNumChan := make(chan int)
	wg := sync.WaitGroup{}
	defer wg.Wait()
	workFunc := func(goRoutineNum int) {
		defer wg.Done()
		for {
			fileNum := <-fileNumChan
			logStartEvery(goRoutineNum, fileNum, 1000)
			f, err := os.CreateTemp(p, "")
			if err != nil {
				log.Fatalf(err.Error())
			}
			f.Close()
			logEndEvery(goRoutineNum, fileNum, 1000)
		}
	}

	numCPUs := runtime.NumCPU()
	for goRoutineNum := 0; goRoutineNum < numCPUs; goRoutineNum++ {
		wg.Add(1)
		go workFunc(goRoutineNum)
	}

	for fileNum := 0; fileNum < NUM_FILES; fileNum++ {
		fileNumChan <- fileNum
	}
}
