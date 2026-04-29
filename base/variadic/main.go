package main

import "fmt"

func sum(nums ...int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}

