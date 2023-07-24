// package main

// import "fmt"

// var score = 99.5

// func main() {
// 	sayHello("mario")

// 	for _, v := range points {
// 		fmt.Println(v)
// 	}

// 	showScore()
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill{
	// bufio is a package that makes it so you can interact with the terminal
	// NewReader is a method and we call 'os.Stdin' which means
	// the operating systems standard input
	reader := bufio.NewReader(os.Stdin)

	// the underscore is a placeholder for an error 
	// when we use the reader with getInput function we created
	name, _ := getInput("Create a new bill name:", reader)

	// we then get rid of the whitespace by adding the strings package then 
	// the method TrimSpace and pass in the name
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip):", reader)
	
	switch option {
	case "a":
			name, _ := getInput("Item name:", reader)
			price, _ := getInput("Item price:", reader)

			// new package strconv - string conversion to convert the string the user input to a float64
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				promptOptions(b)
			}
			b.addItem(name, p)
			fmt.Println("Item added:", name, price)
			// fire off the function again so the user can add a new item8
			promptOptions(b)
	case "t":
			tip, _ := getInput("Enter tip amount ($):", reader)
			
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The tip must be a number")
				promptOptions(b)
			}
			b.updateTip(t)
			fmt.Println("Tip added:", tip)
			// fire off the function again so the user can add a new item8
			promptOptions(b)
	case "s":
			fmt.Println("you chose to save the bill", b)
	default:
		fmt.Println("that was not a valid option..")
		// if they don't choose a letter from the options then 
		// this fires off the function again
		promptOptions(b)
	}

}

func main(){
	mybill := createBill()
	promptOptions(mybill)
	//fmt.Println(mybill)
	// mybill := newBill("Mario's bill")

	// mybill.updateTip(10)
	// mybill.addItem("onion soup", 4.50)
	// mybill.addItem("veg pie", 8.95)
	// mybill.addItem("cake", 8.95)
	// mybill.addItem("cheese burger", 8.95)
	

	// fmt.Println(mybill.format())

}