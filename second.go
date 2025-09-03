package main
import "fmt"

func main(){
	var age int = 25


	// type inference go gueses the type

	var name = "John"

	// short declaration (most common)
	city := "Islamabad"

	//constant
	const PI = 3.14
	
	fmt.Println(age,name,city,PI)
}