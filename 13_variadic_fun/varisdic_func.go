package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num

	}
	return total
}

func main() {
	// n num of parameters can be pass
	num := []int{3, 6, 7} //slice use
	result := sum(num...)
	//result := sum(3 + 7 + 9)
	fmt.Println(result)
	//fmt.Println(1, 2, 3, 4)

}
