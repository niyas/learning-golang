package main

import "fmt"

func main() {
  // specify a key/value pair
  m := make(map[string]int)
  m["Niyas"] = 30
  m["Azhaan"] = 3
  m["Sumi"] = 28
  fmt.Println(m)

  // Print the value of a specific key
  fmt.Println(m["Azhaan"])

  // len() function
  fmt.Println(len(m))

  // Remove a key/pair from map
  delete(m, "Niyas")
  fmt.Println(m)

  // Another way to initialize map
  // Ahead of time
  m2 := map[string]int{"key1": 1, "key2": 2}
  fmt.Println(m2)
 
  // The value and th evalue is present
  val, ok := m2["key1"]
  fmt.Println(val, ok)

  
}
