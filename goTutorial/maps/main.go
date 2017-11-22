package main

import (
	"fmt"
)

func main() {
	// WAY 1:
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// WAY 2:
	//colors := make(map[string]string)
	// adding new value:
	colors["white"] = "#ffffff"
	colors["toDelete"] = "#ffffff"
	// delete a value:
	delete(colors, "toDelete")
	// If m is nil or there is no such element, delete
	// is a no-op.
	delete(colors, "toDeletee")

	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
