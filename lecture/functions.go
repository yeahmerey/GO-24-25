package main 

import (
	"fmt"
)

func add(x int , y int) int{
	return x + y 
}
func multiplicity(x int , y int) int{
	return x * y 
}
func mines(a int , b int) int {
	return a - b
}
func division(a int , b int) int {
	return a / b 
}
func main(){
	fmt.Println(add(10 , 2))
	fmt.Println(multiplicity(10 , 2))
	fmt.Println(mines(10 , 2))
	fmt.Println(division(10 , 2))
}