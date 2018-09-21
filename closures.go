package main

import (
	"fmt"
)

func say_hello(msg string) {
	fmt.Println(msg)
}

func return_anon_func() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}

func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	say_hello("Hello")

	// anonymus function declared
	func(msg string) {
		fmt.Println(msg)
	}("Hello Anonymus")

	print_fun := return_anon_func()
	print_fun("Returned an Anonymous function")

	next_int := int_seq()
	fmt.Println(next_int())
	fmt.Println(next_int())

	next_int2 := int_seq()
	fmt.Println(next_int2())
}
