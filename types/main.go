package main

import (
	"fmt"
	"time"
)

const imax = 3

type blog struct {
	name string
}

func main() {
	t0 := time.Now()
	ch := make(chan string)

	for _, b := range blogList() {
		go download(ch, b)
	}

	for i := 0; i < imax; i++ {
		fmt.Printf("got: %s\n", <-ch)
	}

	fmt.Printf("Took:%v\n", time.Since(t0))
}

func download(ch chan string, b blog) {
	time.Sleep(100 * time.Millisecond)
	ch <- fmt.Sprintf("Downloaded %s", b.name)
}

func blogList() []blog {
	return []blog{
		blog{"apple.com"},
		blog{"meetup.com"},
		blog{"xkcd.com"},
	}
}
