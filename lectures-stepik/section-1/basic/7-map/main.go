package main 

import "fmt"

func main(){
	//инициализация при создании
	//map[key]value
	var user map[string]string = map[string]string{
		"name" : "Merey" , 
		"lastName" : "Kaliyev",
	}
	//cразу с нужной ёмкостью 
	profile := make(map[string]string , 10)

	//количество элементов :
	mapLenght := len(user)

	fmt.Printf("%d %+v\n" , mapLenght , profile)
	//если ключа нет - вернет значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:" , mName)

	// проверка на существование ключа
	mName , mNameExist := user["middleName"]
	fmt.Println("mName :" , mName , "mNameExist:" , mNameExist)

	//пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2" , mNameExist2)

	//удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
	

}