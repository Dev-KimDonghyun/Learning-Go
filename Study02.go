package main
import "fmt"

var myName string = "Donghyun"
var myAge int = 20

func ageCheck(name string, age int) string {
	if age >= 20 {
		return fmt.Sprintf("%s, You Can Drink", name)
	} else {
		return fmt.Sprintf("%s, You Can't Drink", name)
	}
}

func main() {
	result := ageCheck(myName, myAge)
	fmt.Println(result)
}
