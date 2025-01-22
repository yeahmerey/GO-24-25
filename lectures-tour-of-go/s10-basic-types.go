package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var(
	a int = 10 
	b int = 8 
	s string = "tag"
)

func main() {
	fmt.Println(a , b , s)
	fmt.Printf("Type: %T Value:%v", ToBe, ToBe)
	fmt.Println("")
	fmt.Printf("Type: %T Value:%v", MaxInt, MaxInt)
	fmt.Println("")
	fmt.Printf("Type: %T Value:%v", z, z)
}
