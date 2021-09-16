package main

import "fmt"



type bot interface {
		getGreetings() string
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
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string{
	return "hello world"
}

func (spanishBot) getGreetings() string{
	return "Hola!"
}