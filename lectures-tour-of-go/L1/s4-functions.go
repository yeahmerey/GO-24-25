package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}
func mines(x int, y int) int {
	return x - y
}
func multiply(x, y int) int {
	return x * y
}
func main() {
	fmt.Println("x = 44 , y = 55")
	fmt.Println("x + y = ", add(44, 55))
	fmt.Println("x - y = ", mines(44, 55))
	fmt.Println("x * y = ", multiply(44, 55))
}
