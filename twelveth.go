package main 
import "fmt"

func main(){
	studentGrades := map[string]int{
		"Shakeel": 90,
		"saud": 85,

}
// add new key
studentGrades["Ajwad"] = 60
fmt.Println("Updated student grades:", studentGrades)

// access value
fmt.Println("Shakeel's Grade :", studentGrades["Shakeel"])

// check if key exists
grade, ok := studentGrades["Daud"]
if ok {
	fmt.Println("Daud's Grade : ", grade)
}else{
	fmt.Println("Daud's Grade not found")
}
}