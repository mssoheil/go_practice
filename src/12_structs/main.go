package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// value receiver
func (person Person) greet() string {
	return "Hello I'm " + person.firstName + " " + person.lastName + " and I'm " + strconv.Itoa(person.age)
}

// pointer receiver
func (person *Person) hasBirthday() {
	person.age++
}

func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "F" {
		person.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "John", lastName: "Smith", city: "Boston", gender: "M", age: 26}
	fmt.Println(person1)

	person2 := Person{"Sarah", "Brown", "Boston", "F", 27}
	fmt.Println(person2)

	fmt.Println(person1.firstName)

	person1.age++

	fmt.Println(person1)

	person2.hasBirthday()

	fmt.Println(person2.greet())

	person2.getMarried("Williams")
	fmt.Println(person2.greet())

	person1.getMarried("Bernards")
	fmt.Println(person1.greet())

}
