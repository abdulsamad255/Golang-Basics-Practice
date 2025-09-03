package main

import "fmt"

func add(x int, y int) int{
	return x+y
}


func divide(x,y int)(int, int){
	return x/y, x%y
}


func main(){
	result := add(5,7)
	fmt.Println("Result:", result)

	quotient, remainder := divide(10,3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}