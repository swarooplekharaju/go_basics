// in golang fmt acts as a inbuilt string formatter
package main

import "fmt"

func main() {

	o := 05555
	//just like "string.format(multiple options in python ) "
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xadaadad)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

}
