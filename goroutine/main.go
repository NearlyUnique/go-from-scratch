package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t0 := time.Now()
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Doing work ... %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("Took:%v\n", time.Since(t0))
}
