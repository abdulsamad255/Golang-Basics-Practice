package main
import "fmt"

func worker(ch chan string){
	ch <- "Task completed"  // send message
}

func main(){
	ch := make(chan string) // create channel
	go worker(ch)
	message := <-ch  // receive message
	fmt.Println(message)
}