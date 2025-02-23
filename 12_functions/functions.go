package main

import "fmt"

func sub(a, b, c int) int {
	return a + b + c
}

// func getlanguages() (string, string) {
// 	return "golang", "javascript"
// }

//to return func use processIt
// func processIt() func(a int) int {
// 	return func(a int) int {
// 		return 4
// 	}
// }

func main() {
	// fn := processIt()
	// fn(6)
	//fmt.Println(fn)
	// lang1, lang2 := getlanguages()
	// result := getlanguages()
	// fmt.Println(lang1, lang2)
	result := (4 + 6 + 8)
	fmt.Println(result)
}
