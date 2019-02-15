package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, x := range s {
		if x%2 == 0 {
			fmt.Printf(" %v is even", x)
		} else {
			fmt.Printf(" %v is odd", x)
		}
	}
}
