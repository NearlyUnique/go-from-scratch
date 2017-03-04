package main

import (
	"fmt"
	"log"
	"time"
)

const imax = 10

func main() {
	t0 := time.Now()
	ch := make(chan string)

	for i := 0; i < imax; i++ {
		go download(ch, i)
	}

	for i := 0; i < imax; i++ {
		log.Print("\twait read")
		fmt.Printf("got: %s\n", <-ch)
	}

	fmt.Printf("Took:%v\n", time.Since(t0))
}

func download(ch chan string, i int) {
	time.Sleep(100 * time.Millisecond)
	log.Printf("\twait write chan:%d", i)
	ch <- fmt.Sprintf("Downloaded %d", i)
	log.Printf("\twrite complete chan:%d", i)
}
