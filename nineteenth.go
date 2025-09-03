package main
import (
	"fmt"
	"time"
)


func task(name string){
	for i := 1; i <=3; i++{
		fmt.Println(name, "running:", i)
		time.Sleep(500 * time.Millisecond)
	}
}


func main(){
	go task("Worker 1")
	go task("Worker 2")
	go task("Worker 3")

	// Wait for goroutines to finish
	time.Sleep(2 * time.Second)
}