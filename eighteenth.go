package main
import (
	"fmt"
	"time"
)


func printMessage(msg string){
	for i :=1; i<=3; i++{
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}


func main(){
	// Run in goroutine (concurrent)
	go printMessage("Goroutine")

	// Run in main thread
	printMessage("Main")


	// wait for goroutine to finish
	time.Sleep(2 * time.Second)
}


