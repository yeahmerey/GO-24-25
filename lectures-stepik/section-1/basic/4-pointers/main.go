package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a //новый указатель на переменную а

	d := new(int)

	*d = 12
	*c = *d // c = 12 - > a = 12
	*d = 13 // c и a не изменились

	*c = 14 // c=14-> d = 14 , a = 12
	fmt.Println(a, b, c, d)
}
