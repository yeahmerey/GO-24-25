package main

import "fmt"

func main(){
	//значение по умолчанию : in this case 0; 
	var num0 int  
	//значение при инициализации
	var num1 int = 1; 
	//пропуск типа , автомат. определяет
	var num2 = 20 
	fmt.Println(num0 , num1, num2)

	//короткое объявление переменной
	//только новых переменных 

	num := 30 ; 
	//после этого num := 31 не работает
	num += 1  
	fmt.Println("+=" , num)

	//prefix increment is not have , ++num
	num++ ;
	fmt.Println("++", num)

	//camelCase принятый стиль , under_score не принято
	userIndex := 10 
	fmt.Println(userIndex)

	//Объявление нескольких переменных 
	var weight , height int = 10 , 20 

	//присваивание в существующие переменные
	weight , height = 11 , 21 

	//короткое описание
	//хотя-бы одна переменная должна быть новой. 
	weight , age := 12,22

	fmt.Println(weight, height, age)

}