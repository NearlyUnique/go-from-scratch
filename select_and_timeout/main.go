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
	pizzaCollector struct {
		toppingValue int
	}
	collector interface {
		collect(chan string)
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	t0 := time.Now()
	ch := make(chan string)
	var blogs = blogList()

	for _, blog := range blogs {
		go blog.collect(ch)
	}

	select {
	case s := <-ch:
		fmt.Printf("Got %s\n", s)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Ran out of time :-(")
	}

	fmt.Printf("Done in %v\n", time.Since(t0))
}

func (b blog) collect(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	ch <- fmt.Sprintf("Collecting %s ...\n", b.name)
}

func (p pizzaCollector) collect(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	ch <- fmt.Sprintf("PIZZA %d ...\n", p.toppingValue)
}

func blogList() []collector {
	return []collector{
		blog{"expert-talk"},
		blog{"xkcd"},
		blog{"dilbert"},
		pizzaCollector{7},
	}
}
