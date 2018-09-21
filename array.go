package main

import "fmt"

func main() {
 var int_array[5]int
 fmt.	Println(int_array)
 
 var bool_array[10]bool
 fmt.Println(bool_array)

 var str_arr[5]string
 str_arr[0] = "Niyas"
 str_arr[1] = "Sumi"
 fmt.Println(str_arr)

 var aa[5][5]int
 count := 0
 for i:=0; i<5; i++ {
  for j:=0; j<5; j++ {
   aa[i][j] = count
   count = count + 1
  }
 }
 fmt.Println(aa)
}
