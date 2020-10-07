package main

import "fmt"

func main() {
	x := 25
	// p is declared as a pointer to x
	p := &x
	fmt.Println(" values befor changing the pointers are ", x, p)
	//*p is the value which manages the value in x
	*p = 30 + x

	fmt.Println("vales after change in *p are", x)

}
