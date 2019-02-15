package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Print(len(cards))
	fmt.Println(cards)
	fmt.Println("******")
	fmt.Print(cards.shuffle())
}
