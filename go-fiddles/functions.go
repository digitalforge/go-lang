package main

import (
	"fmt"
	"math"
)

//functions
// We can put these outside of the main function and we can then use them in other files as well

// we declare what type the arugment is
func sayGreeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v\n", n)
}

// func can take on multiple arguments. In this we put in a slice and a function
// so we're saying for every name in the string cycle through
// then fire the function we passed through

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// r stands for radius
// in this func we're returning a value - to return a value we must specify in what
// we are returning after the arguments parenthasis - in this case its a float64
func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main() {

	// sayGreeting("Mario")
	// sayGreeting("Luigi")
	// sayBye("Mario")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)

	// so remember %0.3f is how we limit the decimal places in the output. 
	// this limits the formatted string to 3 on each of the a1 & a2 variables
	
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}