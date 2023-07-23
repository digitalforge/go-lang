package main

import (
	"fmt"
)

func main() {
	x := 0

	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5;  i++ {
		fmt.Println("value of i is:", i)
	}

	//looping through a slice of strings
	names := []string{"mario","luigi","yoshi","peach"}

	for i :=0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v\n", index, value)
	}

	//value is a local copy and not the orig slice
	// the underscore is a placeholder if we don't want the index
	for _, value := range names {
		fmt.Printf("the value is %v\n", value)
		value = "new string"
	}

	fmt.Println(names)
}