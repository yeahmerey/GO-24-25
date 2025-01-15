package main 

import "fmt" 

func swap(x , y string) (string , string) {
	return y , x 
}

func swapswap(x , y string)(string  , string){
	return x , y 
}
func swappedMines(x , y int) int{
	return y - x 
}

func main(){
	a , b := swap("hello" , "world"); 
	fmt.Println(a , b)
	c , d := swapswap("smart" , "phone")
	fmt.Println(c , d) ;
	fmt.Println(swappedMines(10 , 2))
}