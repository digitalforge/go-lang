package main

import (
	"fmt"
)

// to use a pointer in the argument we use '*'
func updateName(x *string) {
	*x = "wedge"
}

func main() {
	// first we create a varible with the value of "tifa"
	name := "tifa"
	// then we create a pointer to that value using '&' in front of the variable
	m := &name
	
	fmt.Println("memory address of name is: ", &name)

	
	// fmt.Println("memory address", m)
	// fmt.Println("value at memory address", *m)

	fmt.Println(name)

	updateName(m)
	fmt.Println(name)
}