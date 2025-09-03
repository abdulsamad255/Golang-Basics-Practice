package main
import "fmt"

// slicing 

func main() {
	fruits := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	fmt.Println("Original slice:", fruits)

	// append adds items
	fruits = append(fruits, "Fig", "Grape")
	fmt.Println("After append:", fruits)

	// slice parts of a slice 
	somefruits := fruits[1:3]
	fmt.Println("Sliced portion:", somefruits)

}