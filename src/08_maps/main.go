package main

import "fmt"

func main() {

	emails := make(map[string]string)

	emails["john"] = "john@gmail.com"
	emails["jane"] = "jane@gmail.com"
	emails["sam"] = "sam@gmail.com"

	ages := map[string]int{"sarah": 25, "ben": 28}
	ages["tom"] = 30

	fmt.Println(emails)
	fmt.Println(ages)
	fmt.Println(len(emails))
	fmt.Println(emails["john"])

	delete(emails, "jane")
	fmt.Println(emails)

}
