package main

import "fmt"

func main() {
  sum := add(4, 5)
  sum2 := add3(2, 4, 5)
  fmt.Println(sum)
  fmt.Println(sum2)
}

func add(a int, b int) int {
  return a + b
}

func add3(a, b, c int) int {
  return a + b + c
}
