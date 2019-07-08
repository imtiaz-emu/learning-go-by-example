package main

import (
	"fmt"
)

func main() {
	counter := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			product := fmt.Sprintf("%d", i*j)
			if product[0] == product[len(product)-1] {
				counter++
			}
		}
	}
	fmt.Println(counter)
}
