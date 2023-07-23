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
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip: 0,
	}

	return b
}

//receiver functions
// format bill

// this limits this function to only bill objects
func (b bill) format() string {
	// this is just a copy of the bill not the real one
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		// -25 makes the spacing 25 characters long in front of the V variable
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// total 
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}