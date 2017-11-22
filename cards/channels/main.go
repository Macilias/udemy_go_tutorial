package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	start := makeTimestamp()
	var reported int64

	c := make(chan string)
	r := make(chan int64)

	for _, link := range links {
		go checkLink(link, r, c)
	}
	end := makeTimestamp()
	fmt.Println(reported, "miliseconds reported")
	fmt.Println(end-start, "miliseconds in total")
	fmt.Println(end-start-reported, "miliseconds overhead")

	fmt.Println("go to sleep")
	time.Sleep(10000000000)
	fmt.Println("wacke up")
}

func checkLink(link string, r chan int64, c chan string) {
	start := makeTimestamp()
	_, err := http.Get(link)
	end := makeTimestamp()
	if err != nil {
		fmt.Println(end-start, link, "might be down!")
		r <- end - start
	}
	fmt.Println(end-start, link, "is up!")
	r <- end - start
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
	//return time.Now().UnixNano()
}
