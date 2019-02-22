package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[int]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fff",
	}
	// colors[10] = "#fff"
	// colors[11] = "#15f"
	// delete(colors, 10)
	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
