package main
import "fmt"

type Student struct {
	name string
	age int
	grade string
}

func main() {
	student1 := Student{name: "Alice", age: 20, grade: "A"}
	fmt.Println("Student 1:", student1)

	// Access fields 
	fmt.Println("Name:", student1.name)
}