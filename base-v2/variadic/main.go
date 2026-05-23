package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = num + total
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4)

	fmt.Println(result)
}
