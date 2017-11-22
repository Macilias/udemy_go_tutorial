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

	for _, link := range links {
		reported += checkLink(link)
	}
	end := makeTimestamp()
	fmt.Println(reported, "miliseconds reported")
	fmt.Println(end-start, "miliseconds in total")
	fmt.Println(end-start-reported, "miliseconds overhead")
}

func checkLink(link string) int64 {
	start := makeTimestamp()
	_, err := http.Get(link)
	end := makeTimestamp()
	if err != nil {
		fmt.Println(end-start, link, "might be down!")
		return end - start
	}
	fmt.Println(end-start, link, "is up!")
	return end - start
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
	//return time.Now().UnixNano()
}
