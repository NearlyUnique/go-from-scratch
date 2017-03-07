package main

import (
	"fmt"
	"time"
)

const imax = 3

type (
	blog struct {
		name string
	}
	pizzaDownloader struct {
		toppings int
	}
	downloader interface {
		download(ch chan string)
	}
)

func main() {
	t0 := time.Now()
	ch := make(chan string)
	blogs := blogList()
	for _, b := range blogs {
		go b.download(ch)
	}

	for range blogs {
		fmt.Printf("got: %s\n", <-ch)
	}

	fmt.Printf("Took: %v\n", time.Since(t0))
}

func (b blog) download(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- fmt.Sprintf("Downloaded %s", b.name)
}

func (p pizzaDownloader) download(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- fmt.Sprintf("Pizza with %d toppings", p.toppings)
}

func blogList() []downloader {
	return []downloader{
		blog{"apple.com"},
		blog{"meetup.com"},
		blog{"xkcd.com"},
		pizzaDownloader{toppings: 7},
	}
}
