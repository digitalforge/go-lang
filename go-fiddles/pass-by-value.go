package main

import (
	"fmt"
)

// in this case we're not updating the actual 'name' variable - we are updating the copy of the varaible
func updateName(n string){
	// changing the value of the argument does not change the original value of the var!
	n = "wedge"
}

// how can we fix this? By adding a return value
func updateName2(n string) string {
	n = "wedge"
	return n
}

func updateMenu(y map[string]float64){
	y["coffee"] = 2.99
}

func main() {
	// group A Types => strings, ints, bools, floats, arrays, structs
	name := "tifa"
	name2 := "barret"

	updateName(name)

	// now we make name2 = the return value of the updateName
	name2 = updateName2(name2)

	fmt.Println(name)
	fmt.Println(name2)

	// Group B types => slices, maps, functions

	// REMEMBER - a map is just a set of key value pairs - first we declare we're using a map
	// then enter the types - in this case string : float64 (or number with decimal)
	menu := map[string]float64{
		"pie": 5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)

	fmt.Println(menu)

}