package main

import "fmt"

type englishBot struct{}
type spainishBot struct{}

func main() {
	eb := englishBot{}
	sb := spainishBot{}

	printGreeting(eb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())

}

// func printGreeting(sp spainishBot) {
// 	fmt.Println(sp.getGreeting())
// }

func (englishBot) getGreeting() string {

	return "Hi there!"
}

func (spainishBot) getGreeting() string {

	return "Hola!"
}
