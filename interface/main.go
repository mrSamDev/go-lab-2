package main

import "fmt"

type english struct {
	greeting string
}

type spanish struct {
	greeting string
}

func (e english) getGreeting(name string) string {
	e.greeting = "Hello " + name
	return e.greeting
}

func (s spanish) getGreeting(name string) string {
	s.greeting = "Ola " + name
	return s.greeting
}

func main() {
	e := english{}
	s := spanish{}

	printGreetingE(e)
	printGreetingS(s)
	printGreeting(e, "SIJO")
	printGreeting(s, "Sam")

	printGreetingX(e, "SIJO")
}

func printGreetingE(e english) {
	fmt.Println(e.greeting)
}
func printGreetingS(s spanish) {
	fmt.Println(s.greeting)
}

type greetingInterface interface {
	getGreeting(name string) string
}

func printGreeting(g greetingInterface, name string) {
	fmt.Println(g.getGreeting(name))
}

func printGreetingX[T greetingInterface](greeter T, name string) {
	fmt.Println(greeter.getGreeting(name))
}
