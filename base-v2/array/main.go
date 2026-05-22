package main

import "fmt"

func main() {
	var nums [5]int

	names := [3]string{"Me", "You"}

	nums[1] = 3

	// 2d array
	matrix := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(len(nums))
	fmt.Println(nums)
	fmt.Println(names)
	fmt.Println(matrix)
}
