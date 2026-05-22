package main

import (
	"fmt"
	"slices"
)

func main() {
	// var sl []int
	var nums = make([]int, 0, 5)
	nums = append(nums, 1)
	var nums2 = make([]int, len(nums))
	// x := []int{1, 2, 3, 4}
	// dig := []int{}

	nums = append(nums, 2)
	nums = append(nums, 3)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 3)

	// fmt.Println(sl)
	// fmt.Println(x)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(dig)
	copy(nums2, nums)
	fmt.Println(nums, nums2)
	fmt.Println(nums[0:5])

	x := []int{1, 2, 3}
	y := []int{1, 2, 3, 4}

	fmt.Println(slices.Equal(x, y))

	var matri = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matri)

}
