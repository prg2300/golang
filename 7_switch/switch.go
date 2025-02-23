package main

import (
	"fmt"
)

func main() {
	//switch
	// i := 1
	// switch i {
	// case 1:
	// 	fmt.Println("ONE")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it is weekend")
	// default:
	// 	fmt.Println("it is workday")

	// }

	//type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("other")
		}
	}

	whoAmI(89.7)

}
