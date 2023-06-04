// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"reflect"
)

// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

func main() {

	fmt.Println("Type of hello is", reflect.TypeOf("hello"))
	fmt.Printf("Type of hello is %T\n", "hello")
	fmt.Printf("Type of hello is %[1]T\nAnd the type value of %d is: %[2]T\n", "hello", 2)
}
