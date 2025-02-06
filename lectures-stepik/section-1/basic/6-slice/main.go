package main

import "fmt"

func main() {
	var buf0 []int             //len = 0 , cap =0
	buf1 := []int{}            // len = 0 , cap=0 ;
	buf2 := []int{42}          // len = 1 , cap = 1 ;
	buf3 := make([]int, 0)     // len = 0 , cap = 0 ;
	buf4 := make([]int, 5)     //len = 5 , cap = 5 ;
	buf5 := make([]int, 5, 10) //len=5 ,
	println(buf0, buf1, buf2, buf3, buf4, buf5)

	//Обращение к элементу.
	someInt := buf2[2]
	fmt.Println(someInt)

	//Добавление элементов :
	var buf []int            //len = 0 ,  cap = 0
	buf = append(buf, 9, 10) //len2 , cap= 2
	buf = append(buf, 12)    //len 3 , cap = 4

	//Добавление другого слайса :
	otherBuf := make([]int, 3)
	buf = append(buf, otherBuf...)

	//Просмотр информации про slice :
	var bufLen, bufCap int = len(buf), cap(buf)
	fmt.Println(bufLen, bufCap)

	//-----------------------//\

	bufer := []int{1, 2, 3, 4, 5}
	fmt.Println(bufer)

	//получение среза :
	sl1 := buf[1:4] //2 , 3 ,4
	sl2 := buf[:2]  //1 , 2
	sl3 := buf[2:]  //3 , 4 , 5
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:] //[1 , 2 , 3, 4 ,5]
	newBuf[0] = 9

	//buf int this case , buf=[9 , 2 , 3 , 4 , 5]
	//т.к. та же память

	newBuf = append(newBuf, 6)

	newBuf[0] = 1

	//buf = [9 , 2 , 3 , 4 , 5] //не изменился
	//newBuf = [1 , 2 ,3 ,4 , 5 , 6] //изменился

	fmt.Println(buf, newBuf)
	
	//копирование одного слайса на другой
	newBuf = make([]int ,len(buf) , len(buf))
	copy(newBuf , buf)
	fmt.Println(newBuf)

	//можно копировать в часть существующего слайса 

	ints := []int{1 , 2 , 3 , 4}
	copy(ints[1:3] , []int{5 ,6}) //ints = [1, 5, 6, 4]
	fmt.Println(ints)
	
}
