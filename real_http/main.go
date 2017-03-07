package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	blog struct {
		name string
	}
	pizzaDownloader struct {
		toppingValue int
	}
	downloader interface {
		download(chan string)
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	t0 := time.Now()
	ch := make(chan string)
	var blogs = blogList()

	for _, b := range blogs {
		go b.download(ch)
	}

	select {
	case s := <-ch:
		fmt.Printf("got: %s\n", s)
	case <-time.After(900 * time.Millisecond):
		fmt.Println("Ran out of time :-(")
	}

	fmt.Printf("Took: %v\n", time.Since(t0))
}

func (b blog) download(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	ch <- fmt.Sprintf("Downloaded %s ...\n", b.name)
}

func (p pizzaDownloader) download(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	ch <- fmt.Sprintf("PIZZA %d ...\n", p.toppingValue)
}

func blogList() []downloader {
	geo := NewGeo("http://freegeoip.net/json")
	return []downloader{
		blog{"expert-talk"},
		blog{"xkcd"},
		blog{"dilbert"},
		pizzaDownloader{7},
		&geo,
	}
}
