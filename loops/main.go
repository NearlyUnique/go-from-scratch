package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Doing work ... %d\n", i)
	}
	fmt.Printf("Took:%v\n", time.Since(t0))
}
