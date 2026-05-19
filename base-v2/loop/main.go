package main

import "fmt"

func main() {

	// basic for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while
	x := 0

	for x < 5 {
		fmt.Println(x)
		x++
	}

	// infinite
	for {
		fmt.Println("running")
	}

}
