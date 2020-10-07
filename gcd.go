package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {

		x, y = y, y%x

	}
	return x

}

func main() {
	x, y := 20, 30
	fmt.Println(gcd(x, y))

}
