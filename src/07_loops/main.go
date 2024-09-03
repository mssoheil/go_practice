package main

import "fmt"

func main() {

	// Short method
	i := 1
	for i <= 10 {
		fmt.Println(i)

		// i = i + 1
		i++
	}

	// Long method
	for i := 0; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz", i)
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz", i)
			continue
		}

	}

}
