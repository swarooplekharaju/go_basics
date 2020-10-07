// in this tutorial lets write two functions for value increment and pointer value return
//

package main

import "fmt"

// (argument name   type of argument) type of function return
func value_incrementer(value *int) int {
	//as value is an pointer we are incrementing its  pointer notation to increse the value
	*value++
	return *value

}

func main() {

	x := 20
	//lets declare the pointer for x
	p := &x
	value_incrementer(p)
	fmt.Println("see how the function incremented value of x using its *p pointer", x)

}
