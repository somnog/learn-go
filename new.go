package main

import "fmt"

func PrintWelcome() {
	fmt.Println("Welcome SomNOG!")
}

func Conference() string {
	return "Welcome to the conference"
}

func Year() int {
	return 2026
}

func WelcomeToSomNOG(name string) string {
	return "Welcome to SomNOG, " + name
}
