package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffff"

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, value := range c {
		fmt.Println(color, value)
	}
}
