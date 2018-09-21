package main

import (
	"fmt"
)

func mult(arr ...int) int {
	total := 1
	for _, num := range arr {
		total *= num
	}
	return total
}

func main() {
	fmt.Println("this", "is", "a", "new", "func")

	//Call mult function with as many values as necessary
	fmt.Println(mult(1, 4, 5, 6))

	// Variadic function can also be applied to slices
	nums := []int{1, 2, 3, 4, 5, 8, 9}
	fmt.Println(mult(nums...))
}
