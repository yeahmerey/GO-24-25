package main 

import "fmt"

func main() {
	//простое условие 

	boolVal := true 
	if boolVal { 
		fmt.Println("boolVal is true!")
	}

	mapVal := map[string]string{"name" : "merey"}
	//условие с блоком инициализации , in this case we check is key is available or not ;
	if keyValue , keyExist := mapVal["name"] ; keyExist{
		fmt.Println("name= ", keyValue)
	}
	//получаем только признак существование ключа	
	if _, keyExist := mapVal["name"] ; keyExist {
		fmt.Println("key 'name' exist")
	}

	cond := 1 
	//множественный if/else : 
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2{
		fmt.Println("cond is 2")
	}

	//switch по 1 переменной 
	strVal := "name"
	switch strVal {
	case "name" : 
		fallthrough
	case "test" , "lastName" : 
		//some work	
	default: 
		//some work
	}

	//switch как замена многим ifelse 

	var val1 , val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11 : 
		fmt.Println("first block")
	case val2 > 10 :
		fmt.Println("second block")
	}
Loop:
	for key , val := range mapVal { 
		println("switch in loop" , key , val)
		switch { 
		case key == "lastName":
			break 
			println("dont pront this")
		case key == "firstName" && val == "Merey" : 
			println("switch - break loop here")
			break Loop 	
		}
	}//конец for 
}	