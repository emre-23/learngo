// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	name, last := "carl", "sagan"

	// assignment operation using string concat
	name += " edward"
	last += " "
	last += "emre"
	// equals to this:
	// name = name + " edward"

	fmt.Println(name + " " + last)
}
