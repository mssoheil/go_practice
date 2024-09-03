package main

import "fmt"

func main() {

	ids := []int{33, 27, 7, 2, 11, 9}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	ages := map[string]int{"sarah": 25, "ben": 28}

	for key, value := range ages {
		fmt.Printf("%s: %d\n", key, value)
	}

	for key := range ages {
		fmt.Printf("%s\n", key)
	}

}
