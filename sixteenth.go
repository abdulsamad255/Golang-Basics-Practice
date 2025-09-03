package main
import "fmt"


// Degin interface
type Shape interface{
	Area() float64
}

// Rectangle type

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}


// Circle type	
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main(){
	r := Rectangle{width: 5, height: 10}
	c := Circle{radius: 7}
	printArea(r)
	printArea(c)
}