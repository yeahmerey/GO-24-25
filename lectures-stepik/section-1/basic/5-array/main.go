package main 

import "fmt"

func main(){

	var a1 [3]int ; //инициализация значениями по-умолчанию , [0, 0, 0]
	fmt.Println("a1 short", a1)

	const size = 2; 
	var a2 [size * 2]bool //[false,false,false,false]
	fmt.Println("a2" , a2)

	
	//определение размера при объявлении 
	//қанша элемент берсемде қабылда дегені
	a3 := [...]int{1 , 2 , 3}
	
	fmt.Println("a3" , a3)



}