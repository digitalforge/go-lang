package main //this tells the go compiler it should be compiled into an executable file at the end - running that file will start the program

import "fmt" // this is a package from Go standard library 'fmt' is for formatting strings

// next we creat the 'main' function. Its called main so that Go automaticlly looks for an executes this - there must be one and only one main func in every appliation

func main() {
	// fmt.Println("Hello, ninjas!")

	// How to declare varaibles
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	// this is shorthand - you can't use the shorthand ':=' outside of a function
	nameFour := "yoshi"
	fmt.Println(nameFour)

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//NUMBERS ( INTEGERS )
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	// var numOne int8 = 25 //8 bits
	// var numTwo int8 = -128 // goes from -128 to 127
	// var numThree uint8 = 25 // can't have a negative number goes from 0 - 255

	var scoreOne float32 = -1.5 // num that has a decimal point
	var scoreTwo float64 = 8823423499229239.7
	scoreThree := 1.5

	fmt.Println(scoreOne,scoreTwo,scoreThree)

	



}




