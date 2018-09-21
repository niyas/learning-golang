package main

import "fmt"

type rect struct {
  width, height int
}

func (r rect) area() int {
  return r.width * r.height
}

func (r *rect) peri() int {
  return 2 * r.width * 2 * r.height
}
func main() {
  r := rect{width: 20, height: 10}
  fmt.Println(r.area())
  
  r_ptr := &r;
  fmt.Println(r_ptr.peri())
}
