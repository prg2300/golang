package main

import "fmt"

func counter() func() int { //receive nhi kia pr return hoga
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())

}
