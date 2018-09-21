package main

import "fmt"

func main() {
  // range is similar to range in python
  str_arr := []string{"a", "b", "c", "d"}
  for i, c := range str_arr {
   fmt.Println(i, c)
  }

  // can also range over key/value pair in map
  m := map[string]int {"k1": 1, "k2":2, "k3":3}
  for key, val := range m {
   fmt.Println(key, val)
  }

  // can also just range over the keys
  for key := range m {
   fmt.Println(key)
  }
}
