package demo

import "log"

func RunOutOfMemoryImmediately() float64 {
	f := make([]float64, EightyGigsWorthOfFloats)
	return f[0]
}

func RunOutOfMemory() float64 {
	var ret float64
	for i := 0; i < 10; i++ {
		log.Printf("allocating at iteration %d\n", i)
		f := make([]float64, EightGigsWorthOfFloats)
		ret = f[0]
	}
	return ret
}
