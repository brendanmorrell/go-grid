package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// pretend this had custom logic related to english
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// pretend this had custom logic related to spanish
	return "Hola!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
