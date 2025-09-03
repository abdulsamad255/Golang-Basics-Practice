package main
import "fmt"

func main(){
	// Anonymous function
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Fateh ullah")


	// Closure: Function remembers outer variable
	counter :=0
	increment := func() int{
		counter++
		return counter
	}

	fmt.Println("Counter:", increment())
	fmt.Println("Counter:", increment())
}