package main

import "fmt"


func main() {
	var name string = "John"
	var age int = 30
	// Shorthand
	job, size :=  "Developer", 1.5

	
	age = 28

	const isMale = true

	fmt.Println(name, age, job, size)
	fmt.Println("age", age)
	fmt.Println("isMale", isMale)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMale)
	fmt.Printf("%T\n", size)
}
