package main

import "fmt"

func main() {
	x := 10

	if x < 5 { // &&, ||, !
		fmt.Println("small")
	} else if x < 15 {
		fmt.Println("medium")
	} else {
		fmt.Println("large")
	}

	// short statement
	if y := 10; y > 5 {
		fmt.Println(y)
	}
}
