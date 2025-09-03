package main
import "fmt"

type Student struct {
	name string
	age  int
}

func (s Student) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", s.name, s.age)
}

// Method with pointer receiver (to modify the struct)
func (s *Student) haveBirthday() {
	s.age++
}

func main() {
	student := Student{name: "Alice", age: 20}
	student.greet()
	student.haveBirthday()
	student.greet()
}