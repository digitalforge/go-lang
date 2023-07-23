package main

import (
	"fmt"
	"strings"
)

// remember if we want to return a value from a func
// we need to specify WHAT we are returning after the
// argument parenthasis - the below func returns 2 strings

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	// this is how you just declare a variable and the type
	//initials := []string{}

	// this is how you declare a variable using var
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {
	fn1, sn1 := getInitials("tifa lockhart")

	fmt.Println(fn1, sn1)
}