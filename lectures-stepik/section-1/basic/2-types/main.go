package main

import "fmt"

func main(){
	var i int = 10 
	
	//int8 , int16 , int32 , int64
	var autoInt = -10 
	var bigInt int64 = 1<<32-1
	var unsignedInt uint = 100500
	var unsignedBigInt uint64 = 1<<64-1 

	fmt.Println(i , autoInt ,bigInt , unsignedInt , unsignedBigInt)

	//float : 
	var pi float32 = 3.141
	var e = 2.718 
	goldenRatio := 1.618 
	
	fmt.Println(pi , e , goldenRatio)

	//bool : 
	//по-умолчанию :false 
	var b bool 
	var isOk bool = true ; 
	var success = true 
	cond := true 

	fmt.Println(b , isOk , success , cond)

	//complex64 , complex128
	var c complex128 = -1.1 + 7.12i 
	c2 := -1.1 + 7.12i 
	
	fmt.Println(c, c2)

	var str string //по умолчанию пустая строка
	var hello string =  "Hello \n \t"
	var world string =	`without \n \t`

	fmt.Println(str , hello , world)

	//конкатенация строк 
	helloWorld := "Hello , World"
	andGoodMorning := helloWorld + " and Good Morning!"
	//строки не изменяемы , we cannot assign to helloWorld[0]

	fmt.Println(helloWorld , andGoodMorning)

	//получение длины строки 
	byteLen := len(helloWorld)

	//получение строки подстроки не в байтах , а в символах 
	hello_h :=helloWorld[:12] //Привет , 0-11 байты
	fmt.Println(byteLen , hello_h)

	//----constants--------
	const pi_2 = 3.133
	const(
		hello_word = "Hello"
		e_equals = 2.718 
	)

	


}