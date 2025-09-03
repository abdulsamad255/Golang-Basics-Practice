package main
import (
	"errors"
	"fmt"
)

func divide(x,y int)(int,error){
	if y == 0{
		return 0, errors.New("cannot divide by zero")
}
	return x / y, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}