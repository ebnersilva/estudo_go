package main

import "fmt"

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := "Uouuu"

	const constA = 10
	const xvz int = 1333

	const (
		aa string = "Oiii"
		bb = 66
		cc = 77
	)

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

	fmt.Printf("%v \n", constA)
	fmt.Printf("%v \n", xvz)
	fmt.Printf("%v \n", aa)
	fmt.Printf("%v \n", bb)
	fmt.Printf("%v \n", cc)
}
