package main

import "fmt"

func main() {

	x, y := 0, 1
	fmt.Println("enter n: how many sequence of inputs do you want")

	n := 20
	for i := 0; i <= n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}

}
