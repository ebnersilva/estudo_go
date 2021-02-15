package main

import "fmt"

func main() {
	// For
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Equivalente ao While
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Loop infinito
	for {
		x++
		fmt.Println(x)
		if x == 50 {
			continue
		}

		if x == 100 {
			break
		}
	}
}
