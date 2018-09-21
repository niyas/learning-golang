package main

import "fmt"

func main() {
  fmt.Println("Hello World")
  
  var age int = 30
  var favNum float32 = 1.387767
  const pi float64 = 3.1458765
  fmt.Println(age," ", favNum)
  
  fmt.Printf("%.3f \n", pi)

  fmt.Printf("%T \n", pi) //Prints the datatype of variable 'pi'

  i := 1
  for i <= 10 {
   fmt.Println(i)
   i++
  }
}
