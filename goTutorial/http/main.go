package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		panic("WTF!")
	}
	// By feat approach:
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	// one line example
	//io.Copy(os.Stdout, resp.Body)

	// implement your own writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("CUSTOM WRITER")
	fmt.Println(string(bs))
	return len(bs), nil
}
