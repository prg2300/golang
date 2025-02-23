package main

import (
	"fmt"
)

// -slices dynamic hote hai,phele se len dene ki jrurt nhi h
// -most used construct + useful methods
func main() {
	//uninitialized slice is nill
	// var nums []int
	// fmt.Println(len(nums))

	//nums := make([]int, 2, 5)
	//nums := []int{}

	//capacity-max num of elements
	//fmt.Println(cap(nums))
	// fmt.Println(nums == nil)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	//nums[0] = 3
	//automatically expand len ; dynamic
	//fmt.Println(nums)
	//fmt.Println(cap(nums))

	//copy slice
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//slice operator
	// var nums = []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:3])
	// fmt.Println(nums[2:])

	// slice compare
	// var nums1 = []int{1, 4}
	// var nums2 = []int{1, 2}
	// fmt.Println(slices.Equal(nums1, nums2))

	//2d
	var nums = [2][2]int{{2, 3}, {4, 5}}
	fmt.Println(nums)

}
