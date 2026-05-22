package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	// for i := 0; i < len(n); i++ {
	// 	fmt.Println(n[i])
	// }

	for i, n := range nums {
		fmt.Println(i, n)
	}

	// unicode point rune
		for i, n := range "Prateet"{
		fmt.Println(i, n)
	}
}
