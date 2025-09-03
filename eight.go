package main
import "fmt"


func main(){
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	case "Saturday", 
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek days")
	}
}