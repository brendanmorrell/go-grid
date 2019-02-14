package main

import "fmt"

type deck []string
type suits []string
type values []string

func newDeck() deck {
	cards := deck{}

	cardSuits := suits{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := values{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
