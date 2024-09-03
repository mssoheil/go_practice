package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fruitArr2 := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr2)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])
	fmt.Println(fruitArr[1])
}
