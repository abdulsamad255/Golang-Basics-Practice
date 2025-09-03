package main
import "fmt"

func main(){
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits{
		fmt.Println(index,fruit)
	}
}