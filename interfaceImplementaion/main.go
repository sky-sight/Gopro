package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(&eb)
	printGreeting(&sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (b englishBot) getGreeting() string {
	return "Hi, There"
}

func (s spanishBot) getGreeting() string {
	return "hola"
}
