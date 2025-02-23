package main

import "fmt"

//numbered sequence of specific length

func main() {
	// var nums [4] int
	// nums[2] = 1
	// nums[3] = 4
	// fmt.Println(nums)

	// var vals [5]bool
	// vals[0] = true
	// fmt.Println(vals)

	// var name [3]string
	// name[1] = "pranal"
	// fmt.Println(name)

	//to declare in singla line
	// nums := [4]int{1, 2, 3, 4}
	// fmt.Println(nums)

	//2d arrays
	nums := [2][2]int{{3, 4}, {5, 5}}
	fmt.Println(nums)

	//-fixed size array,is predictable
	//-memory optimization
	//-constant time access

}
