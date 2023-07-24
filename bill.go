package main

import "fmt"

// struct is like a class in js
// it is a blueprint that describes a type of data

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

//receiver functions - these also only create a copy 
// format bill

// this limits this function to only bill objects
func (b *bill) format() string {
	// this is just a copy of the bill not the real one
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		// -25 makes the spacing 25 characters long in front of the V variable
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip", b.tip)
	// total before tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total before tip:", total)
	// total 
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)


	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	// the '*' is the de-referencer but in a struct we need to put it in parenthasis
	// if you're reading this in the future it instead references the pointer instead of the copy
	// so the real value gets updated

	// One last caveat - we really don't need the parenthesis when de-referecning a struck
	// go will automaitically do it
	//(*b).tip = tip
	b.tip = tip
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}