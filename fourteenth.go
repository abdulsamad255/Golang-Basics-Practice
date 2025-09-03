package main
import "fmt"

func main(){
	x := 10
	p := &x // pointer to x

	fmt.Println("Value of x:", x)
	fmt.Println("Value of p:", p) // memory address
	fmt.Println("Value at address p:", *p) // value at that address 

	*p = 20 // change value at address p
	fmt.Println("New value of x:", x)
}