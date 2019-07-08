package main

import (
	"fmt"
)

func main() {
	numbers := []int{12, 32, 34, 21, 3, 42, 67, 45}
	maximum := 0
	for _, i := range numbers {
		if maximum < i {
			maximum = i
		}
	}

	fmt.Println(maximum)
}
