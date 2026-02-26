package main

import "fmt"

func main(){

	var age int

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	switch age {
	case 1:
	fmt.Println("Your age is 1")
		case 2:
	fmt.Println("Your age is 2")	
	case 18:
		fmt.Println("Your age is 18")
	default:
		fmt.Println("Not a valid age")	
	}
}