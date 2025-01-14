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
)

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 6 {
			break
		}
		sum += i
		fmt.Println(i, "--->", sum)
		i += 2 // will be 1,3,5 in sequence

	}
	fmt.Println(sum)

}
