package main

import "fmt"

func main() {
	// age := 15
	// if age >= 18 {
	// 	fmt.Println("Person is is adult")
	// } else if age >= 12 {
	// 	fmt.Println("person is teenager")
	// } else {
	// 	fmt.Println("person is not adult")
	// }
	// var role = "admin"
	// var hasPermissions = false
	// if role == "admin" && hasPermissions {
	// 	fmt.Println("yes")
	// }
	if age := 35; age >= 18 {
		fmt.Println("ADULT", age)
	} else if age >= 12 {
		fmt.Println("teenager", age)
	}

	// go does not have ternary,you will have to use normal if else.

}
