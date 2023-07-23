package main

import "fmt"

func main() {

	// maps
	// consist of [key]value
	// first we use map then pass the key type in brackets 
	// then the value type 
	menu := map[string]float64{
		"soup":4.99,
		"pie":7.99,
		"salad":6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping through map
	// k = 'key'
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string {
		2335634: "mario",
		6523753: "luigi",
		8344647: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2335634])

	// update items inside the map

	phonebook[2335634] = "bowser"
	fmt.Println(phonebook)

	phonebook[8344647] = "yoshi"
	fmt.Println(phonebook)
}