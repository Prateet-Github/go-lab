package main

import "fmt"

// no break in go since auto break
// fallthrough
// multiple cond switch
// type switch

func main() {
	day := 2

	switch day {
	case 1:
		fmt.Println("mon")

	case 2:
		fmt.Println("tue")

	default:
		fmt.Println("invalid")
	}

}
