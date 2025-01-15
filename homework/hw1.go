package main

import (
	"fmt" 
	"math"
)

func calcSqrt(x float64) float64 {
	z := 1.0
	t := 0.00001
	
	for i:=1 ; i <= 10 ; i ++ {
		fmt.Printf("z = %.6f in step %d\n", z , i)
		z = z - (z * z - x) / (2*z)
		if math.Abs(z * z - x) < t {
			fmt.Println("The goal is found !")
			break ; 
		}
	}
	return z ; 
}

func main(){
	x := float64(225)
	result := calcSqrt(x)
	fmt.Println("sqrt(x) = " , result)
	fmt.Println(x, " = " , result * result) 	
}