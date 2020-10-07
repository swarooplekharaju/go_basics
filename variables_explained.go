package main

import "fmt"

func main() {
	// single variable

	i := 22
	// intialization and declaration in one  go
	l := 22

	// declaring multiple variables

	a, b := 22.5, 22

	fmt.Println(i, l, a, b)
	// or you can declare a variable and assign the values with = symbol later
	var x, y bool
	x = true
	y = false
	fmt.Println(x, y)

}
