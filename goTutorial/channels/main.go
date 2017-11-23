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
		"http://vitraum.de",
		"http://www.apple.com/de/",
	}
	var reported int64
	var count int64
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			reported += checkLink(link, c)
			count++
			fmt.Println("average duration", "...........................", reported/count)
		}(l)
	}
}

func checkLink(link string, c chan string) int64 {
	start := makeTimestamp()
	_, err := http.Get(link)
	end := makeTimestamp()
	if err != nil {
		fmt.Println(end-start, link, "might be down!")
		c <- link
		return end - start
	}
	fmt.Println(end-start, link, "is up!")
	c <- link
	return end - start
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
	//return time.Now().UnixNano()
}
