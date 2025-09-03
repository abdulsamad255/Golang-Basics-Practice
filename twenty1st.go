package main
import "fmt"


func main(){
	// Buffered channel with capacity 2

	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	// ch <- 30 // would block because the buffer is full

	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 20
}