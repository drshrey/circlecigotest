package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func main() {
	fmt.Printf("2 + 2 = %d\n", add(2, 2))
	fmt.Printf("2 - 2 = %d\n", subtract(2, 2))
}
