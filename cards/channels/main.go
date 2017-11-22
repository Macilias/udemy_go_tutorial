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

	start := makeTimestamp()
	var reported int64

	c := make(chan string)
	r := make(chan int64)

	for _, link := range links {
		go checkLink(link, r, c)
	}

	for i := 0; i < len(links); i++ {

		var time = <-r
		var lastMessage = <-c
		reported += time
		//fmt.Println(time)
		fmt.Println(lastMessage)
	}

	end := makeTimestamp()
	fmt.Println(reported, "miliseconds reported")
	fmt.Println(end-start, "miliseconds in total")
	fmt.Println(reported-(end-start), "miliseconds saved due to parallelization")

	//fmt.Println("go to sleep")
	//time.Sleep(10000000000)
	//fmt.Println("wacke up")
}

func checkLink(link string, r chan int64, c chan string) {
	start := makeTimestamp()
	_, err := http.Get(link)
	end := makeTimestamp()
	if err != nil {
		fmt.Println(end-start, link, "might be down!")
		r <- end - start
		c <- "Might be down I think"
		return
	}
	fmt.Println(end-start, link, "is up!")
	r <- end - start
	c <- "Yep its up"
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
	//return time.Now().UnixNano()
}
