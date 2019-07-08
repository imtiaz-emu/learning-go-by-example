package main

import (
	"fmt"
)

func main() {
	maxRange := 21
	var i int
	for i = 1; i < maxRange; i++ {
		/* Conditional Version
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
		*/

		// switch case version
		switch {
		case i%15 == 0:
			fmt.Println("fizzbuzz")
		case i%5 == 0:
			fmt.Println("buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}
	}
}
