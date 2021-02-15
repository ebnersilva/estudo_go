package main

import "fmt"

func main() {
	a := "Wesley"
	switch a {
		case "Bob":
			fmt.Println("Hey Bob")
		case "Maria":
			fmt.Println("Hey Maria")
		case "Joao":
			fmt.Println("Hey Joao")
		default:
			fmt.Println("I did not find your name")
	}
}
