package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("original num", num)
}

func main() {
	// address of memory location
	num := 1
	changeNum(num)
	fmt.Println("after changeNum", num)
}
