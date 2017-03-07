package main

import (
	"bufio"
	"io"
	"fmt"
	"math/rand"
	"os"
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
		stringToWrite := []byte(fmt.Sprintf("got: %s", s))
		nBytes, err := writeStringToFile("test-file.txt", stringToWrite)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("wrote %d bytes\n", nBytes)

	case <-time.After(900 * time.Millisecond):
		fmt.Println("Ran out of time :-(")
	}

	fmt.Printf("Took: %v\n", time.Since(t0))
}

func writeData(writer io.Writer, data []byte) (int, error){
	n, err := writer.Write(data)
	return n, err
}

func writeStringToFile(filename string, buff []byte) (int, error) {
	// f, err := os.Create(filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return 0, err
	}

	w := bufio.NewWriter(f)

	// nBytes, err := w.WriteString(str)
	nBytes, err := writeData(w, buff)
	w.Flush();
	return nBytes, err
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
