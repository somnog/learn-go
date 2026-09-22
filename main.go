package main

import "fmt"

func main() {
	fmt.Println("Hello, GO!")
	PrintWelcome()
	// How to define variables in go

	var conferenceName string = "Go Conference"

	fmt.Println(conferenceName)
	conference := Conference()
	fmt.Println(conference)
	year := Year()
	fmt.Println(year)

	welcome := WelcomeToSomNOG("Ahmed")
	fmt.Println(welcome)
}
