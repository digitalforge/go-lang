package main

import (
	"fmt"
)

func main() {
	// booleans & conditionals
	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	names := []string{"mario","luigi","yoshi","peach","bowser"}

	//this is a bit like a for in loop

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			// this tells the loops to go back up to the start and continue with the loop 
			// instead of doing the next line. So it will skip printing out the index and value of 1
			// and continue to 2
			continue 

			
		}

		if index > 2 {
				fmt.Println("breaking at pos", index)
				break // this breaks out of the loop completely
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)


	}
}