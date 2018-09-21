package main

import (
	"fmt"
)

type employee struct {
	first_name string
	last_name  string
	age        int
}

func main() {
	fmt.Println(employee{first_name: "Mohammed", last_name: "Niyas", age: 30})
	fmt.Println(employee{"Mohammed", "Azhaan", 3})

	emp := employee{"Sumaiya", "Niyas", 28}

	fmt.Println(emp.first_name)
	fmt.Println(emp.last_name)
	fmt.Println(emp.age)

	emp_ptr := &emp
	fmt.Println(emp_ptr.first_name)

	emp_ptr.first_name = "Sumi"
	fmt.Println(emp_ptr.first_name)
	fmt.Println(emp.first_name)
}
