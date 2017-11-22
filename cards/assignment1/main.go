package main

import "fmt"

func main() {

	slice := make([]int, 10)

	for i := 0; i < 10; i++ {
		slice[i] = i
	}

	for i := 0; i < 10; i++ {
		if slice[i]%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
