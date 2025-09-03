package main
import "fmt"

func main(){

	// standard for loop
	for i :=1; i<=5; i++{
		fmt.Println("Iteration:", i)
	}



	// while-style loop

	j :=1
	for j <=3 {
		fmt.Println("j is: ", j)
		j++
	}
}