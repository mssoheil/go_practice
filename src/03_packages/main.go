package main

import (
	"fmt"
	"go_practice/src/03_packages/utils"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(utils.Reverse("John"))

}
