package main //this tells the go compiler it should be compiled into an executable file at the end - running that file will start the program

import (
	"fmt" // this is a package from Go standard library 'fmt' is for formatting strings
	"sort"
	"strings"
)

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

	age := 43
	name := "Jon"

	// Print does the same thig as Println only it doesn't automatically print on a new line unless we 
	// add a '/n'
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Print Line
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf - Formatted Strings
	// '%v' means varaible. It matters what order we use them in. It will
	// find the first variable and use that and so on
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	// '%q is double quotes. It will not work on numbers
	fmt.Printf("my age is %q and my name is %q \n", age, name)

	// '%T gets the type of varaible
	fmt.Printf("age is of type %T \n", age)

	// '%f' is for floats ( numbers with decimal points in them )
	fmt.Printf("you scored %f points \n", 225.55)
	// We can limit the amount of numbers after the decimal points
	// In this case if we only use 1 after decimal point it will round up to 255.6
	fmt.Printf("you scored %0.1f points \n", 225.55)

	// Sprintf saves formatted strings in a variable
	str := fmt.Sprintf("my age is %v and my name is %v \n", age, name)

	fmt.Println("The saved string is:", str)

	// arrays and slices
	// arrays in go have a fixed length and they cannot change
	// we declare the number of items and the type:

	// var ages [3]int = [3]int{20,25,30}
	var ages = [3]int{20,25,30}

	names := [4]string{"mario","luigi","bowser","peach"}
	names[1] = "yoshi"

	fmt.Println(names, len(names))

	fmt.Println(ages, len(ages)) // we can print out the length of array using len(ages)

	// slices allow us to change array length
	// when we use an empty squre bracket it creates a slice and not and array
	// so we don't need to specify the a fixed length

	scores := []int{100,50,60}
	scores[2] = 25

	// we can append items to slices
	
	fmt.Println(scores, len(scores))
	// this alters the orig slice
	scores = append(scores,90)

	fmt.Println(scores, len(scores))



	// slice ranges - a way to get a range of elem of an existing arr or slice and store them in a new slice

	rangeOne := names[1:3] // from pos 1 to 3 but not including 3 - its inclusive of the 1st number no the 2nd
	rangeTwo := names[2:] // from 2 to the end
	rangeThree := names[:3] // from start up to pos 3 but not including pos 3
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	// since these are slices we can append 
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

	teams := []string{"Bruins","Leafs","Habs","Kings","Golden Knights"}

	fmt.Println(teams, len(teams))

	greeting := "hello there friends!"

	//strings is a package we can use on strings. Here are some examples
	// this will return a boolean if the greeting contains 'hello'


	fmt.Println(strings.Contains(greeting, "hello"))
	// modifiedGreeting := strings.ReplaceAll(greeting, "hello", "hey")
	// fmt.Println(modifiedGreeting)
	// this will replace 'hello' with 'hi'. Also this doesn't override the varaible it writes it to a new one
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) 
	// fmt.Println(strings.Index(greeting, "ll"))

	// fmt.Println(strings.Split(greeting, " "))

	// fmt.Println(strings.ToUpper(greeting))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort does alter the original value of the slice
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi","mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
	

	



}




