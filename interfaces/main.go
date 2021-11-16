package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// fmt.Println(eb.getGreeting(), sb.getGreeting())

	printGreeting(eb)
	printGreeting(sb)
}
