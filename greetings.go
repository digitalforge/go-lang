// package main

// import "fmt"

// // the main function only appears once in the whole program
// // there should only be one

// // remember we can't use the shorthand ':=' unless we are using it in a func
// // in this case we aren't so we have to use 'var'
// var points = []int{20,90,100,45,70}


// func sayHello(n string) {
// 	fmt.Println("hello", n)
// }

// func showScore() {
// 	fmt.Println("you scored this many points:", score)
// }

// // since this file is accessed by go when we run the main.go file we also need to run
// // this one so that go has access to the vars and funcs

// // to do this we : go run main.go greetings.go